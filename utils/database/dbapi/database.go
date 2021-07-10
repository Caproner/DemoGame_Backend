package dbapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Caproner/DemoGame_Backend/utils/database/ormredis"

	r "github.com/Caproner/DemoGame_Backend/app/global/record"
)

func init(){
	ormredis.InitRedis()
}

func FindPlayer(openId string) (*r.Player, error) {
	p, err := findPlayerInRedis(openId)
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("not databse ")
	}
	return p, nil
}

func SavePlayer(p *r.Player) error {
	return setPlayerInRedis(p)
}

func ItemLenAdd(key string) int64{
	return findLenAddInRedis(key)
}


func findPlayerInRedis(openId string) (*r.Player, error){
	rtype, err := ormredis.RDB().KVGet(openId)
	if err != nil{
		return nil, err
	}
	p := &r.Player{}
	_ = json.Unmarshal([]byte(rtype.(string)), p)
	return p, nil
}

func setPlayerInRedis(p *r.Player) error{
	err := ormredis.RDB().KVSet(p.OpenId,*p)
	fmt.Println(err)
	return err
}

func findLenAddInRedis(key string)int64{
	return ormredis.RDB().Incr(key)
}