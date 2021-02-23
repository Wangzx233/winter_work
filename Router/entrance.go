package Router

import (
	"1/Controller"
	"1/middleware"
	"github.com/gin-gonic/gin"
)
func Entrance() {
	r := gin.Default()

	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		//ping
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200,gin.H{
				"msg":"pong",
			})
		})

		//登录
		v1.GET("user/login",Controller.Login)
		v1.GET("user/isLogin",Controller.IsLogin)

		//第三方登录
		v1.GET("user/oauth",Controller.Oauth)

		v1.POST("user/register",Controller.Register)
		v1.GET("user/logout",middleware.Auth(),Controller.Cancel)

		//创建视频
		v1.POST("video",middleware.Auth(),Controller.CreatVideo)

		//分区视频和相关推荐
		v1.GET("video/list",Controller.FindVideo)

		//推荐视频
		v1.GET("video/hot",Controller.HotVideo)

		//删除视频
		v1.DELETE("video",middleware.Auth(),Controller.DeleteVideo)

		//视频详细页
		v1.GET("video",Controller.ShowVideo)
		v1.PUT("video/online_add",Controller.OnlineAdd)
		v1.PUT("video/online_sub",Controller.OnlineSub)
		v1.GET("video/info",Controller.VideoInfo)


		//头像获取
		v1.GET("user/head_picture",Controller.FindPicture)

		//查看评论
		v1.GET("video/show_comment",Controller.ShowComment)

		//发表评论
		v1.POST("video/create_comment",middleware.Auth(),Controller.CreateComment)

		//用户信息
		v1.GET("user/info",Controller.UserInfo)

		//搜索
		v1.GET("video/search",Controller.SearchVideo)

		//用户的视频
		v1.GET("video/user_video",Controller.UserVideo)

		//弹幕
		v1.POST("video/barrage",Controller.SentBarrage)
		v1.GET("video/barrage",Controller.ShowBarrage)

		//三连
		v1.PUT("video/like",middleware.Auth(),Controller.Like)
		v1.PUT("video/coin",middleware.Auth(),Controller.Coin)
		v1.PUT("video/collection",middleware.Auth(),Controller.Collection)
	}

	err:=r.Run(":8080")
	if err!=nil{
		panic(err)
	}
}