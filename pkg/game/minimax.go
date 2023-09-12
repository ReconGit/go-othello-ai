package game

import (
	"github.com/huandu/go-clone/generic"
	"math"
	"math/rand"
)

// MinimaxMove returns the best move for the given game state using Minimax with alpha-beta pruning
func MinimaxMove(game Othello, depth int) [2]int {
	possible_moves := game.GetValidMoves()
	if len(possible_moves) == 0 {
		panic("Minimax: No valid moves!")
	}
	// if only one move, return it
	if len(possible_moves) == 1 {
		return possible_moves[0]
	}
	// if round is less than 3, return random move
	round := get_round(game.Board)
	if round < 3 {
		return possible_moves[rand.Intn(len(possible_moves))]
	}
	// increase depth based on round
	if round >= 50 {
		depth += 10
	} else if round > 40 {
		depth += 2
	} else if round > 30 {
		depth += 1
	}

	my_turn := game.State
	best_move := possible_moves[0]
	best_value := math.MinInt32
	for _, move := range possible_moves {
		simulation := clone.Clone(game)
		simulation.MakeMove(move)
		value := minimax(simulation, my_turn, depth-1, math.MinInt32, math.MaxInt32)
		if value > best_value {
			best_move = move
			best_value = value
		}
		if value >= 300 {
			break
		}
	}
	return best_move
}

func minimax(game Othello, my_turn State, depth, alpha, beta int) int {
	State := game.State
	if State == BLACK_WON {
		if my_turn == BLACK_TURN {
			return 300
		} else {
			return -300
		}
	} else if State == WHITE_WON {
		if my_turn == WHITE_TURN {
			return 300
		} else {
			return -300
		}
	} else if State == DRAW {
		return 0
	}
	// if depth is 0, return heuristic Board evaluation
	if depth <= 0 {
		return evaluate_board(game.Board, my_turn)
	}

	possible_moves := game.GetValidMoves()
	best_value := math.MaxInt32
	if State == my_turn {
		best_value = math.MinInt32
	}
	for _, move := range possible_moves {
		simulation := clone.Clone(game)
		simulation.MakeMove(move)
		value := minimax(simulation, my_turn, depth-1, alpha, beta)
		// alpha-beta pruning
		if State == my_turn {
			best_value = max(best_value, value)
			alpha = max(alpha, best_value)
		} else {
			best_value = min(best_value, value)
			beta = min(beta, best_value)
		}
		if alpha >= beta {
			break
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
			cell := game_board[y][x]
			if cell == BLACK {
				if my_turn == BLACK_TURN {
					score += REWARDS[y][x]
				} else {
					score -= REWARDS[y][x]
				}
			} else if cell == WHITE {
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
