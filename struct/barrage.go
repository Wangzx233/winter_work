package _struct

type Barrage struct {
	Id			int			`form:"id" json:"id"`
	Uid			string		`form:"uid" json:"uid"`
	Vid			string		`form:"vid" json:"vid"`
	Content		string		`form:"content" json:"content"`
	time		string		`form:"time" json:"time"`
}
