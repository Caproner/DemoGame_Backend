package ormredis

import (
	"encoding/json"
	"errors"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"github.com/go-redis/redis"
)

//var defaultctx = context.Background()
var defaultRedis *DBRedis

type DBRedis struct {
	redisClient *redis.Client
}

func InitRedis(){
	defaultRedis = &DBRedis{}
	defaultRedis.redisClient = redis.NewClient(&redis.Options{
		Addr : "127.0.0.1:6379",
		Password : "",
		DB: 0,
	})

	pong, err := defaultRedis.redisClient.Ping().Result()
	if err != nil{
		log.Error(err)
	}
	log.Info(pong)
}

// 外部调用，通常来说事dbapi调用
func RDB() *DBRedis{
	return defaultRedis
}

func (drs *DBRedis)KVSet(key string, value interface{}) error{
	marvalue, err := json.Marshal(value)
	if err != nil{
		return err
	}
	exists, err := drs.redisClient.Exists( key).Result()
	if err!= nil{
		panic(err)
	}
	if exists > 0{
		_,_ = drs.redisClient.Del( key).Result()
	}
	_, err = drs.redisClient.Set( key, marvalue,0).Result()
	return err

}

func (drs *DBRedis)KVGet(key string) (interface{}, error){
	r, err := drs.redisClient.Get( key).Result()
	if err == redis.Nil {
		return nil, errors.New("redis empty")
	}
	return r, err
}

func (drs *DBRedis)Incr(key string)int64{
	num, _ := drs.redisClient.Incr( key).Result()
	return num
}