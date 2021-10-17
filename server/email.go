package server

import (
	"ACAT/dao"
	"ACAT/errmsg"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"net/http"
)

func Sentemail(c *gin.Context,student_num int,flag bool)bool {
	var code string //邮件内容

	//获取用户邮箱
	email := dao.Checkemail(student_num)

	//发送邮件
	if flag == true { //发送成功邮件
		code = "成功"
	} else { //发送失败邮件
		code = "失败"
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "473172339@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "ACAT面试") // 邮件标题
	m.SetBody("text/html", code)     // 邮件内容
	d := gomail.NewDialer("smtp.qq.com", 465, "473172339@qq.com", "sredxirbwncccade")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("***%s\n", err.Error())
		c.JSON(
			http.StatusOK, gin.H{
				"status":  errmsg.Error_Fail_Email,
				"message": errmsg.Geterrmsg(errmsg.Error_Fail_Email),
			},
		)
		return false
	} else {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  200,
				"message": errmsg.Geterrmsg(200),
			},
		)
	}
	return true
}
