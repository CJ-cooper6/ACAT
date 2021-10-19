package errmsg

// 错误信息

const(
	Success=200
	Error=500

	//code =1000...管理员模块的错误

	Error_Adminname_Used=1001
	Error_AdminPassword_Wrong=1002
	Error_Admin_No_Right=1003
	Error_Token_Exist=1004
	Error_Token_Runtime=1005
	Error_Token_Wrong=1006
	Error_Token_Type_Wrong=1007
	Error_No_Power = 1008
	Error_Fail_Email=1009

)

var codemsg =map[int]string{
	Success: "OK",
	Error: "Fail",
	Error_Adminname_Used: "管理员用户名已存在",
	Error_AdminPassword_Wrong: "密码错误",
	Error_Admin_No_Right: "管理员不存在，该学号未注册",
	Error_Token_Exist: "Token不存在",
	Error_Token_Runtime: "Token已过期",
	Error_Token_Wrong: "Token不正确",
	Error_Token_Type_Wrong: "Token格式错误",
	Error_No_Power:"没有权限",
	Error_Fail_Email:"发送邮箱失败",
}

func Geterrmsg(code int)string{

	return codemsg[code]
}
