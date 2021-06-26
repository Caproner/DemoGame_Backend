package test

import (
	_ "github.com/go-sql-driver/mysql"

	//"github.com/go-xorm/xorm"

	"github.com/Caproner/DemoGame_Backend/database/orm"
	"testing"
)
type User struct{
	Id int `xorm:" not null INT(11)"`
	UserName string `xorm:"not null VARCHAR(40)"`
}

func TestOrm(t *testing.T){
	orm1 := orm.DefaultEngine()

	user := new(User)
	user.Id = 3
	user.UserName ="dk3"
	aff, err := orm1.Insert(user)
	t.Log(aff, err)
}