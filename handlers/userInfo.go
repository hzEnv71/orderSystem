package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"orderSystem/models"
	"strconv"
)

type User models.User

var user models.User

// GetListByPage 根据分页获取用户信息
func (User) GetListByPage(c *gin.Context) {
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", string(0)))
	nickName, userPhone, userSex := c.Query("nickName"), c.Query("user_phone"), c.Query("user_sex")
	data, total := user.GetListByPage(pageIndex, nickName, userPhone, userSex)
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

// DeleteById 根据id删除用户信息
func (User) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("DeleteById strconv:", err)
		return
	}
	msg, err := user.DeleteById(id)
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
