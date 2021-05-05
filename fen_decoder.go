package fen_decoder

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	asciNum_0       = 48
	asciNumBefore_a = 96
	initialPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

type Board struct {
	initialPosition  string
	board            map[int]map[int]rune
	originalMoves    []string
	movesWithFigures []string
	countMoves       int
}

func (b *Board) SetInitialPosition(s string) {
	b.initialPosition = s
	b.Init()
}

func (b *Board) CountMoves() int {
	return b.countMoves
}

func (b *Board) MakeMoves(moves string) error {
	amoves := strings.Split(moves, " ")
	for _, m := range amoves {
		if err := b.MakeMove(m); err != nil {
			return err
		}
	}
	return nil
}

func (b *Board) MakeMove(move string) error {
	valid_move := regexp.MustCompile(`[a-h][1-8][a-h][1-8]`)
	if !valid_move.MatchString(move) {
		return errors.New("Wrong move param")
	}

	v1 := int(move[0]) - asciNumBefore_a
	h1 := int(move[1]) - asciNum_0

	v2 := int(move[2]) - asciNumBefore_a
	h2 := int(move[3]) - asciNum_0

	if b.board[h1][v1] == 0 {
		return fmt.Errorf("The is no figure on the %s cell", move[0:2])
	}

	b.board[h2][v2] = b.board[h1][v1]
	b.board[h1][v1] = 0
	b.originalMoves = append(b.originalMoves, move)
	b.movesWithFigures = append(b.movesWithFigures, string(b.board[h2][v2])+" "+move)
	b.countMoves++
	return nil
}

func (b *Board) GetMovesWithFigures() string {
	return strings.Join(b.movesWithFigures, ", ")
}

func (b *Board) Init() {
	if b.initialPosition == "" {
		b.initialPosition = initialPosition
	}
	if b.board == nil {
		b.board = make(map[int]map[int]rune)
		for i := 1; i <= 8; i++ {
			b.board[i] = map[int]rune{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
		}
	}
	b.buildBoardMap()
}

func (b *Board) buildBoardMap() {
	re := regexp.MustCompile(`([rnbqkbnrRNBQKBNR1-8pP]{1,8})`)
	horizontals := re.FindAllString(b.initialPosition, 8)

	for i, j := 0, len(horizontals)-1; i < j; i, j = i+1, j-1 {
		horizontals[i], horizontals[j] = horizontals[j], horizontals[i]
	}

	for i, horizontal := range horizontals {
		vertical := 1
		for _, char := range horizontal {
			if char >= '1' && char <= '8' {
				vertical += int(char) - asciNum_0
			} else {
				b.board[i+1][vertical] = char
				vertical++
			}
		}
	}
}
