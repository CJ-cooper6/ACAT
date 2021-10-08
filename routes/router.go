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


	 router:=r.Group("admin")//管理员
	 router.Use(middleware.JwtToken())
	 {
		 router.GET("getmainpage",api.Getmainpage)	//获取主页面
		 router.POST("appraise/:student_num",api.Appraise)	 // 评价

		 superadmin:=router.Group("superadmin")//超级管理员

		 superadmin.Use(middleware.Checkpower())
		 {
			 superadmin.GET("getadminpage",api.Getadminpage)	//获取所有管理员界面
			 superadmin.POST("changepower",api.Changepower)		//修改权限
		 }


	 }
		r.Run(utils.HttpPort)
}