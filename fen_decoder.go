package fen_decoder

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	asciNum_0          = 48
	asciNumBefore_a    = 96
	initialPosition    = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	shortWhiteCastling = "e1g1"
	longWhiteCastling  = "e1c1"
	shortBlackCastling = "e8g8"
	longBlackCastling  = "e8c8"
)

var castlingAlias = map[string]string{
	shortWhiteCastling: "O-O",
	longWhiteCastling:  "O-O-O",
	shortBlackCastling: "O-O",
	longBlackCastling:  "O-O-O",
}

type Board struct {
	initialPosition    string
	board              map[int]map[int]rune
	originalMoves      []string
	movesWithFigures   []string
	movesWithShortForm []string
	countMoves         int
}

func (b *Board) SetInitWithPosition(s string) {
	b.Reset()
	b.initialPosition = s
	b.Init()
}
func (b *Board) Reset() {
	b.board = nil
	b.originalMoves = nil
	b.movesWithFigures = nil
	b.movesWithShortForm = nil
	b.countMoves = 0
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
	v1, h1 := getIntCell(move[:2])

	if b.board[h1][v1] == 0 {
		return fmt.Errorf("The is no figure on the %s cell", move[0:2])
	}

	b.originalMoves = append(b.originalMoves, move)
	b.addMoveWithFigure(move)
	b.addShortMove(move)
	b.swapFigures(move)
	b.countMoves++
	return nil
}

func (b *Board) swapFigures(move string) {
	if b.isCastling(move) {
		switch move {
		case shortWhiteCastling:
			b.board[1][5] = 0
			b.board[1][6] = 'R'
			b.board[1][7] = 'K'
			b.board[1][8] = 0
		case longWhiteCastling:
			b.board[1][5] = 0
			b.board[1][4] = 'R'
			b.board[1][3] = 'K'
			b.board[1][1] = 0
		case shortBlackCastling:
			b.board[8][5] = 0
			b.board[8][6] = 'r'
			b.board[8][7] = 'k'
			b.board[8][8] = 0
		case longBlackCastling:
			b.board[8][5] = 0
			b.board[8][4] = 'r'
			b.board[8][3] = 'k'
			b.board[8][1] = 0
		}
	} else {
		v1, h1 := getIntCell(move[:2])
		v2, h2 := getIntCell(move[2:4])
		b.board[h2][v2] = b.board[h1][v1]
		b.board[h1][v1] = 0
	}
}

func (b *Board) GetMovesWithFigures() string {
	return strings.Join(b.movesWithFigures, ", ")
}

func (b *Board) GetMovesWithShortForm() string {
	return strings.Join(b.movesWithShortForm, ", ")
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

func (b *Board) addMoveWithFigure(move string) {
	v, h := getIntCell(move[:2])
	b.movesWithFigures = append(b.movesWithFigures, string(b.board[h][v])+" "+move)
}

func (b *Board) addShortMove(move string) {
	short_move := ""
	if b.isCastling(move) {
		short_move = castlingAlias[move]
	} else {
		v1, h1 := getIntCell(move[:2])
		v2, h2 := getIntCell(move[2:4])

		figure := b.board[h1][v1]
		short_move = string(figure) + " " + move

		switch {
		case figure == 'n' || figure == 'N':
			if !b.canTwoKnightMove(figure, h2, v2) {
				short_move = string(figure) + " " + move[2:4]
			}
		case figure == 'r' || figure == 'R':
			if !b.canTwoFigureLineMove(figure, h2, v2) {
				short_move = string(figure) + " " + move[2:4]
			}
		case figure == 'p' || figure == 'P':
			if v1 == v2 {
				short_move = move[2:4]
			}
		}
	}
	b.movesWithShortForm = append(b.movesWithShortForm, short_move)
}

func (b *Board) canTwoKnightMove(figure rune, hor, ver int) bool {
	ar := []int{2, 1, -2, 1, 2, -1, -2, -1, 2}
	count := 0
	for i := 0; i < len(ar)-1; i++ {
		potetial_hor, potetial_ver := hor+ar[i], ver+ar[i+1]
		if potetial_hor < 1 || potetial_hor > 8 || potetial_ver < 1 || potetial_ver > 8 {
			continue
		}
		if b.board[potetial_hor][potetial_ver] == figure {
			count++
		}
	}
	return count >= 2
}

func (b *Board) canTwoFigureLineMove(figure rune, hor, ver int) bool {
	count := 0
	// Check right direction
	for i := ver + 1; i <= 8; i++ {
		if b.board[hor][i] == figure {
			count++
		} else if b.board[hor][i] != 0 {
			break
		}
	}
	// Check left direction
	for i := ver - 1; i >= 1; i-- {
		if b.board[hor][i] == figure {
			count++
		} else if b.board[hor][i] != 0 {
			break
		}
	}

	// Check top direction
	for i := hor + 1; i <= 8; i++ {
		if b.board[i][ver] == figure {
			count++
		} else if b.board[i][ver] != 0 {
			break
		}
	}
	// Check bottom direction
	for i := hor - 1; i >= 1; i-- {
		if b.board[i][ver] == figure {
			count++
		} else if b.board[i][ver] != 0 {
			break
		}
	}
	return count >= 2
}

func (b *Board) isCastling(move string) bool {
	v, h := getIntCell(move[:2])
	if b.board[h][v] == 'K' || b.board[h][v] == 'k' {
		if _, ok := castlingAlias[move]; ok == true {
			return true
		}
	}
	return false
}

func getIntCell(move string) (int, int) {
	return int(move[0]) - asciNumBefore_a, int(move[1]) - asciNum_0
}

func getCell(ver, hor int) string {
	return string(rune(ver+asciNumBefore_a)) + string(rune(hor+asciNum_0))
}
