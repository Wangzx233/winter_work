package Controller

import (
	"1/Model"
	_struct "1/struct"
	"github.com/gin-gonic/gin"
)

func SentBarrage(c *gin.Context)  {
	uid,err := c.Request.Cookie("uid");if err!=nil{
		c.JSON(400,err)
		return
	}
	var ba _struct.Barrage
	ba.Uid=uid.Value

	err = c.ShouldBind(ba);if err!=nil{
		c.JSON(400,err)
		return
	}

	err = Model.SentBarrage(ba);if err!=nil{
		c.JSON(400,err)
		return
	}
}

func ShowBarrage(c *gin.Context)  {
	vid := c.Query("vid")

	err,res := Model.ShowBarrage(vid);if err!=nil{
		c.JSON(400,err)
		return
	}

	c.JSON(200,res)
}
