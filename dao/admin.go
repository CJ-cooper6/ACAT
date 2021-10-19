package dao

import (
	"ACAT/errmsg"
	"ACAT/model"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

func CheckAdminNum(num string)int{	//检查学号是否注册过
	var admin model.Admin
	if errors.Is(model.Db.Where("admin_num=?",num).First(&admin).Error,gorm.ErrRecordNotFound)==true{
		return errmsg.Success
	}
	return errmsg.Error_Adminname_Used
}

func SaveAdmin(admin *model.Admin){	//注册管理员信息
	admin.Admin_Password=ScryotPW(admin.Admin_Password)
	err:=model.Db.Create(admin).Error
	if err!=nil{
		fmt.Println(err)
	}
}



func CheckLogin(admin_num string,password string)(int,string){	//检查学号密码是否正确
	var admin model.Admin
	model.Db.Where("admin_num = ?",admin_num).First(&admin)
	if admin.ID == 0{
		return errmsg.Error_Admin_No_Right,""
	}
	if ScryotPW(password)!=admin.Admin_Password{
		return errmsg.Error_AdminPassword_Wrong,""
	}

	return errmsg.Success,admin.Admin_Name
}

func ScryotPW(password string)string{	//密码加密
	const Kenlen =10
	salt := make([]byte,8)
	salt = []byte{12,32,4,6,66,22,222,12}
	HashPW,err:=scrypt.Key([]byte(password),salt,16384,8,1,Kenlen)
	if err != nil{
		log.Fatal(err)
	}
	fpw:=base64.StdEncoding.EncodeToString(HashPW)
	return fpw
}

func CheckRole(admin_num string)int{
	var admin model.Admin
	model.Db.Where("admin_num = ?",admin_num).First(&admin)

return admin.Role

}





