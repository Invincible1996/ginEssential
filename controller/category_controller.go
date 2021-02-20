package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goEssential/common"
	"github.com/goEssential/model"
	"github.com/goEssential/response"
	"gorm.io/gorm"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名必填")
		return
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	panic("implement me")
}

func (c CategoryController) Show(ctx *gin.Context) {
	panic("implement me")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	panic("implement me")
}
