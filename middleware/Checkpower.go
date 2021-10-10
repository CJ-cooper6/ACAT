package middleware

import (
	"ACAT/errmsg"
	"ACAT/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Checkpower()gin.HandlerFunc{		//检查是否有超级管理员的权限
	return func(c *gin.Context){

		admin_num,_:=c.Get("admin_num")
		_=utils.E.LoadPolicy()	//重新从数据库中加载策略
		ok, _:= utils.E.Enforce(admin_num,c.Request.URL.Path,"change")
		if ok == true {
			// 有权限

			c.Next()
		} else {
			// 拒绝请求，抛出异常
			c.JSON(
				http.StatusOK, gin.H{
					"power":false,
					"status":  errmsg.Error_No_Power,
					"message": errmsg.Geterrmsg(errmsg.Error_No_Power),
				},
			)
			c.Abort()
			return
		}

	}
}
