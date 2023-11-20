package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int    `gorm:"type:int;not null;auto increment" json:"id" label:"用户id"`
	NickName  string `gorm:"type:string;column:nickName" json:"nickName"  label:"用户名称"`
	UserSex   string `gorm:"type:string" json:"user_sex" label:"性别"`
	UserPhone string `gorm:"type:string" json:"user_phone" label:"手机号"`
	UserEmail string `gorm:"type:varchar(100)" json:"user_email" label:"用户邮箱"`
	UserPhoto string `gorm:"type:varchar(100)" json:"user_photo" label:"用户头像"`
}

// 根据分页获取用户信息
func (User) GetListByPage(pageIndex int, nickName string, userPhone string, userSex string) ([]User, int64) {
	var data []User
	var total int64
	err := Db.Where("nickName LIKE ? and user_phone LIKE ? and user_sex LIKE ?", "%"+nickName+"%", "%"+userPhone+"%", "%"+userSex+"%").Limit(10).Offset((pageIndex - 1) * 10).Find(&data).Count(&total).Error
	if err != nil {
		fmt.Println("Select err:", err)
		return nil, 0
	}
	//Db.Model(&User{}).Where("nickName LIKE ? and user_phone LIKE ? and user_sex LIKE ?", "%"+nickName+"%", "%"+userPhone+"%", "%"+userSex+"%")
	//if err != nil {
	//	fmt.Println("Db.Model(&data).Where err:", err)
	//	return data, 0
	//}
	return data, total
}

// 根据id删除用户信息
func (User) DeleteById(id int) (string, error) {
	err := Db.Delete(&User{Id: id}).Error
	if err != nil {
		return "删除失败", err
	}
	return "删除成功", err
}
