package routes

import (
	"ACAT/api"
	"ACAT/middleware"
	"ACAT/utils"
	"github.com/gin-gonic/gin"
)
func InitRouter() {

	 r := gin.Default()

		r.POST("/admin/login",api.Login)
		r.POST("/admin/regist",api.Regist)


	 router:=r.Group("admin")
	 router.Use(middleware.JwtToken())
	 {
		 router.GET("getmainpage",api.Getmainpage)	//获取主页面

		 //todo 评价
		 router.POST("appraise/:student_num",api.Appraise)
		 //todo 权限管理

	 }


		r.Run(utils.HttpPort)
}