package api

import (
	"ACAT/dao"
	"ACAT/errmsg"
	"ACAT/model"
	"ACAT/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Changepower(c *gin.Context){	//T 赋权
	var admin model.Admin
	c.ShouldBind(&admin)

	role:=strconv.Itoa(admin.Role)

	code:=dao.Changepower(admin.Admin_Num,admin.Role)	//先删除casbin表中的策略，更新admin表中的role
	_=utils.E.LoadPolicy()	//重新从数据库中加载策略
	add,_:=utils.E.AddGroupingPolicy(admin.Admin_Num,role)	//增加策略

		fmt.Println(add)


	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.Geterrmsg(code),
		},
	)

}
