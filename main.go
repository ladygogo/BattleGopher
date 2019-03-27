package main

import (
	"fmt"

	"github.com/ladygogo/BattleGopher/gameboard"
	"github.com/ladygogo/BattleGopher/player"
)

func main() {
	g, err := gameboard.InitializeBoard(5)
	fmt.Println(err)
	fmt.Println(g.MaskedBoard)
	fmt.Println(g.UnmaskedBoard)
	fmt.Println(g.CheckForHit(2, 2))
	fmt.Println(g.CheckForHit(0, 1))
	fmt.Println(g.MaskedBoard)
	fmt.Println(g.UnmaskedBoard)
	fmt.Println("Are all gophers sunk?", g.AllGophersSunk())
	g.PrettyPrint()

	player, err := player.InitializePlayer("Linda")
	fmt.Println(player.Name, player.Score)
	player.NewBoard(5)
	player.Gameboard.PrettyPrint()
}
