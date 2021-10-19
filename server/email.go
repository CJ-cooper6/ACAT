package server

import (
	"ACAT/dao"
	"ACAT/errmsg"
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"html/template"
	"net/http"
)

const Successmail = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body>
    <div>
    <div class="post" style="width: 400px; height: 300px; margin: 0 auto; color: rgba(233, 126, 107, 0.5); border: 1px solid rgba(233, 118, 98, 0.5); box-shadow: 5px 5px 20px;">
        <p href="javascript:;" style="color: rgba(233, 126, 107); position :relative; top:45px; left:50px; text-decoration:none ;font-size:20px; font-weight:800;">亲爱的 {{.Name}}同学，你好！</p>
        <div class="content" style=" width: 350px; height: 150px;position: relative; top: 70px;left: 25px;word-break: break-all;">
            <p>&nbsp;&nbsp;&nbsp;&nbsp;恭喜你通过了ACAT一面，接下来进入二面。请加入qq群号： {{.Group}}</p>
            <p>&nbsp;&nbsp;&nbsp;&nbsp;我们期待你优秀的表现!</p>
            <p class="address" style="position: relative; top: 0; left: 100px;">——ACAT计算机应用技术协会</p>
        </div>
    </div>
</div>
</body>

</html>`
const Failmail =`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>

</head>

<body>
    <div>
        <div class="post" style="width: 400px; height: 300px; margin: 0 auto; color: rgba(87, 80, 79, 0.5); border: 1px solid rgba(87, 80, 79, 0.5); ; box-shadow: 5px 5px 20px;">
            <p href="javascript:;" style="color: rgb(83, 83, 83); position :relative; top:45px; left:50px; text-decoration:none ;font-size:20px; font-weight:800;">亲爱的{{.Name}}同学，你好！</p>
            <div class="content" style=" width: 350px; height: 150px;position: relative; top: 70px;left: 25px;word-break: break-all;">
                <p>&nbsp;&nbsp;&nbsp;&nbsp;结合你整体情况，经过慎重考虑，很遗憾你没有通过ACAT一面。</p>
                <p>&nbsp;&nbsp;&nbsp;&nbsp;别灰心，待你准备好ACAT随时欢迎你来霸面。</p>
                <p class="address" style="position: relative; top: 0; left: 100px;">——ACAT计算机应用技术协会</p>
            </div>
        </div>
    </div>
</body>

</html>
`

var QQgroup =map[int]string{	//各组的二面qq群
	2:		"755890635",	//前端
	3:		"163486972",	//Go
	4:		"935462280",	//Java
	5:		"664173684",	//服务端
	6:		"537460758",	//机器学习
}

type News struct {
	Name string
	Group string 	//qq群
}

func Sentemail(c *gin.Context,student_num int,flag bool)bool {
	var N News
	//获取学生姓名，邮箱,方向
	email,name,role:= dao.Checknameemail(student_num)
	t := template.New("Mail") //创建一个模板
	buffer := new(bytes.Buffer)
	if flag == true { //发送成功邮件
		N=News{name,Getgroup(role)}
		t, _ = t.Parse(Successmail)
	} else { // 发送失败邮件
		N=News{name,""}
		t, _ = t.Parse(Failmail)
	}

	_=t.Execute(buffer,N)	//将加载完的html放到buffer中
	m := gomail.NewMessage()
	m.SetHeader("From", "473172339@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "ACAT") // 邮件标题
	m.SetBody("text/html", buffer.String())     // 邮件内容,从buffer中获取
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

func Getgroup(role int)string{
	return QQgroup[role]
}