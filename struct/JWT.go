package _struct

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type Payload struct {
	Iss string `json:"iss"`
	Exp string `json:"exp"`
	Iat string `json:"iat"`
	Id  uint    `json:"id"`
}
type JWT struct {
	Header    Header
	Payload   Payload
	Signature string
	Token     string
}


