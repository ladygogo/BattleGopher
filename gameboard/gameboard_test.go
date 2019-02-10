package gameboard

import(
	"testing"
	"math/rand"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestInitializeBoard(t *testing.T) {
	randomNumber = rand.New(rand.NewSource(1))

	g := GameBoard{BoardDimension: 4 }
	err := g.InitializeBoard()

	require.NoError(t, err)
	assert.Equal(t, 4, len(g.UnmaskedBoard))
	assert.Equal(t, 4, len(g.UnmaskedBoard[0]))

	expectedBoard := [][]int{[]int{1, 1, 1, 1}, []int{1, 0, 1, 0}, []int{0, 0, 0, 1}, []int{0, 1, 0, 0}}
	expectedMaskedBoard := [][]string{[]string{ "?", "?", "?", "?"}, []string{"?", "?", "?", "?"}, []string{"?", "?", "?", "?"}, []string{"?", "?", "?", "?"}}
	assert.Equal(t, expectedBoard, g.UnmaskedBoard)
	assert.Equal(t, expectedMaskedBoard, g.MaskedBoard)
}

func TestInitializeGameBoardInvalidBoardDimension(t *testing.T) {
	g := GameBoard{BoardDimension: -1 }

	err := g.InitializeBoard()

	assert.Error(t, err)
	assert.Equal(t, errors.New("Board dimension is less than 1"), err)
}

func TestCheckForHitGopherFound(t *testing.T) {
	randomNumber = rand.New(rand.NewSource(1))

	g := GameBoard{BoardDimension: 4 }
	g.InitializeBoard()

	hit, err := g.CheckForHit(0, 0)

	require.NoError(t, err)
	assert.True(t, true, hit)
}

func TestCheckForHitGopherNotFound(t *testing.T) {
	randomNumber = rand.New(rand.NewSource(1))

	g := GameBoard{BoardDimension: 4 }
	g.InitializeBoard()

	hit, err := g.CheckForHit(1, 1)

	assert.False(t, hit)
	assert.Nil(t, err)
}

func TestCheckForHitInvalidRowDimension(t *testing.T) {
	g := GameBoard{BoardDimension: 4 }
	g.InitializeBoard()

	hit, err := g.CheckForHit(-1, 1)

	assert.False(t, hit)
	assert.EqualError(t, err, "row dimension is not valid")
}

func TestCheckForHitInvalidColumnDimension(t *testing.T) {
	g := GameBoard{BoardDimension: 4 }
	g.InitializeBoard()

	hit, err := g.CheckForHit(1, -2)

	assert.False(t, hit)
	assert.EqualError(t, err, "col dimension is not valid")
}

func TestAllGophersSunkTrue(t *testing.T) {
	randomNumber = rand.New(rand.NewSource(3))

	g := GameBoard{BoardDimension: 2}
	g.InitializeBoard()

	hit, _ := g.CheckForHit(0, 1)

	assert.True(t, hit)

	allSunk := g.AllGophersSunk()
	assert.True(t, allSunk)
}

func TestAllGophersSunkFalse(t *testing.T) {
	randomNumber = rand.New(rand.NewSource(3))

	g := GameBoard{BoardDimension: 2}
	g.InitializeBoard()

	allSunk := g.AllGophersSunk()
	assert.False(t, allSunk)
}
