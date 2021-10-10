package dao

import (
	"ACAT/errmsg"
	"ACAT/model"
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)
//	修改权限

func Changepower(admin_num string,role int)int{
	var admin model.Admin
	var cas gormadapter.CasbinRule

	err:=model.Db.Where("v0=?",admin_num).Delete(&cas).Error
	if err != nil{
		fmt.Println(err)
	}
	err=model.Db.Model(&admin).Where("admin_num = ?", admin_num).Update("role", role).Error
	if err!=nil{
		fmt.Println(err)
		return errmsg.Error
	}
	return errmsg.Success
}
