package routes

import (
	"ACAT/api"
	"ACAT/middleware"
	"ACAT/utils"
	"github.com/gin-gonic/gin"
)
func InitRouter() {
	 r := gin.Default()
	r.Use(middleware.CORS())
		r.POST("api/admin/login",api.Login)
		r.POST("api/admin/regist",api.Regist)

	 router:=r.Group("api/admin")//管理员
	 router.Use(middleware.JwtToken())
	 {
	 	 router.GET("gointerview/:student_num",api.Gointerview)// 点击去面试，修改面试状态，正在面试
		 router.POST("appraise/:student_num",api.Appraise)	 // 评价
		 router.GET("uninterview",api.UnInterview)		//获取未面试
		 router.GET("interviewing",api.Interviewing)		//获取面试中的
		 router.GET("interviewed",api.Interviewed)	//获取已面试的
		 router.GET("passinterview",api.Passinterview)	//获取通过面试的
		 router.GET("successemail/:student_num",api.Successemail)	//发送面试成功的邮件
		 router.GET("failemail/:student_num",api.Failemail)	//发送面试失败的邮件
		 router.GET("getappraise/:student_num",api.Getappraise)	 //查看单个用户的评价信息

		 superadmin:=router.Group("superadmin")//超级管理员
		 superadmin.Use(middleware.Checkpower())
		 	{
				 superadmin.GET("getadminpage",api.Getadminpage)	//获取所有管理员界面
				 superadmin.POST("changepower",api.Changepower)		//修改权限
			 }
	 }

		_=r.Run(utils.HttpPort)
}