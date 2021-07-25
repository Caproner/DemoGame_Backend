package pmoney

import (
	rp "github.com/Caproner/DemoGame_Backend/include/global"
)

func GetTypeMoney(t int, p *rp.Player)int64{
	if v,ok := p.Money[t];ok{
		return v
	}else{
		return int64(0)
	}
}

func SetTypeMoney(k int,v int64, p *rp.Player)bool{
	p.Money[k] = v
	return true
}

func SyncMoney(kv map[int]int64,  p *rp.Player)bool{
	p.Money = kv
	return true
}

