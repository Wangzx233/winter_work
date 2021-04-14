package Controller

import (
	"1/Model"
	_struct "1/struct"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(context *gin.Context)  {
	var user _struct.UserLogin
	err:=context.ShouldBind(&user)
	if err!=nil{
		context.JSON(800,"bind_err")
		return
	}
	res,rs:=Model.Register(user.UserName,user.Password)
	if res {
		context.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": rs,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": rs,
		})
	}
}

func Login(context *gin.Context) {
	var user _struct.UserLogin
	err:=context.ShouldBind(&user)
	if err!=nil{
		context.JSON(900,err)
		return
	}
	res,uid:=Model.Login(user.UserName,user.Password)
	if res{
		//context.SetCookie("uid",user.UserName, 9999, "/", "/", false, false)

		//jwt
		jwt := NewJWT(uid)
		bytes, err := json.Marshal(jwt)
		if err!=nil {
			log.Println(err)
		}
		context.Header("Authorization","Bearer token="+string(bytes))

		context.JSON(200,gin.H{
			"code":		200,
			"message":	"你好"+user.UserName,
		})

	}else {
		context.JSON(200,gin.H{
			"code":		20000,
			"message":	"用户名或密码错误",
		})
	}
}
func Cancel(context *gin.Context)  {
	context.SetCookie("uid","",-1, "/", "/", false, false)
}

func IsLogin(c *gin.Context)  {
	//_, err := c.Request.Cookie("uid");if err==nil{
	//	c.Redirect(http.StatusMovedPermanently, "http://115.29.201.183:443/")
	//}
	_, exists := c.Get("jwt")
	if exists {
		c.JSON(200,gin.H{
			"code":		200,
			"message":	"已登录",
		})
	}else {
		c.JSON(200,gin.H{
			"code":		200,
			"message":	"未登录",
		})
	}
}