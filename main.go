package main

import (
	"fmt"

	"github.com/schwerdt/BattleGopher/gameboard"
)

func main() {
	g := gameboard.GameBoard{BoardDimension: 5}
	err := g.InitializeBoard()
	fmt.Println(err)
	fmt.Println(g.MaskedBoard)
	fmt.Println(g.UnmaskedBoard)
	fmt.Println(g.CheckForHit(2, 2))
	fmt.Println(g.CheckForHit(0, 1))
	fmt.Println(g.MaskedBoard)
	fmt.Println(g.UnmaskedBoard)
	fmt.Println("Are all gophers sunk?", g.AllGophersSunk())
	g.PrettyPrint()

}
