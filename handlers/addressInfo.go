package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"orderSystem/models"
	"strconv"
)

type Address models.Address

var address models.Address

// GetListByPage 根据分页获取用户地址信息
func (Address) GetListByPage(c *gin.Context) {
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", string(0)))
	nickName, phone := c.Query("nickName"), c.Query("phone")
	data, total := address.GetListByPage(pageIndex, nickName, phone)
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

// 根据id删除收货地址信息

func (Address) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("strconv:", err)
		return
	}
	msg, err := address.DeleteById(id)
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
