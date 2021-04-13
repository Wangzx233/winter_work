package Controller

import (
	_struct "1/struct"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"time"
)

func NewJWT(id uint) _struct.JWT {
	var jwt _struct.JWT
	jwt.Header = _struct.Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	jwt.Payload = _struct.Payload{
		Iss: "bibi",
		Exp: strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat: strconv.FormatInt(time.Now().Unix(), 10),
		Id:  id,
	}
	h, _ := json.Marshal(jwt.Header)
	p, _ := json.Marshal(jwt.Payload)
	baseh := base64.StdEncoding.EncodeToString(h)
	basep := base64.StdEncoding.EncodeToString(p)
	secret := baseh + "." + basep
	key := "pass"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(secret))
	s := mac.Sum(nil)
	jwt.Signature = base64.StdEncoding.EncodeToString(s)
	jwt.Token = secret + "." + jwt.Signature
	return jwt
}