package api

import (
	"ACAT/dao"
	"ACAT/errmsg"
	"ACAT/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UnInterview(c *gin.Context){	//未面试的

server.Interviewpage(c,0)
}

func Interviewing(c *gin.Context){	//面试中的
	server.Interviewpage(c,1)
}

func Interviewed(c *gin.Context){	//完成面试的
	server.Interviewpage(c,2)

}

func Passinterview(c *gin.Context){	// 返回通过面试的
server.Interviewpage(c,3)
}


func Getadminpage(c *gin.Context){	//获取管理员的页面

	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))

	n,_:=c.Get("admin_num")
	admin_Num:=n.(string)

	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	admins,code,total:=dao.Getadminpage(pageSize,pageNum,admin_Num)
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
