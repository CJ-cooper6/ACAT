package api

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func S(){
	a, _ := gormadapter.NewAdapter("mysql", "root:123@tcp(127.0.0.1:3306)/casbin",true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("./model.conf", a)

	sub := "154" // 想要访问资源的用户。
	obj :="data1" // 将被访问的资源。
	act := "read" // 用户对资源执行的操作。

	added,err:= e.AddPolicy("eve", "data3", "read")
	println(added)
	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理err
		fmt.Println("发生错误")
	}

	if ok == true {
		// 允许alice读取data1
		fmt.Println("有权限")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("没有权限")
	}
}
