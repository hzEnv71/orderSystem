package handlers

import (
	"github.com/gin-gonic/gin"
	"orderSystem/models"
)

type Comment models.Comment

// GetListByPage 根据分页获取菜品评论分类信息
func (comment Comment) GetListByPage(c *gin.Context) {}

// DeleteById 根据id删除评论
func (comment Comment) DeleteById(c *gin.Context) {}

// SetCommentShowById 根据id显示评论
func (comment Comment) SetCommentShowById(c *gin.Context) {}

// SetCommentHideById 根据id隐藏评论
func (comment Comment) SetCommentHideById(c *gin.Context) {}
