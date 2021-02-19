package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goEssential/common"
	"github.com/goEssential/dto"
	"github.com/goEssential/model"
	"github.com/goEssential/response"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {

	DB := common.GetDB()
	// 获取参数
	name := c.PostForm("name")

	phone := c.PostForm("phone")
	password := c.PostForm("password")

	log.Println(name, password, phone)
	if len(name) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名必填",
		})
		return
	}

	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码至少为6位",
		})
		return
	}

	// 查询手机号是否存在
	if isPhoneExist(DB, phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已经存在",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "加密错误",
		})
		return
	}

	//创建用户
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hashedPassword),
	}
	response.Success(c, nil, "注册成功")
	DB.Create(&newUser)
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("生成token异常 : %v", err)
		return
	}
	response.Success(c, gin.H{
		"token": token,
	}, "登录成功")

}

func GetUserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User)),
		},
	})
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	return user.ID != 0
}
