package api

import (
	"ACAT/dao"
	"ACAT/server"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Successemail(c *gin.Context){
	student_num,_:= strconv.Atoi(c.Param("student_num"))
	b:=server.Sentemail(c,student_num,true)
	if b == false{
		c.Abort()
		return
	}
	// 修改面试状态，已通过面试
	dao.ChangeInterview_State(student_num,3)

}
func Failemail(c *gin.Context){
	student_num,_:= strconv.Atoi(c.Param("student_num"))
	b:=server.Sentemail(c,student_num,false)
	if b == false{
		c.Abort()
		return
	}
	// 修改面试状态，未通过面试
	dao.ChangeInterview_State(student_num,4)
}
