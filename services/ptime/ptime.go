package ptime

import (
	rp "github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/include/localvar/timevar"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
)

func GetTime(t string, p *rp.Player)int64{
	if v,ok := p.TimeClock[t];ok{

		return int64(v.(float64))
	}else{
		return int64(0)
	}
}

func SetTime(i string, t int64, p *rp.Player)bool{
	p.TimeClock[i] = t
	return true
}

func GetTimes(p *rp.Player)map[string]int64{
	m := make(map[string]int64)
	for k,v := range p.TimeClock{
		switch k {
		case timevar.LoginTimeType:
			k = "1"
		case timevar.LastActionTimeType:
			k = "2"
		case timevar.OfflineTimeType:
			k = "3"
		default:
			k = "0"
		}
		m[k] = tr.InterfaceToInt64(v)
	}
	return m
}
