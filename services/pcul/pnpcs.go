package pcul

import (
	rp "github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/include/localvar/cultivatevar"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
)

// GetNpcS 获得养成——NPC的数据
func GetNpcS(p *rp.Player)[]int{
	if r, ok := p.Cultivate[cultivatevar.NpcS];ok{
		return tr.InterfaceToIntList(r)
	}else{
		return []int{}
	}
}
// SetNpcS 设置养成——NPC的数据
func SetNpcS(d []int, p *rp.Player)bool{
	p.Cultivate[cultivatevar.NpcS] = d
	return true
}

// SyncNpcS 设置养成——NPC的数据
func SyncNpcS(d []int, p *rp.Player)bool{
	//log.Info(d)
	p.Cultivate[cultivatevar.NpcS] = d
	return true
}