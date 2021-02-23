package Controller

import (
	"1/Model"
	_struct "1/struct"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ShowVideo(c *gin.Context) {
	id := c.Query("id")

	intid, err := strconv.Atoi(id);if err!=nil{
		c.String(http.StatusBadRequest, "Error:%s", err.Error())
		return
	}

	v,err := Model.ShowVideo(intid);if err!=nil{
		c.String(http.StatusBadRequest, "Error:%s", err.Error())
		return
	}

	c.JSON(200,v)

}

func VideoInfo(c *gin.Context)  {
	var video_info _struct.VideoInfo
	id := c.Query("id")
	intid, err := strconv.Atoi(id);if err!=nil{
		c.JSON(600, err.Error())
		return
	}

	video_info,err = Model.VideoInfo(intid);if err!=nil{
		c.JSON(601, err.Error())
		return
	}
	c.JSON(200,video_info)
}

func OnlineAdd(c *gin.Context)  {
	vid := c.Query("vid")
	Model.OnlineAdd(vid)
}

func OnlineSub(c *gin.Context)  {
	vid := c.Query("vid")
	Model.OnlineSub(vid)
}

func UserVideo(c *gin.Context) {
	uid := c.Query("uid")

	err,res	:= Model.UserVideo(uid);if err!=nil{
		c.String(http.StatusBadRequest, "Error:%s", err.Error())
		return
	}

	c.JSON(200,res)

}