package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"orderSystem/models"
	"os"
	"strconv"
)

type Food models.Food

var food models.Food

// GetListByPage 根据分页获取用户信息
func (Food) GetListByPage(c *gin.Context) {
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", string(0)))
	foodName, categoryName := c.Query("food_name"), c.Query("category_name")
	data, total := food.GetListByPage(pageIndex, foodName, categoryName)
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

// SetFoodInfoWeight 设置菜品权重
func (Food) SetFoodInfoWeight(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("strconv id_ Err:", err)
		return
	}
	weight, err := strconv.Atoi(c.Query("weight"))
	if err != nil {
		fmt.Println("strconv weight_ Err:", err)
		return
	}
	msg, err := food.SetFoodInfoWeight(id, weight)
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

// Add 新增菜品
func (Food) Add(c *gin.Context) {
	var _food models.Food
	err := c.ShouldBindJSON(&_food)
	if err != nil {
		fmt.Println("c.ShouldBindJSON_err:", err)
		return
	}
	foodObject, err := food.Add(_food)
	if err != nil {
		fmt.Println("Create Err:", err)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    "新增成功",
			"data":   foodObject,
		},
	)
}

// DeleteById 根据id删除菜品
func (Food) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	fmt.Println("id ======", id)
	if err != nil {
		fmt.Println("strconv:", err)
		return
	}
	msg, err := food.DeleteById(id)
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

// FindById 根据id查找菜品
func (Food) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		fmt.Println("strconv:", err)
	}
	data, err := food.FindById(id)
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
func (Food) Update(c *gin.Context) {
	err := c.ShouldBindJSON(&food)
	if err != nil {
		fmt.Println("c.ShouldBindJSON_err:", err)
		return
	}
	msg, err := food.Update(&food)
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

func (Food) UploadFoodPhoto(c *gin.Context) {
	file, err := c.FormFile("food_img")
	if err != nil {
		fmt.Println("form-file _err:", err)
		return
	}
	in, _ := file.Open()
	defer in.Close()
	out, _ := os.Create("web/src/assets/img/" + file.Filename)
	defer out.Close()
	io.Copy(out, in)
	//c.File(file.Filename) //将文件传给前端
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    "图片上传成功",
			"data":   "src/assets/img/" + file.Filename,
		},
	)
}
