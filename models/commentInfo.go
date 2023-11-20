package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Uid       int       `gorm:"type:int;not null;auto increment" json:"id" label:"用户id"`
	StartTime time.Time `gorm:"type:datetime" json:"start_time" label:"订单开始时间"`
	EndTime   time.Time `gorm:"type:datetime" json:"end_time" label:"订单结束时间"`
}

/**
 *
 * @param {*} pageIndex  评论页码
 * @param {*} food_name  菜品名称
 * @param {*} start_time  订单开始时间
 * @param {*} end_time 订单结束时间
 * @returns  {Promise<AxiosResponse>}
 */
// 根据分页获取菜品评论分类信息
func (comment Comment) getListByPage(pageIndex, food_name, start_time, end_time int) {}

// 根据id删除评论
func (comment Comment) deleteById(id int) {}

// 根据id显示评论
func (comment Comment) setCommentShowById(id int) {}

// 根据id隐藏评论
func (comment Comment) setCommentHideById(id int) {}
