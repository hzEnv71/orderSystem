package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

/*
 * 编辑管理员
 *  @param {*} id 菜品id
 * @param {*} admin_name 管理员名称
 * @param {*} admin_sex 管理员性别
 * @param {*} admin_tel 管理员手机号
 * @param {*} admin_email 管理员邮箱
 * @param {*} admin_photo 管理员头像
 * @param {*} admin_type 管理员类型
 * @returns {Promise<AxiosResponse>}
 */
type Admin struct {
	gorm.Model
	Id         int    `gorm:"type:int;not null;auto increment" json:"id"`
	AdminName  string `gorm:"type:varchar(20)" json:"admin_name"  label:"管理员名称"`
	AdminPwd   string `gorm:"type:varchar(20)" json:"admin_pwd"  label:"管理员密码"`
	AdminSex   string `gorm:"type:varchar(20)" json:"admin_sex" label:"管理员性别"`
	AdminTel   string `gorm:"type:varchar(20)" json:"admin_tel" label:"管理员手机号"`
	AdminEmail string `gorm:"type:varchar(100)" json:"admin_email" label:"管理员邮箱"`
	AdminPhoto string `gorm:"type:varchar(100)" json:"admin_photo" label:"管理员头像"`
	AdminType  int    `gorm:"type:int" json:"admin_type" label:"管理员类型"`
}

// 管理员登录检测
func (Admin) CheckLogin(uname string, pwd string) (Admin, string) {
	var admin Admin
	err := Db.Where("admin_name=?", uname).Find(&admin).Error
	if err != nil {
		fmt.Println("find_err:", err)

	}
	if admin.AdminPwd == pwd {
		return admin, "验证成功"
	}
	return Admin{}, "验证失败"
}

// 根据分页获取管理员信息信息
func (Admin) GetListByPage(pageIndex int, adminName, adminSex, adminTel, adminEmail string) ([]Admin, int64) {

	var data []Admin
	var total int64
	err := Db.Where("admin_name LIKE ? and admin_sex LIKE ? and admin_tel LIKE ? and admin_email LIKE ?", "%"+adminName+"%", "%"+adminSex+"%", "%"+adminTel+"%", "%"+adminEmail+"%").Limit(10).Offset((pageIndex - 1) * 10).Find(&data).Count(&total).Error
	if err != nil {
		fmt.Println("Select err:", err)
		return nil, 0
	}

	return data, total
}

// 根据id删除管理员信息
func (Admin) DeleteById(id int) (string, error) {
	err := Db.Delete(&Admin{Id: id}).Error
	if err != nil {
		return "删除失败", err
	}
	return "删除成功", err
}

// 根据id查找管理员信息
func (Admin) FindById(id int) (Admin, error) {
	var data Admin
	err := Db.Where("id=?", id).Find(&data).Error
	if err != nil {
		fmt.Println("FindById err:", err)
	}
	return data, err
}

// 编辑管理员
func (Admin) Update(admin Admin) (string, error) {
	//err := Db.Model(&Admin{}).Where("id = ?", admin.Id).Updates(Food{FoodName: food.FoodName, Price: food.Price, Cid: food.Cid, FoodImg: food.FoodImg, FoodDesc: food.FoodDesc, Weight: food.Weight}).Error
	admin.CreatedAt = time.Now()
	err := Db.Model(&Admin{}).Where("id = ?", admin.Id).Save(admin).Error
	if err != nil {
		return "修改失败", err
	}
	return "修改成功", err
}

// 添加管理员
func (Admin) Add(admin Admin) (Admin, error) {
	err := Db.Create(&admin).Error
	if err != nil {
		return Admin{}, err
	}
	return admin, err
}

// 获取所有管理员信息
func (Admin) GetAllList() ([]Admin, error) {
	var data []Admin
	err := Db.Find(&data).Error
	if err != nil {
		fmt.Println("Select err:", err)
		return nil, err
	}
	return data, err
}

// 获取总计信息

func (Admin) GetTotalInfo() {}

// 获取菜品分类总数
func (Admin) GetCategoryFoodCount() {}

// 上传头像
func (Admin) UploadAdminPhoto() (string, error) {
	return "s", nil
}
