package Controller

import (
	"1/Model"
	"github.com/gin-gonic/gin"
)

func FindPicture(c *gin.Context)  {
	uid, err := c.Request.Cookie("uid")
	if err == nil {
		res,err := Model.UserInfo(uid.Value);if err!=nil{
			c.JSON(6001,err)
			return
		}
		c.JSON(200,res.PictureURL)
	} else {
		c.JSON(200,"https://static.hdslb.com/images/akari.jpg")//未登录头像
	}
}
