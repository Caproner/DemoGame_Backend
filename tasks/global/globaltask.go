package global

import "github.com/Caproner/DemoGame_Backend/services/playersvr"

func init() {

}

func StartGlobalTask() {
	go playersvr.DefaultSvr().Loop()
}
