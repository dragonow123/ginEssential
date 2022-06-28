package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oceanlearn.teach/ginessential/common"
	"oceanlearn.teach/ginessential/model"
	"oceanlearn.teach/ginessential/response"
	"oceanlearn.teach/ginessential/vo"
	"strconv"
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

	return CategoryController{db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
		return
	}
	//var requestCategory model.Category
	//ctx.Bind(&requestCategory)

	//if requestCategory.Name == "" {
	//	response.Fail(ctx, "数据验证错误，分类名称必填", nil)
	//}
	category := model.Category{Name: requestCategory.Name}
	//c.DB.Create(&requestCategory)
	c.DB.Create(&category)

	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	//var requestCategory model.Category
	//ctx.Bind(&requestCategory)

	//if requestCategory.Name == "" {
	//	response.Fail(ctx, "数据验证错误，分类名称必填", nil)
	//}
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
		return
	}

	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category
	if err := c.DB.First(&updateCategory, categoryId).Error; err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}

	// 更新分类
	//map
	//struct
	//name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category
	if err := c.DB.First(&category, categoryId).Error; err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, "删除失败，请重试", nil)
	}
	response.Success(ctx, nil, "")
	return
}
