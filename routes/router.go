package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orderSystem/handlers"
)

func InitRouter() {
	r := gin.Default()
	//r.SetTrustedProxies([]string{"127.0.0.1"})
	//r.LoadHTMLGlob("./static/pages/*")
	r.StaticFS("/src", http.Dir("./web"))
	address := handlers.Address{}
	addressG := r.Group("/addressInfo")
	{
		addressG.GET("/getListByPage", address.GetListByPage)
		addressG.GET("/deleteById", address.DeleteById)
	}
	admin := handlers.Admin{}
	adminG := r.Group("/adminInfo")
	{
		adminG.POST("/checkLogin", admin.CheckLogin)
		adminG.GET("/getListByPage", admin.GetListByPage)
		adminG.GET("/deleteById", admin.DeleteById)
		adminG.GET("/findById", admin.FindById)
		adminG.POST("/update", admin.Update)
		adminG.POST("/add", admin.Add)
		adminG.GET("/getAllList", admin.GetAllList)
		adminG.GET("/getTotalInfo", admin.GetTotalInfo)
		adminG.GET("/getCategoryFoodCount", admin.GetCategoryFoodCount)
		adminG.POST("/uploadAdminPhoto", admin.UploadAdminPhoto)
	}
	category := handlers.Category{}
	categoryG := r.Group("/categoryInfo")
	{
		categoryG.GET("/getListByPage", category.GetListByPage)
		categoryG.GET("/add", category.Add)
		categoryG.GET("/getAllList", category.GetAllList)
		categoryG.GET("/deleteById", category.DeleteById)
		categoryG.GET("/findById", category.FindById)
		categoryG.POST("/update", category.Update)

	}
	comment := handlers.Comment{}
	commentG := r.Group("/commentInfo")
	{
		commentG.GET("/getListByPage", comment.GetListByPage)
		commentG.GET("/deleteById", comment.DeleteById)
		commentG.GET("/setCommentShowById", comment.SetCommentShowById)
		commentG.GET("/setCommentHideById", comment.SetCommentHideById)
	}
	food := handlers.Food{}
	foodG := r.Group("/foodInfo")
	{
		foodG.GET("/getListByPage", food.GetListByPage)
		foodG.GET("/setFoodInfoWeight", food.SetFoodInfoWeight)
		foodG.POST("/add", food.Add)
		foodG.GET("/deleteById", food.DeleteById)
		foodG.GET("/findById", food.FindById)
		foodG.POST("/update", food.Update)
		foodG.POST("/uploadFoodImg", food.UploadFoodPhoto)
	}
	order := handlers.Order{}
	orderG := r.Group("/orderInfo")
	{
		orderG.GET("/getListByPage", order.GetListByPage)
		orderG.GET("/updateOrderToDeliver", order.UpdateOrderToDeliver)
	}
	user := handlers.User{}
	userG := r.Group("/userInfo")
	{
		userG.GET("/getListByPage", user.GetListByPage)
		userG.GET("/deleteById", user.DeleteById)
	}
	r.Run(":8080")

}
