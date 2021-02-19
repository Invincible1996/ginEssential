package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goEssential/common"
	"github.com/goEssential/model"
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

func (c2 CategoryController) Create(c *gin.Context) {
	panic("implement me")
}

func (c2 CategoryController) Update(c *gin.Context) {
	panic("implement me")
}

func (c2 CategoryController) Show(c *gin.Context) {
	panic("implement me")
}

func (c2 CategoryController) Delete(c *gin.Context) {
	panic("implement me")
}
