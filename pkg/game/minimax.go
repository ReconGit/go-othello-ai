package game

import (
	"math"
	"math/rand"
)

// MinimaxMove returns the best move for the given game state using Minimax with alpha-beta pruning.
func MinimaxMove(game Othello, depth int) [2]int {
	possible_moves := game.GetValidMoves()
	possible_moves_len := len(possible_moves)
	
	if possible_moves_len == 0 {
		panic("Minimax: No valid moves!")
	}
	if possible_moves_len == 1 {
		return possible_moves[0]
	}
	round := get_round(game.Board)
	if round < 3 {
		return possible_moves[rand.Intn(possible_moves_len)]
	}

	if round >= 50 {
		depth += 10 // end game solver
	} else if round > 40 {
		depth += 2
	} else if round > 30 {
		depth += 1
	}
	my_turn := game.State
	best_move := possible_moves[0]
	best_value := math.MinInt32
	result_chan := make(chan [2]int, possible_moves_len)

	for i := range possible_moves {
		// parallelize each move in a goroutine to speed up the search
		go func(move_idx int) {
			simulation := game.DeepCopy()
			simulation.MakeMove(possible_moves[move_idx])
			value := minimax(simulation, my_turn, depth-1, math.MinInt32, math.MaxInt32)
			result_chan <- [2]int{value, move_idx}
		}(i)
	}
	for i := 0; i < possible_moves_len; i++ {
		result := <-result_chan
		if result[0] > best_value {
			best_value = result[0]
			best_move = possible_moves[result[1]]
		}
		if best_value >= 300 {
			// good enough move found
			// goroutines will finish in the background but we don't care about the results
			break
		}
	}
	return best_move
}

func minimax(game Othello, my_turn State, depth, alpha, beta int) int {
	state := game.State
	switch state {
	case BLACK_WON:
		if my_turn == BLACK_TURN {
			return 300
		} else {
			return -300
		}
	case WHITE_WON:
		if my_turn == WHITE_TURN {
			return 300
		} else {
			return -300
		}
	case DRAW:
		return 0
	}
	if depth <= 0 {
		return evaluate_board(game.Board, my_turn)
	}

	var best_value int
	if state == my_turn {
		best_value = math.MinInt32
	} else {
		best_value = math.MaxInt32
	}
	possible_moves := game.GetValidMoves()
	for i := range possible_moves {
		simulation := game.DeepCopy()
		simulation.MakeMove(possible_moves[i])
		value := minimax(simulation, my_turn, depth-1, alpha, beta)

		if state == my_turn {
			best_value = max(best_value, value)
			alpha = max(best_value, alpha)
		} else {
			best_value = min(best_value, value)
			beta = min(best_value, beta)
		}
		if alpha >= beta {
			break // prune
		}
	}
	return best_value
}

var REWARDS = [8][8]int{
	{80, -20, 20, 10, 10, 20, -20, 80},
	{-20, -40, -10, -10, -10, -10, -40, -20},
	{20, -10, 10, 0, 0, 10, -10, 20},
	{10, -10, 0, 5, 5, 0, -10, 10},
	{10, -10, 0, 5, 5, 0, -10, 10},
	{20, -10, 10, 0, 0, 10, -10, 20},
	{-20, -40, -10, -10, -10, -10, -40, -20},
	{80, -20, 20, 10, 10, 20, -20, 80},
}

func evaluate_board(game_board [8][8]Cell, my_turn State) int {
	score := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			switch game_board[y][x] {
			case BLACK:
				if my_turn == BLACK_TURN {
					score += REWARDS[y][x]
				} else {
					score -= REWARDS[y][x]
				}
			case WHITE:
				if my_turn == WHITE_TURN {
					score += REWARDS[y][x]
				} else {
					score -= REWARDS[y][x]
				}
			}
		}
	}
	return score
}

func get_round(game_board [8][8]Cell) int {
	round := -3
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if game_board[y][x] == BLACK || game_board[y][x] == WHITE {
				round++
			}
		}
	}
	return round
}
