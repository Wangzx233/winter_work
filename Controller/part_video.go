package Controller

import (
	"1/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindVideo(c *gin.Context)  {
	part := c.Query("part")

	err,res	:= Model.FindVideo(part);if err!=nil{
		c.String(http.StatusBadRequest, "Error:%s", err.Error())
		return
	}

	c.JSON(200,res)
}