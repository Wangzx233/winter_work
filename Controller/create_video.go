package Controller

import (
	"1/Model"
	_struct "1/struct"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func CreatVideo(context *gin.Context)  {
	var video _struct.Video
	err := context.ShouldBind(&video);if err!=nil{
		context.String(5000,"绑定失败")
		return
	}

	//读取上传的视频
	v,err :=context.FormFile("video")
	if err!=nil{
		context.String(5000,"文件读取失败")
		return
	}
	videoName := v.Filename
	//读取上传的封面
	cover,err :=context.FormFile("cover")
	if err!=nil{
		context.String(5000,"文件读取失败")
		return
	}
	coverName := cover.Filename

	//获取当前用户cookie
	username,err := context.Request.Cookie("uid")
	if err!=nil{
		context.String(5001,"获取cookie失败")
		return
	}
	uid, err := strconv.Atoi(username.Value);if err!=nil{
		context.String(http.StatusBadRequest, "Error:%s", err.Error())
		return
	}
	video.UID=uid

	//保存视频
	url := "./"+video.Part+"/"+username.Value+"/"+videoName
	fmt.Println(url)

	//存入数据库
	video.VideoURL=url+"/"+videoName
	video.CoverURL=url+"/"+coverName
	err = Model.CreateVideo(video);if err!=nil{
		context.String(http.StatusBadRequest, "数据写入失败 Error:%s", err.Error())
		return
	}

	//继续保存视频
	err=os.MkdirAll(url,0777);if err!=nil{
		context.String(http.StatusBadRequest, "文件夹创建失败 Error:%s", err.Error())
	}

	if err := context.SaveUploadedFile(v, url+"/"+videoName); err != nil {
		context.String(http.StatusBadRequest, "视频保存失败 Error:%s", err.Error())
		return
	}
	if err := context.SaveUploadedFile(cover, url+"/"+coverName); err != nil {
		context.String(http.StatusBadRequest, "封面保存失败 Error:%s", err.Error())
		return
	}

	context.JSON(http.StatusOK, "上传文件成功")
}