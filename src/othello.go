package src

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
	board       [8][8]Cell
	state       State
	black_score int
	white_score int
}

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
		board:       board,
		state:       BLACK_TURN,
		black_score: 2,
		white_score: 2,
	}
}

func (o *Othello) MakeMove(position [2]int) {
	// sanity checks
	if o.state != BLACK_TURN && o.state != WHITE_TURN {
		panic("Cannot make move because the game is over.")
	}
	if o.board[position[0]][position[1]] != VALID {
		panic("Position is not valid.")
	}

	// update board
	reverse := BLACK
	if o.state == WHITE_TURN {
		reverse = WHITE
	}
	o.board[position[0]][position[1]] = reverse
	for _, cell := range o.flipped_cells(position) {
		o.board[cell[1]][cell[0]] = reverse
	}
	o.update_state()
}

func (o *Othello) GetValidMoves() [][2]int {
	if o.state != BLACK_TURN && o.state != WHITE_TURN {
		return [][2]int{}
	}
	valid_moves := [][2]int{}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if o.board[y][x] == VALID {
				valid_moves = append(valid_moves, [2]int{y, x})
			}
		}
	}
	return valid_moves
}

func (o *Othello) update_state() {
	o.black_score = 0
	o.white_score = 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if o.board[y][x] == BLACK {
				o.black_score++
			} else if o.board[y][x] == WHITE {
				o.white_score++
			}
		}
	}
	if o.is_board_full() || o.black_score == 0 || o.white_score == 0 {
		o.decide_winner()
		return
	}
	// switch turn
	if o.state == BLACK_TURN {
		o.state = WHITE_TURN
	} else {
		o.state = BLACK_TURN
	}
	// update valid cells
	o.update_valid_cells()
	if len(o.GetValidMoves()) == 0 {
		// switch turn again
		if o.state == BLACK_TURN {
			o.state = WHITE_TURN
		} else {
			o.state = BLACK_TURN
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
			if o.board[y][x] == EMPTY || o.board[y][x] == VALID {
				return false
			}
		}
	}
	return true
}

func (o *Othello) decide_winner() {
	if o.black_score > o.white_score {
		o.state = BLACK_WON
	} else if o.black_score < o.white_score {
		o.state = WHITE_WON
	} else {
		o.state = DRAW
	}
}

func (o *Othello) update_valid_cells() {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if o.board[y][x] == VALID {
				o.board[y][x] = EMPTY
			}
			if o.board[y][x] == EMPTY && len(o.flipped_cells([2]int{x, y})) > 0 {
				o.board[y][x] = VALID
			}
		}
	}
}

func (o *Othello) flipped_cells(position [2]int) [][2]int {
	player := BLACK
	if o.state == WHITE_TURN {
		player = WHITE
	}
	opponent := BLACK
	if o.state == WHITE_TURN {
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

func (o *Othello) flipped_cells_in_direction(x int, y int, dx int, dy int, player Cell, opponent Cell) [][2]int {
	flipped := [][2]int{}
	x += dx
	y += dy
	for x >= 0 && x < 8 && y >= 0 && y < 8 && o.board[y][x] == opponent {
		flipped = append(flipped, [2]int{x, y})
		x += dx
		y += dy
	}
	if !(x >= 0 && x < 8 && y >= 0 && y < 8) || o.board[y][x] == player {
		return [][2]int{}
	}
	return flipped
}
