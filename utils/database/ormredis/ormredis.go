package ormredis

import (
	"encoding/json"
	"errors"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"github.com/go-redis/redis"
	"time"
)

//var defaultctx = context.Background()
var defaultRedis *dBRedis

type dBRedis struct {
	redisClient *redis.Client
}

// InitRedis 初始化redis数据库
func InitRedis(){
	defaultRedis = &dBRedis{}
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

// RDB 外部调用，通常来说事dbapi调用
func RDB() *dBRedis{
	return defaultRedis
}

func (drs *dBRedis)KVSet(key string, value interface{}) error{
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

func (drs *dBRedis)KVGet(key string) (interface{}, error){
	r, err := drs.redisClient.Get( key).Result()
	if err == redis.Nil {
		return nil, errors.New("redis empty")
	}
	return r, err
}

func (drs *dBRedis)Incr(key string)int64{
	num, _ := drs.redisClient.Incr( key).Result()
	return num
}

func (drs *dBRedis)IsTimelessKeyActive(k string) bool{
	if _, err := drs.redisClient.Get(k).Result();err != redis.Nil{
		return true
	}else{
		return false
	}
}

func (drs *dBRedis)TimeLessKeyUpate(k string) bool{
	if _, err := drs.redisClient.Get(k).Result();err != redis.Nil{
		_,_ = drs.redisClient.Del(k).Result()
		if _, e := drs.redisClient.Set(k,0, 180 * time.Second).Result();e != nil{
			log.Info(e)
		}
		return true
	}else{
		if _, e := drs.redisClient.Set(k,0, 180 * time.Second).Result();e != nil{
			log.Info(e)
		}
		return false
	}
}