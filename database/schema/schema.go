package schema

type User struct{
	Id int `xorm:" not null INT(11)"`
	Username string `xorm:"not null VARCHAR(40)"`
}