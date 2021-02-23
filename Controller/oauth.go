package Controller

import (
	"1/Model"
	_struct "1/struct"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
var conf = _struct.OauthConf{
	ClientID:     "c022882a207c3d946f39",
	ClientSecret: "c7b0ace6c6340b975d33da9040cc21b1225fc4f7",
	RedirectUrl:  "http://115.29.201.183/api/v1/user/oauth",
}


func Oauth(c *gin.Context)  {
	code := c.Request.URL.Query().Get("code")

	//token
	var tokenAuthUrl = GetTokenAuthUrl(code)
	var token *Token
	 token, err := GetToken(tokenAuthUrl); if err != nil {
		c.JSON(400,err)
		return
	}
	userInfo,err := GetUserInfo(token); if err != nil {
		c.JSON(400,err)
		return
	}
	if Model.Oauth(userInfo,token.AccessToken){
		c.SetCookie("uid",token.AccessToken, 9999, "/", "/", false, false)
	}
}

func GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		conf.ClientID, conf.ClientSecret, code,
	)
}

// 获取 token
func GetToken(url string) (*Token, error) {

	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}

	// 将响应体解析为 token，并返回
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

// 获取用户信息
func GetUserInfo(token *Token) (map[string]interface{}, error) {

	// 形成请求
	var userInfoUrl = "https://api.github.com/user"	// github用户信息获取接口
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}