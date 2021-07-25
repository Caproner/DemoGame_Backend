package pcul

import (
	rp "github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/include/localvar/cultivatevar"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
)

// GetBuildS 获得养成——建筑物的数据
func GetBuildS(p *rp.Player)[]int{
	if r, ok := p.Cultivate[cultivatevar.BuildS];ok{
		return tr.InterfaceToIntList(r)
	}else{
		return []int{}
	}
}
// SetBuildS 设置养成——建筑物数据
func SetBuildS(d []int, p *rp.Player)bool{
	p.Cultivate[cultivatevar.BuildS] = d
	return true
}
// SyncBuildS 设置养成——建筑物数据
func SyncBuildS(d []int, p *rp.Player)bool{
	//log.Info(d)
	p.Cultivate[cultivatevar.BuildS] = d
	return true
}
