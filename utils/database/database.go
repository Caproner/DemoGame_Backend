package database

import (
	"errors"

	r "github.com/Caproner/DemoGame_Backend/app/global/record"
)

func FindPlayer(openId string) (*r.Player, error) {
	return &r.Player{}, errors.New("not databse ")
}

func SavePlayer(p *r.Player) error {
	return nil
}
