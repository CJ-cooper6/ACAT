package api

import (
	"ACAT/dao"
	"ACAT/errmsg"
	"ACAT/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)



//func Getmainpage(c *gin.Context){	//TODO 主页面
//	var admin model.Admin
//	 var(
//		 list []model.Interview
//		 total int64
//	 	 code int
//	 )
//
//	n,_:=c.Get("user_num")
//	admin.Admin_Num=n.(string)
//	admin.Role=dao.CheckRole(admin.Admin_Num)	//获得角色
//
//	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
//	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
//	if pageSize == 0{
//		pageSize = -1
//	}
//	if pageNum == 0{
//		pageNum = -1
//	}
//
//	switch admin.Role {
//	case 0:
//		list,code,total=nil,errmsg.Success,0
//	case 1:	//返回所有学生信息
//		list,code,total=dao.Getmainpage(pageSize,pageNum)
//	default:		//返回相对应方向的学生信息
//		list,code,total=dao.Getappoint(admin.Role,pageSize,pageNum)
//	}
//	c.JSON(
//		http.StatusOK, gin.H{
//			"status":  code,
//			"list":list,
//			"total":total,
//			"message": errmsg.Geterrmsg(code),
//		},
//	)
//}

func UnInterview(c *gin.Context){	//未面试的

server.Interviewpage(c,0)
}

func Interviewing(c *gin.Context){	//面试中的
	server.Interviewpage(c,1)
}

func Interviewed(c *gin.Context){	//完成面试的
	server.Interviewpage(c,2)

}

func Getadminpage(c *gin.Context){	//获取管理员的页面

	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	admins,code,total:=dao.Getadminpage(pageSize,pageNum)
	c.JSON(
		http.StatusOK, gin.H{
			"power":true,
			"status":  code,
			"admins":admins,
			"total":total,
			"message": errmsg.Geterrmsg(code),
		},
	)

}
