package api

import (
	"ACAT/dao"
	"ACAT/errmsg"
	"ACAT/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Appraise(c *gin.Context){		//评价
	var interview model.Interview
	//获取被评价者学号
	student_num,_:= strconv.Atoi(c.Param("student_num"))
	//获取评价,评价人,分数
	c.ShouldBind(&interview)
	//保存评价
dao.SaveAppraise(student_num,&interview)
	//更改面试状态
code:=dao.ChangeInterview_State(student_num,2)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.Geterrmsg(code),
		},
	)

}

func Getappraise(c *gin.Context){
	student_num,_:= strconv.Atoi(c.Param("student_num"))
	view,code:=dao.Getappraise(student_num)
	c.JSON(
		http.StatusOK, gin.H{
			"view":view,
			"status":  code,
			"message": errmsg.Geterrmsg(code),
		},
	)
}