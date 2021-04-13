package middleware

import (
	_struct "1/struct"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		jwt, err := checkToken(token)
		fmt.Println(jwt)
		if err == nil {
			c.Next()
		} else {
			c.Abort()
			c.JSON(500,[]byte("请先登录"))
		}
	}
}

func checkToken(token string) (jwt _struct.JWT,err error) {
	err = errors.New("token error")
	arr := strings.Split(token, ".")
	if len(arr) < 3 {
		fmt.Println("59------", err)
		return
	}
	baseh := arr[0]
	h, err := base64.StdEncoding.DecodeString(baseh)
	if err != nil {
		fmt.Println("decode header", err)
		return
	}
	err = json.Unmarshal(h, &jwt.Header)
	if err != nil {
		fmt.Println("unmarshal header", err)
		return
	}
	basep := arr[1]
	p, err := base64.StdEncoding.DecodeString(basep)
	if err != nil {
		fmt.Println("decode payload", err)
		return
	}
	err = json.Unmarshal(p, &jwt.Payload)
	if err != nil {
		fmt.Println("unmarshal payload", err)
		return
	}
	bases := arr[2]
	s1, err := base64.StdEncoding.DecodeString(bases)
	if err != nil {
		fmt.Println("decode bases", err)
		return
	}
	se := baseh + "." + basep
	w := []byte("pass")
	mac := hmac.New(sha256.New, w)
	mac.Write([]byte(se))
	s2 := mac.Sum(nil)
	if string(s1) != string(s2) {
		return
	} else {
		jwt.Signature = arr[2]
		jwt.Token = token
	}
	return jwt, nil
}