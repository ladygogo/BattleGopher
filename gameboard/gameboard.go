package gameboard

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var randomNumber = rand.New(rand.NewSource(time.Now().UnixNano()))

type GameBoard struct {
	BoardDimension int
	UnmaskedBoard  [][]int
	MaskedBoard    [][]string
}

func (g *GameBoard) InitializeBoard() error {
	if g.BoardDimension < 1 {
		return errors.New("Board dimension is less than 1")
	}

	g.UnmaskedBoard = make([][]int, g.BoardDimension)
	g.MaskedBoard = make([][]string, g.BoardDimension)

	for col := 0; col < g.BoardDimension; col++ {
		g.UnmaskedBoard[col] = make([]int, g.BoardDimension)
		g.MaskedBoard[col] = make([]string, g.BoardDimension)
	}

	for row := 0; row < g.BoardDimension; row++ {
		for col := 0; col < g.BoardDimension; col++ {
			g.UnmaskedBoard[row][col] = randomNumber.Intn(2)
			g.MaskedBoard[row][col] = "?"
		}
	}
	return nil

}

func (g *GameBoard) CheckForHit(row int, col int) (bool, error) {
	if row < 0 || row >= g.BoardDimension {
		return false, errors.New("row dimension is not valid")
	}
	if col < 0 || col >= g.BoardDimension {
		return false, errors.New("col dimension is not valid")
	}
	if g.UnmaskedBoard[row][col] == 1 {
		g.MaskedBoard[row][col] = "H"
		return true, nil
	} else {
		g.MaskedBoard[row][col] = "M"
		return false, nil
	}
}

func (g GameBoard) AllGophersSunk() bool {
	for row := 0; row < g.BoardDimension; row++ {
		for col := 0; col < g.BoardDimension; col++ {
			if g.UnmaskedBoard[row][col] == 1 && g.MaskedBoard[row][col] == "?" {
				return false
			}
		}
	}
	return true
}

func (g GameBoard) PrettyPrint() {
	for row := 0; row < g.BoardDimension; row++ {
		for col := 0; col < g.BoardDimension; col++ {
			fmt.Printf("  %s  ", g.MaskedBoard[row][col])
		}
		fmt.Println()
	}
}
