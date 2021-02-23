package Controller

import (
	"1/Model"
	_struct "1/struct"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context)  {
	var com _struct.Comment

	err := c.ShouldBind(&com);if err!=nil{
		c.JSON(400,err)
		return
	}


	ck, err := c.Request.Cookie("uid");if err!=nil{
		c.JSON(400,err)
		return
	}

	//ick, err := strconv.Atoi(ck.Value);if err!=nil{
	//	c.String(http.StatusBadRequest, "Error:%s", err.Error())
	//	return
	//}
	com.Uid = ck.Value




	err = Model.CreateComment(com)
	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}
}
