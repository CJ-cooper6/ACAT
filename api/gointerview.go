package api

import (
	"ACAT/dao"
	"ACAT/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Gointerview(c *gin.Context){
	student_num,_:= strconv.Atoi(c.Param("student_num"))

	code:=dao.ChangeInterview_State(student_num,1)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.Geterrmsg(code),
		},
	)

}
