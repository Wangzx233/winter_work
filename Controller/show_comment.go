package Controller

import (
	"1/Model"
	_struct "1/struct"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ShowComment(c *gin.Context)  {
	vid := c.Query("vid")

	intid, err := strconv.Atoi(vid);if err!=nil{
		c.String(http.StatusBadRequest, "Error:%s", err.Error())
		return
	}

	mp := []_struct.Comment{}
	err,mp = Model.ShowComment(intid);if err!=nil{
		c.String(http.StatusBadRequest, "Error:%s", err.Error())
		return
	}

	c.JSON(200,mp)


}
