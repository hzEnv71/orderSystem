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

type Admin models.Admin

var admin models.Admin

type UP struct {
	UName string `json:"zh"`
	Pwd   string `json:"admin_pwd"`
}

// CheckLogin 管理员登录检测
func (Admin) CheckLogin(c *gin.Context) {
	var up UP
	err := c.ShouldBindJSON(&up)
	if err != nil {
		fmt.Println("shouldBindJSON_err:", err)
		return
	}
	data, msg := admin.CheckLogin(up.UName, up.Pwd)

	if msg == "验证失败" {
		c.JSON(
			http.StatusOK, gin.H{
				"status": "failure",
				"msg":    msg,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    msg,
			"data": gin.H{
				"loginUserInfo": data,
				"token":         'a',
			},
		},
	)
}

// GetListByPage 根据分页获取管理员信息信息
func (Admin) GetListByPage(c *gin.Context) {

	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", string(0)))
	adminName, adminSex, adminTel, adminEmail := c.Query("admin_name"), c.Query("admin_sex"), c.Query("admin_tel"), c.Query("admin_email")
	data, total := admin.GetListByPage(pageIndex, adminName, adminSex, adminTel, adminEmail)
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

// DeleteById 根据id删除管理员信息
func (Admin) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("strconv:", err)
		return
	}
	msg, err := admin.DeleteById(id)
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

// FindById 根据id查找管理员信息
func (Admin) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		fmt.Println("strconv:", err)
	}
	data, err := admin.FindById(id)
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

// Update 编辑管理员
func (Admin) Update(c *gin.Context) {
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		fmt.Println("c.ShouldBindJSON_err:", err)
		return
	}
	msg, err := admin.Update(admin)
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

// Add 添加管理员
func (Admin) Add(c *gin.Context) {
	var _admin models.Admin
	err := c.ShouldBindJSON(&_admin)
	if err != nil {
		fmt.Println("c.ShouldBindJSON_err:", err)
		return
	}
	foodObject, err := admin.Add(_admin)
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

// GetAllList 获取所有管理员信息
func (Admin) GetAllList(c *gin.Context) {
	data, err := admin.GetAllList()
	if err != nil {
		fmt.Println("getAllList_err:", err)
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

// GetTotalInfo 获取总计信息

func (Admin) GetTotalInfo(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    "获取数据成功",
			"data": gin.H{
				"user_total":  "324.00",
				"order_total": "570.00",
				"food_total":  "146.00",
				"money_total": "96542.70",
			},
		},
	)
}

// GetCategoryFoodCount 获取菜品分类总数
func (Admin) GetCategoryFoodCount(c *gin.Context) {

}

// 上传头像
func (Admin) UploadAdminPhoto(c *gin.Context) {
	file, err := c.FormFile("admin_photo")
	if err != nil {
		fmt.Println("form-file _err:", err)
		return
	}
	in, _ := file.Open()
	defer in.Close()
	out, _ := os.Create("web/src/assets/img/" + file.Filename)
	defer out.Close()
	io.Copy(out, in)
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"msg":    "图片上传成功",
			"data":   "src/assets/img/" + file.Filename,
		},
	)
}
