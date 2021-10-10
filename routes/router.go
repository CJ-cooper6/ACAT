package routes

import (
	"ACAT/api"
	"ACAT/middleware"
	"ACAT/utils"
	"github.com/gin-gonic/gin"
)
func InitRouter() {

	 r := gin.Default()

	//todo 未面试
	//todo 面试中
	//todo 完成面试
	//todo 发送成功的邮件
	//todo 发送失败的邮件
		r.POST("/admin/login",api.Login)
		r.POST("/admin/regist",api.Regist)


	 router:=r.Group("admin")//管理员
	 router.Use(middleware.JwtToken())
	 {

		 router.POST("appraise/:student_num",api.Appraise)	 // 评价
		 router.GET("uninterview",api.UnInterview)		//获取未面试
		 router.GET("interviewing",api.Interviewing)		//获取面试中的
		 router.GET("interviewed",api.Interviewed)	//获取已面试的

		 router.POST("successemail/:student_num",api.Successemail)
		 router.POST("failemail/:student_num",api.Failemail)


		 superadmin:=router.Group("superadmin")//超级管理员
		 superadmin.Use(middleware.Checkpower())
		 	{
				 superadmin.GET("getadminpage",api.Getadminpage)	//获取所有管理员界面
				 superadmin.POST("changepower",api.Changepower)		//修改权限
			 }


	 }
		r.Run(utils.HttpPort)
}