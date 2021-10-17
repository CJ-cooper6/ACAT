package server

import (
	"ACAT/dao"
	"ACAT/errmsg"
	"ACAT/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Interviewpage(c *gin.Context,interview_state int){
	var admin model.Admin
	var(
		list []model.Interview
		total int64
		code int
	)
	n,_:=c.Get("admin_num")
	admin.Admin_Num=n.(string)
	admin.Role=dao.CheckRole(admin.Admin_Num)	//获得角色
	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}

	switch admin.Role {
	case 0:
		list,code,total=nil,errmsg.Success,0
	default:
		list,code,total=dao.GetInterview(admin.Role,pageSize,pageNum,interview_state)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"list":list,
			"total":total,
			"message": errmsg.Geterrmsg(code),
		},
	)

}
