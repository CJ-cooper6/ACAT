package api

import (
	"ACAT/server"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Successemail(c *gin.Context){
	student_num,_:= strconv.Atoi(c.Param("student_num"))
	server.Sentemail(c,student_num,true)

}
func Failemail(c *gin.Context){
	student_num,_:= strconv.Atoi(c.Param("student_num"))
	server.Sentemail(c,student_num,false)
}
