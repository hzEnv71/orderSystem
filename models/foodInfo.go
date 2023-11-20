package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Id        int     `gorm:"type:int;not null;auto increment;column:id" json:"id"`
	FoodName  string  `gorm:"type:varchar(20);column:food_name" json:"food_name"  label:"菜品名称"`
	Price     float64 `gorm:"type:double;column:price" json:"price" label:"菜品价格"`
	Cid       int     `gorm:"type:int;column:cid" json:"cid" label:"菜品分类cid"`
	FoodDesc  string  `gorm:"type:varchar(100);column:food_desc" json:"food_desc" label:"菜品描述"`
	FoodImg   string  `gorm:"type:varchar(100);column:food_img" json:"food_img" label:"菜品图片"`
	Weight    float64 `gorm:"type:double;column:weight" json:"weight" label:"菜品权重"`
	SaleCount int     `gorm:"type:int;column:sale_count" json:"saleCount" label:"菜品权重"`
}
type FoodCategory struct {
	gorm.Model
	Id        int      `gorm:"type:int;not null;auto increment;column:id" json:"id"`
	FoodName  string   `gorm:"type:varchar(20);column:food_name" json:"food_name"  label:"菜品名称"`
	Price     string   `gorm:"type:varchar(20)" json:"price" label:"菜品价格"`
	Cid       int      `gorm:"type:int;column:cid" json:"cid" label:"菜品分类cid"`
	FoodDesc  string   `gorm:"type:varchar(100);column:food_desc" json:"food_desc" label:"菜品描述"`
	FoodImg   string   `gorm:"type:varchar(100);column:food_img" json:"food_img" label:"菜品图片"`
	Weight    float64  `gorm:"type:double;column:weight" json:"weight" label:"菜品权重"`
	SaleCount int      `gorm:"type:int;column:sale_count" json:"saleCount" label:"菜品权重"`
	Categorys Category `gorm:"foreignKey:cid" json:"categoryInfo"`
}

// 根据分页获取用户信息
func (Food) GetListByPage(pageIndex int, foodName, categoryName string) ([]FoodCategory, int64) {
	var data []FoodCategory
	var total int64
	err := Db.
		Joins("Categorys").
		Where("food_name LIKE ?", "%"+foodName+"%").
		Where("category_name LIKE ?", "%"+categoryName+"%").
		Limit(10).Offset((pageIndex - 1) * 10).Find(&data).Count(&total).Error
	if err != nil {
		fmt.Println("Select err:", err)
		return nil, 0
	}
	return data, total
}

// 设置菜品权重
func (Food) SetFoodInfoWeight(id, weight int) (string, error) {
	err := Db.Model(&Food{}).Where("id = ?", id).Update("weight", weight).Error
	if err != nil {
		return "修改失败", err
	}
	return "修改成功", err
}

// 新增菜品
func (Food) Add(food Food) (Food, error) {
	err := Db.Create(&food).Error
	if err != nil {
		return Food{}, err
	}
	return food, err
}

// 根据id删除菜品
func (Food) DeleteById(id int) (string, error) {
	err := Db.Delete(&Food{Id: id}).Error
	if err != nil {
		return "删除失败", err
	}
	return "删除成功", err
}

// 根据id查找菜品
func (Food) FindById(id int) (Food, error) {
	var data Food
	err := Db.Where("id=?", id).Find(&data).Error
	if err != nil {
		fmt.Println("FindById err:", err)
	}
	return data, err
}

// 编辑菜品
func (Food) Update(food *Food) (string, error) {
	//fmt.Println("food_created_at:   ", food.CreatedAt)
	//food.CreatedAt = time.Now()
	err := Db.Model(&Food{}).Where("id = ?", food.Id).Updates(Food{FoodName: food.FoodName, Price: food.Price, Cid: food.Cid, FoodImg: food.FoodImg, FoodDesc: food.FoodDesc, Weight: food.Weight}).Error
	if err != nil {
		return "修改失败", err
	}
	return "修改成功", err
}
