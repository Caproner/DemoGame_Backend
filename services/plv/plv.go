package plv

import (
	rp "github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/utils/log"
)

func SyncLv(lv int, exp int64,  p *rp.Player)bool{
	if p.Lv > lv{
		log.Infof("sync lv:%d is error to %d", lv, p.Lv)
		return false
	}else if p.Lv == lv && p.Exp > exp{
		log.Infof("sync exp:%d is error to %d", exp, p.Exp)
		return false
	}
	//log.Info(lv, exp)
	p.Lv = lv
	p.Exp = exp
	return true
}
