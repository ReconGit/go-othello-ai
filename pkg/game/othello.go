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

// NewOthello returns a new Othello game.
func NewOthello() Othello {
	// initialize board
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

// MakeMove makes a move on the board.
func (o *Othello) MakeMove(position [2]int) {
	// sanity checks
	if o.State != BLACK_TURN && o.State != WHITE_TURN {
		panic("Cannot make move because the game is over.")
	}
	if o.Board[position[0]][position[1]] != VALID {
		panic("Position is not valid.")
	}
	// update board
	reverse := BLACK
	if o.State == WHITE_TURN {
		reverse = WHITE
	}
	o.Board[position[0]][position[1]] = reverse
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
				valid_moves = append(valid_moves, [2]int{y, x})
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
	// switch turn
	if o.State == BLACK_TURN {
		o.State = WHITE_TURN
	} else {
		o.State = BLACK_TURN
	}
	// update valid cells
	o.update_valid_cells()
	if len(o.GetValidMoves()) == 0 {
		// switch turn again
		if o.State == BLACK_TURN {
			o.State = WHITE_TURN
		} else {
			o.State = BLACK_TURN
		}
		// update valid cells again
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
	player := BLACK
	if o.State == WHITE_TURN {
		player = WHITE
	}
	opponent := BLACK
	if o.State == WHITE_TURN {
		opponent = WHITE
	}
	flipped := [][2]int{}
	x := position[0]
	y := position[1]
	for _, direction := range [][2]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}} {
		flipped = append(flipped, o.flipped_cells_in_direction(x, y, direction[0], direction[1], player, opponent)...)
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
	if !(x >= 0 && x < 8 && y >= 0 && y < 8) || o.Board[y][x] == player {
		return [][2]int{}
	}
	return flipped
}
