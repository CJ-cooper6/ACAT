package api

import (
	"ACAT/dao"
	"ACAT/errmsg"
	"ACAT/middleware"
	"ACAT/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context){



	var admin model.Admin
	var token string
	_=c.ShouldBindJSON(&admin)
	code,admin_name:=dao.CheckLogin(admin.Admin_Num,admin.Admin_Password)

	if code==errmsg.Success{
		token,code=middleware.SetToken(admin.Admin_Num)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"admin_name":admin_name,
			"status":  code,
			"token":token,
			"message": errmsg.Geterrmsg(code),
		},
	)

}

func Regist(c *gin.Context){
	var admin model.Admin

	_=c.ShouldBind(&admin)

	code:=dao.CheckAdminNum(admin.Admin_Num)	//判断学号是否注册
	if code == errmsg.Success{
		dao.SaveAdmin(&admin)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.Geterrmsg(code),

		},
	)

}
