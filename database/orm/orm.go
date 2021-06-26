package orm

import (
	"github.com/Caproner/DemoGame_Backend/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var ormEngine *xorm.Engine

func DefaultEngine() *xorm.Engine{
	return ormEngine
}

func init(){
	var err error
	var sqlType = "mysql"
	var sqlSourceName = "root:root@tcp(127.0.0.1:3306)/test"
	if ormEngine, err = xorm.NewEngine(sqlType, sqlSourceName);err != nil{
		log.Error(err)
	}
	if err := ormEngine.Ping();err!= nil{
		log.Error(err)
	}
	log.Info("test database engine had ready")
	ormEngine.ShowSQL(true)
	ormEngine.SetMaxIdleConns(2)
	ormEngine.SetMaxOpenConns(1)
	preFixMapper := core.NewPrefixMapper(core.SameMapper{},"")
	ormEngine.SetMapper(preFixMapper)
}

// 创建table表 未测试
func CreateTable(table ...interface{}) (err error){
	err = ormEngine.CreateTables(table)
	return
}
// 向数据表插入数据，接口数据未校验，需要外部校验完传入
func Insert(data interface{})(affect int64, err error){
	affect, err = ormEngine.Insert(data)
	return
}


