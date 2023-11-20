package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"orderSystem/models"
	"strconv"
	"time"
)

type Order models.Order

var order models.Order

// GetListByPage 根据分页获取订单信息
func (Order) GetListByPage(c *gin.Context) {
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", string(0)))
	startTime, endTime, orderStatus, uid := c.Query("start_time"), c.Query("end_time"), c.Query("order_status"), c.Query("uid")
	layout := "2006-01-02 15:04:05.000000"
	//fmt.Println("startTime====", startTime, "endTime====", endTime, "orderStatus===", orderStatus)
	if startTime == "" {
		startTime = "1000-01-01 00:00:00.000000"
	}
	_startTime, err := time.Parse(layout, startTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	if endTime == "" {
		endTime = "2050-01-01 00:00:00.000000"
	}
	_endTime, err := time.Parse(layout, endTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println("------startTime====", _startTime, "--------endTime====", _endTime)
	data, total := order.GetListByPage(pageIndex, _startTime, _endTime, orderStatus, uid)
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

// UpdateOrderToDeliver 根据id发货
func (Order) UpdateOrderToDeliver(c *gin.Context) {

}
