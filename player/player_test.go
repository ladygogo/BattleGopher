package player

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializePlayer(t *testing.T) {
	player, err := InitializePlayer("Gophie")
	assert.Equal(t, nil, err)
	assert.Equal(t, "Gophie", player.Name)
}

func TestInitializePlayerBlankName(t *testing.T) {
	player, err := InitializePlayer("")
	assert.Equal(t, errors.New("Player name cannot be blank!"), err)
	assert.Equal(t, Player{}, player)
}

func TestNewBoard(t *testing.T) {
	player, err := InitializePlayer("Gophette")
	assert.NoError(t, err)

	err = player.NewBoard(2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, player.Gameboard.BoardDimension)
}

func TestNewBoardInvalidDimension(t *testing.T) {
	player, err := InitializePlayer("Gophette")
	assert.NoError(t, err)

	err = player.NewBoard(0)
	assert.Error(t, err)
	assert.Error(t, errors.New("Board dimension is less than 1"), err)
}
