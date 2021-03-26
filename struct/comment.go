package _struct

type Comment struct {
	Id		string			`form:"id" json:"id" gorm:"primary_key"`
	Content	string			`form:"content" json:"content"`
	Uid		string			`form:"uid" json:"uid"`
	Vid		string			`form:"vid" json:"vid"`
}