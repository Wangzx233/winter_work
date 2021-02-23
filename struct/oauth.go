package _struct

type OauthConf struct {
	ClientID		string
	ClientSecret	string
	RedirectUrl		string
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}
