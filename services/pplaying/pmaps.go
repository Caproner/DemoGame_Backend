package pplaying

import (
	rp "github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/include/localvar/cultivatevar"
	"github.com/Caproner/DemoGame_Backend/include/localvar/playingvar"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
)

// GetMapS 获得养成——NPC的数据
func GetMapS(p *rp.Player)map[string]interface{}{
	if r, ok := p.Playing[playingvar.MapS];ok{
		return tr.InterfaceToMStringFace(r)
	}else{
		return make(map[string]interface{})
	}
}
// SetNpcS 设置养成——NPC的数据
func SetMapS(d []int, p *rp.Player)bool{
	p.Cultivate[cultivatevar.NpcS] = d
	return true
}

// SyncNpcS 设置养成——NPC的数据
func SyncMapS(d map[string]interface{}, p *rp.Player)bool{
	//log.Info(d)
	p.Playing[playingvar.MapS] = d
	return true
}
