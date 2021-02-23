package Controller

import (
	"1/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchVideo(c *gin.Context) {
	keyword := c.Query("keyword")

	err,res	:= Model.SearchVideo(keyword);if err!=nil{
		c.String(http.StatusBadRequest, "Error:%s", err.Error())
		return
	}

	c.JSON(200,res)
}
