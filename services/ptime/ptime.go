package ptime

import (
	rp "github.com/Caproner/DemoGame_Backend/include/global"
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
