package fen_decoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitWithDefaultPosition(t *testing.T) {
	board := &Board{}
	board.Init()
	expected := map[int]map[int]rune{
		1: {1: 'R', 2: 'N', 3: 'B', 4: 'Q', 5: 'K', 6: 'B', 7: 'N', 8: 'R'},
		2: {1: 'P', 2: 'P', 3: 'P', 4: 'P', 5: 'P', 6: 'P', 7: 'P', 8: 'P'},
		3: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		4: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		5: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		6: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		7: {1: 'p', 2: 'p', 3: 'p', 4: 'p', 5: 'p', 6: 'p', 7: 'p', 8: 'p'},
		8: {1: 'r', 2: 'n', 3: 'b', 4: 'q', 5: 'k', 6: 'b', 7: 'n', 8: 'r'},
	}

	assert.EqualValues(t, expected, board.board)

}

func TestSetInitialPosition(t *testing.T) {
	board := &Board{}
	// Position after e2-e4, d7-d5
	board.SetInitWithPosition("rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 2")
	expected := map[int]map[int]rune{
		1: {1: 'R', 2: 'N', 3: 'B', 4: 'Q', 5: 'K', 6: 'B', 7: 'N', 8: 'R'},
		2: {1: 'P', 2: 'P', 3: 'P', 4: 'P', 5: 0, 6: 'P', 7: 'P', 8: 'P'},
		3: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		4: {1: 0, 2: 0, 3: 0, 4: 0, 5: 'P', 6: 0, 7: 0, 8: 0},
		5: {1: 0, 2: 0, 3: 0, 4: 'p', 5: 0, 6: 0, 7: 0, 8: 0},
		6: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		7: {1: 'p', 2: 'p', 3: 'p', 4: 0, 5: 'p', 6: 'p', 7: 'p', 8: 'p'},
		8: {1: 'r', 2: 'n', 3: 'b', 4: 'q', 5: 'k', 6: 'b', 7: 'n', 8: 'r'},
	}

	assert.EqualValues(t, expected, board.board)
}

func TestMakeMove(t *testing.T) {
	board := &Board{}

	board.Init()

	if err := board.MakeMove("q2e4"); err == nil {
		t.Fatal("Must be wrong move param")
	}

	if err := board.MakeMove("d3d4"); err == nil {
		t.Fatal("The is no figure on the d3 cell")
	}

	err := board.MakeMove("e2e4")
	if err != nil {
		t.Fatal("The move e2e4 must be without error")
	}
	board.MakeMove("d7d5")

	// Position after e2-e4, d7-d5 moves
	expected := map[int]map[int]rune{
		1: {1: 'R', 2: 'N', 3: 'B', 4: 'Q', 5: 'K', 6: 'B', 7: 'N', 8: 'R'},
		2: {1: 'P', 2: 'P', 3: 'P', 4: 'P', 5: 0, 6: 'P', 7: 'P', 8: 'P'},
		3: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		4: {1: 0, 2: 0, 3: 0, 4: 0, 5: 'P', 6: 0, 7: 0, 8: 0},
		5: {1: 0, 2: 0, 3: 0, 4: 'p', 5: 0, 6: 0, 7: 0, 8: 0},
		6: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		7: {1: 'p', 2: 'p', 3: 'p', 4: 0, 5: 'p', 6: 'p', 7: 'p', 8: 'p'},
		8: {1: 'r', 2: 'n', 3: 'b', 4: 'q', 5: 'k', 6: 'b', 7: 'n', 8: 'r'},
	}

	assert.EqualValues(t, []string{"e2e4", "d7d5"}, board.originalMoves)
	assert.EqualValues(t, 2, board.CountMoves())
	assert.EqualValues(t, expected, board.board)
}

func TestMakeMoves(t *testing.T) {
	board := &Board{}

	board.Init()

	err := board.MakeMoves("e2e4 d7d5 g1f3")
	if err != nil {
		t.Fatalf("The move %s must be without error", err)
	}
	// Position after e2-e4, d7-d5 moves
	expected := map[int]map[int]rune{
		1: {1: 'R', 2: 'N', 3: 'B', 4: 'Q', 5: 'K', 6: 'B', 7: 0, 8: 'R'},
		2: {1: 'P', 2: 'P', 3: 'P', 4: 'P', 5: 0, 6: 'P', 7: 'P', 8: 'P'},
		3: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 'N', 7: 0, 8: 0},
		4: {1: 0, 2: 0, 3: 0, 4: 0, 5: 'P', 6: 0, 7: 0, 8: 0},
		5: {1: 0, 2: 0, 3: 0, 4: 'p', 5: 0, 6: 0, 7: 0, 8: 0},
		6: {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
		7: {1: 'p', 2: 'p', 3: 'p', 4: 0, 5: 'p', 6: 'p', 7: 'p', 8: 'p'},
		8: {1: 'r', 2: 'n', 3: 'b', 4: 'q', 5: 'k', 6: 'b', 7: 'n', 8: 'r'},
	}

	assert.EqualValues(t, []string{"e2e4", "d7d5", "g1f3"}, board.originalMoves)
	assert.EqualValues(t, []string{"P e2e4", "p d7d5", "N g1f3"}, board.movesWithFigures)
	assert.EqualValues(t, []string{"e4", "d5", "N f3"}, board.movesWithShortForm)
	assert.EqualValues(t, 3, board.CountMoves())
	assert.EqualValues(t, expected, board.board)
}

func TestCanTwoKnightMove(t *testing.T) {
	board := &Board{}

	board.Init()
	assert.False(t, board.canTwoKnightMove('N', 6, 5))
	assert.False(t, board.canTwoKnightMove('N', 4, 5))

	// Two white Knights are on c5 and g5
	// They can both move on e4 and e6
	board.SetInitWithPosition("8/8/8/2N3N1/8/8/8/8 w - - 0 1")
	// Check e6
	assert.True(t, board.canTwoKnightMove('N', 6, 5))
	// Check e4
	assert.True(t, board.canTwoKnightMove('N', 4, 5))

	// Two white Knights are on g5 and h8
	// They can both move on f7
	board.SetInitWithPosition("7N/8/8/6N1/8/8/8/8 w - - 0 1")
	// Check f7
	assert.True(t, board.canTwoKnightMove('N', 7, 6))
	assert.False(t, board.canTwoKnightMove('N', 7, 7))
}

func TestСanTwoFigureLineMove(t *testing.T) {
	board := &Board{}
	board.Init()

	// Two white Rooks are on a5 and h8
	// They can both move on a8 and h5
	board.SetInitWithPosition("7R/8/8/R7/8/8/8/8 w - - 0 1")
	assert.True(t, board.canTwoFigureLineMove('R', 8, 1), "Check a8")
	assert.True(t, board.canTwoFigureLineMove('R', 5, 8), "Check h8")

	// We added black knight on e8
	// And now, there is one white rook(a5) that can move on a8
	board.SetInitWithPosition("4n2R/8/8/R7/8/8/8/8 w - - 0 1")
	assert.False(t, board.canTwoFigureLineMove('R', 8, 1), "Check a8")
	// h8 also is available
	assert.True(t, board.canTwoFigureLineMove('R', 5, 8), "Check h8")
}

func TestGetMovesWithFigures(t *testing.T) {
	board := &Board{}

	board.Init()
	board.MakeMoves("e2e4 d7d5 g1f3")
	assert.EqualValues(t, "P e2e4, p d7d5, N g1f3", board.GetMovesWithFigures())
}

func TestGetMovesWithShortForm(t *testing.T) {
	board := &Board{}

	board.Init()
	board.MakeMoves("e2e4 d7d5 g1f3")
	assert.EqualValues(t, "e4, d5, N f3", board.GetMovesWithShortForm())
}

func TestGetIntCell(t *testing.T) {
	v, h := getIntCell("d7")
	assert.EqualValues(t, 7, h)
	assert.EqualValues(t, 4, v)

}

func TestGetCell(t *testing.T) {
	assert.EqualValues(t, "e2", getCell(5, 2))
}
