# 西安邮电大学-ACAT计算机应用技术协会-纳新后台系统后端部分



### 介绍

这是西安邮电大学2021年纳新系统后台，旨在更好的查看、管理学生信息。

与小伙伴合作完成，前端 Vue   后端Golang，这里仅介绍后端部分。



### 目录结构

```shell
├─api	 		(api层)
├─config		(配置包)
├─dao			(数据库)
├─dist			(Vue打包文件)
├─errmsg		(错误处理)
├─log			(日志文件)
├─middleware	        (中间件层)  
├─model		        (模型层)
├─routes		(路由层) 
├─server		(部分功能函数)
└─utils			(初始化内容)

```



### 实现功能

1. 用户密码加密存储
2. JWT认证
3. 基于Casbin的权限管理
4. 列表分页
5. 自定义日志功能
6. 发送邮件



### 技术栈

- Golang
  - Gin
  - jwt-go
  - go-mail
  - go-ini
  - logrus
  - Casbin
- MYSQL
- Nginx



### 项目预览

- 登录界面

![image-20211017163428480](https://gitee.com/CJ-cooper6/picgo/raw/master/image-20211017163428480.png)

### 部分接口文档

https://www.showdoc.com.cn/1637571656793053



### 运行&部署

###### **前端部署**

登录实验室服务器，在/etc/nginx目录下修改nginx.conf文件

![image-20211017164119201](https://gitee.com/CJ-cooper6/picgo/raw/master/image-20211017164119201.png)

根据需要选择自己的端口

**注意**：要做Frp内网端口映射，否则无法使用



将前端打包好的文件dist位置要和配置文件中root位置相同



###### **后端**

1. 克隆项目

2. 转到下面文件夹下

3. 安装依赖

4. 初始化项目配置config.ini

   ```go
   ACAT/config/config.ini
   
   [server]
   #debug 开发模式 ， release 生产模式
   AppMode = debug
   JwtKey = 89js82js72	#JWT密钥，随机字符串即可
   HttpPort = :7301	# 项目端口
   
   
   
   [database]
   Db = mysql	#数据库类型，不能变更为其他形式
   DbHost = localhost	# 数据库地址
   Dbport = 3306	# 数据库端口
   Dbuser = root	 # 数据库用户名
   DbPassWord=		# 数据库用户密码
   DbName = ACAT	ginblog # 数据库名
   CasbinDbName = ACAT		#存放Casbin策略的数据库名
   ```

5. 创建数据库

   ```
   create database ACAT;
   ```

   

6. 启动项目

```shell
 go run main.go
  或者编译好后台不挂断运行
  go build
 sudo nohup ACAT  > nohup_bluebell.log 2>&1 &	
 
 //> nohup_bluebell.log表示将命令的标准输出重定向到 nohup_bluebell.log 文件
 
 //2>&1表示将标准错误输出也重定向到标准输出中，结合上一条就是把执行命令的输出都定向到 nohup_bluebell.log 文件
 
```



此时，项目启动，你可以访问页面

```
后台首页
http://192.168.1.197:7300

默认管理员：admin123  密码：123456
```



### 不足

数据库设计感觉不太好，重复字段过多，之后应使用联结或合成一张表



## Thanks for free JetBrains Open Source license

感谢JetBrains免费开源授权

<img src="https://gitee.com/wejectchan/ginblog/raw/master/upload/jet.png" alt="img" style="zoom: 25%;" />

