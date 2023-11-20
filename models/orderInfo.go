package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Id          int       `json:"id"`
	Uid         int       `json:"uid"`
	SubmitTime  time.Time `json:"submit_time"`
	OrderStatus string    `json:"order_status"`
	OrderMoney  string    `json:"order_money"`
	PayTime     time.Time `json:"pay_time"`
	DeliverTime time.Time `json:"deliver_time"`
	Aid         int       `json:"aid"`
}

type OrderAddressUser struct {
	gorm.Model
	Id          int       `json:"id"`
	Uid         int       `json:"uid"`
	SubmitTime  time.Time `json:"submit_time"`
	OrderStatus string    `json:"order_status"`
	OrderMoney  string    `json:"order_money"`
	PayTime     time.Time `json:"pay_time"`
	DeliverTime time.Time `json:"deliver_time"`
	Aid         int       `json:"aid"`
	Addresss    Address   `gorm:"foreignKey:aid" json:"addressInfo"`
	Users       User      `gorm:"foreignKey:uid" json:"userInfo"`
}

// GetListByPage 根据分页获取订单信息
func (Order) GetListByPage(pageIndex int, startTime, endTime time.Time, orderStatus string, uid string) ([]OrderAddressUser, int64) {
	var data []OrderAddressUser
	var total int64
	err := Db.
		Joins("Addresss").
		Joins("Users").
		Where("submit_time > ?", startTime).
		Where("submit_time < ?", endTime).
		Where("order_status LIKE ?", "%"+orderStatus+"%").
		Where("order.uid LIKE ?", "%"+uid+"%").
		Limit(10).Offset((pageIndex - 1) * 10).Find(&data).Count(&total).Error
	if err != nil {
		fmt.Println("Select err:", err)
		return nil, 0
	}
	//fmt.Println(data)
	return data, total
}

// UpdateOrderToDeliver 根据id发货
func (Order) UpdateOrderToDeliver(id int) {}
