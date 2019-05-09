package main

import (
	"fmt"
	"github.com/ladygogo/BattleGopher/player"
)

func main() {
	player1, _ := player.InitializePlayer("Linda")
	player2, _ := player.InitializePlayer("Player 2")
	playerArray := [...]player.Player{player1, player2}

	for i, _ := range playerArray {
		playerArray[i].NewBoard(2)
	}

	fmt.Println("Welcome to Battle Gopher!")
	fmt.Println("---------------------")

	var row, column int
	playerIndex := 0

	for playerArray[0].Gameboard.AllGophersSunk() == false && playerArray[1].Gameboard.AllGophersSunk() == false {
		fmt.Printf("%s's Turn\n", playerArray[(playerIndex)%2].Name)

		playerArray[(playerIndex + 1)%2].Gameboard.PrettyPrint()	

		fmt.Print("Enter a row-> ")
	    fmt.Scanf("%d", &row)
		fmt.Print("Enter a column-> ")
	    fmt.Scanf("%d", &column)
	
	    hit, err := playerArray[(playerIndex + 1)%2].Gameboard.CheckForHit(row - 1, column - 1)

	    if err != nil {
	    	fmt.Println(err)
	    	fmt.Println("Try again\n")
	    } else {
	    	playerArray[(playerIndex + 1)%2].Gameboard.PrettyPrint()	
	    }

	    if hit == false && err == nil {
	    	playerIndex ++
	    }
	}
}
