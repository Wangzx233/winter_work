package Controller

import (
	"1/Model"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context)  {
	uid := c.Query("uid")

	u,err := Model.UInfo(uid);if err!=nil{
		c.JSON(500,"bind_err")
		return
	}
	c.JSON(200,u)
}
