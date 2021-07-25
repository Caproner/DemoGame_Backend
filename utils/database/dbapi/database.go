package dbapi

import (
	"encoding/json"
	"errors"
	rp "github.com/Caproner/DemoGame_Backend/include/global"

	"github.com/Caproner/DemoGame_Backend/utils/database/ormredis"

)

func init() {
	ormredis.InitRedis()
}

// FindPlayer 查找是否已经有该玩家数据
func FindPlayer(openID string) (*rp.Player, error) {
	p, err := findPlayerInRedis(openID)
	// fmt.Println(err)
	if err != nil {
		return nil, errors.New("not databse ")
	}
	return p, nil
}

// SavePlayer 保存玩家数据
func SavePlayer(p *rp.Player) error {
	return setPlayerInRedis(p)
}

// ItemLenAdd 自增项返回值
func ItemLenAdd(key string) int64 {
	return findLenAddInRedis(key)
}

func CheckUserToken(openId string) bool{
	key := UserToken(openId)
	return ormredis.RDB().IsTimelessKeyActive(key)
}

func UpdateToken(openId string) bool{
	key := UserToken(openId)
	return ormredis.RDB().TimeLessKeyUpate(key)
}

func UserToken(openId string) string{
	key := "ToKen_" + openId
	return key
}

func findPlayerInRedis(openID string) (*rp.Player, error) {
	rtype, err := ormredis.RDB().KVGet(openID)
	if err != nil {
		return nil, err
	}
	p := &rp.Player{}
	_ = json.Unmarshal([]byte(rtype.(string)), p)
	return p, nil
}

func setPlayerInRedis(p *rp.Player) error {
	err := ormredis.RDB().KVSet(p.OpenID, *p)
	// fmt.Println(err)
	return err
}

func findLenAddInRedis(key string) int64 {
	return ormredis.RDB().Incr(key)
}
