package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"orderSystem/models"
	"strconv"
)

type Category models.Category

var category models.Category

// GetListByPage 根据分页获取菜品分类信息
func (Category) GetListByPage(c *gin.Context) {
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", string(0)))
	categoryName := c.Query("category_name")
	data, total := category.GetListByPage(pageIndex, categoryName)
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    "获取数据成功",
			"data": gin.H{
				"pageIndex":  pageIndex,
				"totalCount": total,
				"pageCount":  (total + 9) / 10,
				"listData":   data,
				"pageStart":  pageIndex,
				"pageEnd":    (total + 9) / 10,
			},
		},
	)
}

// Add 新增分类
func (Category) Add(c *gin.Context) {
	categoryObject, err := category.Add(c.Query("category_name"))
	if err != nil {
		fmt.Println("Create Err:", err)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    "新增成功",
			"data":   categoryObject,
		},
	)
}

// GetAllList 获取所有菜品分类信息
func (Category) GetAllList(c *gin.Context) {
	data, err := category.GetAllList()
	if err != nil {
		fmt.Println("category.GetAllList() err:", err)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    "新增成功",
			"data":   data,
		},
	)

}

// DeleteById 根据id删除菜品分类
func (Category) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("strconv:", err)
		return
	}
	msg, err := category.DeleteById(id)
	if err != nil {
		fmt.Println("Delete Err:", err)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    msg,
		},
	)
}

// FindById 根据id查找菜品分类信息
func (Category) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("strconv:", err)
		return
	}
	data, err := category.FindById(id)
	if err != nil {
		fmt.Println("FindById Err:", err)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"flag": "success",
			"msg":  "查询成功",
			"data": data,
		},
	)
}

// Update 编辑菜品
func (Category) Update(c *gin.Context) {
	err := c.ShouldBindJSON(&category)
	if err != nil {
		fmt.Println("c.ShouldBindJSON_err:", err)
		return
	}
	msg, err := category.Update(&category)
	if err != nil {
		fmt.Println("Update Err:", err)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    msg,
		},
	)
}
