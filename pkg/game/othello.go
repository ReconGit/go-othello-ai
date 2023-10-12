package game

type Cell int

const (
	EMPTY Cell = iota
	BLACK
	WHITE
	VALID
)

type State int

const (
	BLACK_TURN State = iota
	WHITE_TURN
	BLACK_WON
	WHITE_WON
	DRAW
)

type Othello struct {
	Board      [8][8]Cell
	State      State
	BlackScore int
	WhiteScore int
}

func NewOthello() Othello {
	board := [8][8]Cell{}
	board[3][3] = WHITE
	board[3][4] = BLACK
	board[4][3] = BLACK
	board[4][4] = WHITE
	board[2][3] = VALID
	board[3][2] = VALID
	board[4][5] = VALID
	board[5][4] = VALID

	return Othello{
		Board:      board,
		State:      BLACK_TURN,
		BlackScore: 2,
		WhiteScore: 2,
	}
}

func (o *Othello) DeepCopy() Othello {
	board := [8][8]Cell{}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			board[y][x] = o.Board[y][x]
		}
	}
	return Othello{
		Board:      board,
		State:      o.State,
		BlackScore: o.BlackScore,
		WhiteScore: o.WhiteScore,
	}
}

// MakeMove makes a move on the board and updates the game state.
func (o *Othello) MakeMove(position [2]int) {
	if o.State != BLACK_TURN && o.State != WHITE_TURN {
		panic("Can't make move: Game is over")
	}
	if o.Board[position[1]][position[0]] != VALID {
		panic("Can't make move: Position invalid" + string(rune(o.Board[position[1]][position[0]])))
	}

	var reverse Cell
	if o.State == BLACK_TURN {
		reverse = BLACK
	} else {
		reverse = WHITE
	}
	o.Board[position[1]][position[0]] = reverse
	for _, cell := range o.flipped_cells(position) {
		o.Board[cell[1]][cell[0]] = reverse
	}
	o.update_state()
}

// GetValidMoves returns a list of positions that can be played.
func (o *Othello) GetValidMoves() [][2]int {
	if o.State != BLACK_TURN && o.State != WHITE_TURN {
		return [][2]int{}
	}

	valid_moves := [][2]int{}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if o.Board[y][x] == VALID {
				valid_moves = append(valid_moves, [2]int{x, y})
			}
		}
	}
	return valid_moves
}

func (o *Othello) update_state() {
	o.BlackScore = 0
	o.WhiteScore = 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if o.Board[y][x] == BLACK {
				o.BlackScore++
			} else if o.Board[y][x] == WHITE {
				o.WhiteScore++
			}
		}
	}
	if o.is_board_full() || o.BlackScore == 0 || o.WhiteScore == 0 {
		o.decide_winner()
		return
	}

	if o.State == BLACK_TURN {
		o.State = WHITE_TURN
	} else {
		o.State = BLACK_TURN
	}
	o.update_valid_cells()
	// if the next player has no valid moves, skip their turn
	if len(o.GetValidMoves()) == 0 {
		if o.State == BLACK_TURN {
			o.State = WHITE_TURN
		} else {
			o.State = BLACK_TURN
		}
		o.update_valid_cells()
		if len(o.GetValidMoves()) == 0 {
			o.decide_winner()
		}
	}
}

func (o *Othello) is_board_full() bool {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if o.Board[y][x] == EMPTY || o.Board[y][x] == VALID {
				return false
			}
		}
	}
	return true
}

func (o *Othello) decide_winner() {
	if o.BlackScore > o.WhiteScore {
		o.State = BLACK_WON
	} else if o.BlackScore < o.WhiteScore {
		o.State = WHITE_WON
	} else {
		o.State = DRAW
	}
}

func (o *Othello) update_valid_cells() {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if o.Board[y][x] == VALID {
				o.Board[y][x] = EMPTY
			}
			if o.Board[y][x] == EMPTY && len(o.flipped_cells([2]int{x, y})) > 0 {
				o.Board[y][x] = VALID
			}
		}
	}
}

func (o *Othello) flipped_cells(position [2]int) [][2]int {
	var player Cell
	var opponent Cell
	if o.State == BLACK_TURN {
		player = BLACK
		opponent = WHITE
	} else {
		player = WHITE
		opponent = BLACK
	}
	flipped := [][2]int{}
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {1, 1}, {1, -1}, {-1, 0}, {-1, 1}, {-1, -1}}
	for _, direction := range directions {
		flipped = append(flipped, o.flipped_cells_in_direction(position[0], position[1], direction[0], direction[1], player, opponent)...)
	}
	return flipped
}

func (o *Othello) flipped_cells_in_direction(x, y, dx, dy int, player, opponent Cell) [][2]int {
	flipped := [][2]int{}
	x += dx
	y += dy
	for x >= 0 && x < 8 && y >= 0 && y < 8 && o.Board[y][x] == opponent {
		flipped = append(flipped, [2]int{x, y})
		x += dx
		y += dy
	}
	if !(x >= 0 && x < 8 && y >= 0 && y < 8) || o.Board[y][x] != player {
		return [][2]int{}
	}
	return flipped
}
