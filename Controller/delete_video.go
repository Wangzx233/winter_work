package Controller

import (
	"1/Model"
	"github.com/gin-gonic/gin"
)

func DeleteVideo(c *gin.Context)  {
	vid := c.Query("vid")
	uid,err := c.Request.Cookie("uid");if err!=nil{
		c.JSON(400,err)
		return
	}

	err = Model.DeleteVideo(uid.Value,vid)
	;if err!=nil{
		c.JSON(400,err)
		return
	}
	c.JSON(200,"success")
}