package player

import (
	"errors"
////////"fmt"
	"github.com/ladygogo/BattleGopher/gameboard"
////////"github.com/davecgh/go-spew/spew"
)

type Player struct {
	Name string
	Score int
	Gameboard gameboard.GameBoard
}

func InitializePlayer(name string) (Player, error) {
	player := Player{Name: name, Score: 0}
	if player.Name == "" {
		return Player{}, errors.New("Player name cannot be blank!")
	}
	return player, nil
}

func (p *Player) NewBoard(boardDimension int) error {
	g, err := gameboard.InitializeBoard(boardDimension)
	if err != nil {
		return err
	}
	p.Gameboard = g
	return nil
}
