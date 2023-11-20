package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	Id           int       `gorm:"type:int;not null;auto increment" json:"id" lable:"菜品id"`
	CategoryName string    `gorm:"type:varchar(20)" json:"category_name" label:"菜品分类名称"`
	CreateTime   time.Time `gorm:"type:time" json:"create_time" label:"菜品添加时间"`
}

// 根据分页获取菜品分类信息
func (Category) GetListByPage(pageIndex int, categoryName string) ([]Category, int64) {
	var data []Category
	var total int64
	err := Db.Where("category_name LIKE ?", "%"+categoryName+"%").Limit(10).Offset((pageIndex - 1) * 10).Find(&data).Count(&total).Error
	if err != nil {
		fmt.Println("Select err:", err)
		return nil, 0
	}
	return data, total
}

// 新增分类
func (Category) Add(categoryName string) (Category, error) {
	var category = Category{CategoryName: categoryName, CreateTime: time.Now()}
	err := Db.Create(&category).Error
	if err != nil {
		return Category{}, err
	}
	return category, err
}

// 获取所有菜品分类信息
func (Category) GetAllList() ([]Category, error) {
	var data []Category
	err := Db.Find(&data).Error
	if err != nil {
		fmt.Println("Select err:", err)
		return nil, err
	}
	return data, err
}

// 根据id删除菜品分类
func (Category) DeleteById(id int) (string, error) {
	err := Db.Delete(&Category{Id: id}).Error
	if err != nil {
		return "删除失败", err
	}
	return "删除成功", err
}

// 根据id查找菜品分类信息
func (Category) FindById(id int) (Category, error) {
	var data Category
	err := Db.Where("id=?", id).Find(&data).Error
	if err != nil {
		fmt.Println("FindById err:", err)
	}
	return data, err
}

// 编辑菜品
func (Category) Update(category *Category) (string, error) {
	//err := Db.Model(&Category{Id: id}).Update("categoryName", categoryName).Error
	err := Db.Model(&Category{}).Where("id = ?", category.Id).Update("category_name", category.CategoryName).Error
	if err != nil {
		return "修改失败", err
	}
	return "修改成功", err
}
