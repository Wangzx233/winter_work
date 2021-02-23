package Controller

import (
	"1/Model"
	"github.com/gin-gonic/gin"
)
//点赞
func Like(c *gin.Context)  {
	vid := c.Query("vid")

	uid, err := c.Request.Cookie("uid");if err!=nil{
		c.JSON(400,err)
		return
	}

	if !Model.Like(uid.Value,vid){
		c.JSON(400,"已点赞")
	}
}

//投币
func Coin(c *gin.Context)  {
	 vid := c.Query("vid")
	 to_uid := c.Query("uid")
	 uid, err := c.Request.Cookie("uid");if err!=nil{
		c.JSON(400,err)
		return
	 }

	Model.Coin(uid.Value,vid,to_uid)
}

//点赞
func Collection(c *gin.Context)  {
	vid := c.Query("vid")

	uid, err := c.Request.Cookie("uid");if err!=nil{
		c.JSON(400,err)
		return
	}

	if !Model.Collection(uid.Value,vid){
		c.JSON(400,"已收藏")
	}
}
