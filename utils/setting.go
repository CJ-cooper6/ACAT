package utils

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gopkg.in/ini.v1"
)


var(
	AppMode string
	HttpPort string

	JwtKey string
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
	CasbinDbName string



)

var E *casbin.Enforcer
func init(){
	file ,err := ini.Load("./config/config.ini")
	if err!= nil {
		fmt.Println("配置文件读取错误，检查文件路径：",err)
	}

	LoadServer(file)
	LoadDb(file)

	//casbin初始化
	dsn:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser,DbPassWord,DbHost,DbPort,CasbinDbName)
	a, _ := gormadapter.NewAdapter(Db, dsn,true)
	E, _ = casbin.NewEnforcer("./config/model.conf", a)


}

func LoadServer(file *ini.File){
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	JwtKey =file.Section("server").Key("JwtKey").MustString("89js82js72")
	HttpPort =file.Section("server").Key("HttpPort").MustString("3000")
}

func LoadDb(file *ini.File){
	Db =file.Section("database").Key("Db").MustString("mysql")
	DbHost =file.Section("database").Key("DbHost").MustString("localhost")
	DbPort=file.Section("database").Key("Dbport").MustString("3306")
	DbUser =file.Section("database").Key("Dbuser").MustString("root")
	DbPassWord=file.Section("database").Key("DbPassWord").MustString("")
	DbName =file.Section("database").Key("DbName").MustString("ACAT")
	CasbinDbName =file.Section("database").Key("CasbinDbName").MustString("casbin")
}

