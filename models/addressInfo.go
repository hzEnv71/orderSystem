package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Address struct { //一个用户有多个地址
	gorm.Model
	Id         int    `gorm:"type:int;not null;auto increment" json:"id" label:"收货id"`
	Address    string `json:"address" label:"联系人地址"`
	Phone      string `json:"phone" label:"联系人手机号"`
	PersonName string `json:"person_name" label:"联系人名称"`
	Tag        string `json:"tag" label:"地址标签"`
	Uid        int    `json:"uid" label:"用户id"`
}
type AddressUser struct { //一个用户有多个地址
	gorm.Model
	Id         int    `gorm:"type:int;not null;auto increment" json:"id" label:"收货id"`
	Address    string `json:"address" label:"联系人地址"`
	Phone      string `json:"phone" label:"联系人手机号"`
	PersonName string `json:"person_name" label:"联系人名称"`
	Tag        string `json:"tag" label:"地址标签"`
	Uid        int    `json:"uid" label:"用户id"`
	Users      User   `gorm:"foreignKey:uid" json:"userInfo"`
}

// GetListByPage 根据分页获取用户地址信息
func (address Address) GetListByPage(pageIndex int, nickName string, phone string) ([]AddressUser, int64) {
	var data []AddressUser
	var total int64
	err := Db.
		Joins("Users").
		Where("nickName LIKE ?", "%"+nickName+"%").
		Where("phone LIKE ?", "%"+phone+"%").
		Limit(10).Offset((pageIndex - 1) * 10).Order("id").Find(&data).Count(&total).Error
	if err != nil {
		fmt.Println("Select err:", err)
		return nil, 0
	}
	fmt.Println(data)
	return data, total
}

// 根据id删除收货地址信息

func (address Address) DeleteById(id int) (string, error) {
	err := Db.Where("id=?", id).Delete(&Address{}).Error
	if err != nil {
		return "删除失败", err
	}
	return "删除成功", err
}
