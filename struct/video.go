package _struct

type Video struct {
	ID			int		`form:"vid" json:"vid"`
	UID			int		`form:"uid" json:"uid"`
	Title		string	`form:"title" json:"title"`
	Info		string	`form:"info" json:"info"`
	Part		string	`form:"part" json:"part"`
	VideoURL	string	`form:"video_url" json:"video_url"`
	CoverURL	string	`form:"cover_url" json:"cover_url"`
	UpTime		string	`form:"up_time" json:"up_time"`
}

type VideoInfo struct {
	ID			int		`form:"vid" json:"vid"`
	UID			int		`form:"uid" json:"uid"`
	Plays		int		`form:"plays" json:"plays"`
	Likes		int		`form:"likes" json:"likes"`
	Coins		int		`form:"coins" json:"coins"`
	Collections	int		`form:"collections" json:"collections"`
	Online		int		`form:"online" json:"online"`
}