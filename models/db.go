package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB
var err error

func (Address) TableName() string {
	return "address"
}
func (Admin) TableName() string {
	return "admin"
}
func (Category) TableName() string {
	return "category"
}
func (Comment) TableName() string {
	return "comment"
}
func (Food) TableName() string {
	return "food"
}
func (Order) TableName() string {
	return "order"
}
func (User) TableName() string {
	return "user"
}
func (FoodCategory) TableName() string {
	return "food"
}
func (AddressUser) TableName() string {
	return "address"
}
func (OrderAddressUser) TableName() string {
	return "order"
}
func InitDb() {

	dsn := "root:Li20031202@tcp(127.0.0.1:3306)/order?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
	}

	sqlDB, err := Db.DB()
	if err != nil {
		fmt.Println("err::", err)
		return
	}
	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	//_ = Db.AutoMigrate(&Address{}, &Admin{}, &Category{}, &Comment{}, &Food{}, &Order{}, &User{})
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
