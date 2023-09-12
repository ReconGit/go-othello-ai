package src

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	GAMES           = 10
	MINIMAX_DEPTH   = 1
	MCTS_ITERATIONS = 100

	MAGENTA_ANSI = "\x1b[35m"
	BLUE_ANSI    = "\x1b[34m"
)

func RunBenchmarks() {
	fmt.Printf("%sRunning benchmarks...%s\n", MAGENTA_ANSI, RESET_ANSI)
	start_time := time.Now()

	fmt.Printf("%s\nRandom vs Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmarkGame(random_move, random_move, 0, 0)

	fmt.Printf("%s\nBLACK Minimax vs WHITE Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmarkGame(MinimaxMove, random_move, MINIMAX_DEPTH, 0)

	fmt.Printf("%s\nWHITE Minimax vs BLACK Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmarkGame(random_move, MinimaxMove, 0, MINIMAX_DEPTH)

	fmt.Printf("%s\nMinimax vs Minimax%s\n", BLUE_ANSI, RESET_ANSI)
	benchmarkGame(MinimaxMove, MinimaxMove, MINIMAX_DEPTH, MINIMAX_DEPTH)

	// fmt.Printf("%s\nBLACK MCTS vs WHITE Random%s\n", BLUE_ANSI, RESET_ANSI)
	// benchmarkGame(MctsMove, random_move, MCTS_ITERATIONS, 0)

	// fmt.Printf("%s\nWHITE MCTS vs BLACK Random%s\n", BLUE_ANSI, RESET_ANSI)
	// benchmarkGame(random_move, MctsMove, 0, MCTS_ITERATIONS)

	// fmt.Printf("%s\nMCTS vs MCTS%s\n", BLUE_ANSI, RESET_ANSI)
	// benchmarkGame(MctsMove, MctsMove, MCTS_ITERATIONS, MCTS_ITERATIONS)

	// fmt.Printf("%s\nBLACK Minimax vs WHITE MCTS%s\n", BLUE_ANSI, RESET_ANSI)
	// benchmarkGame(MinimaxMove, MctsMove, MINIMAX_DEPTH, MCTS_ITERATIONS)

	// fmt.Printf("%s\nBLACK MCTS vs WHITE Minimax%s\n", BLUE_ANSI, RESET_ANSI)
	// benchmarkGame(MctsMove, MinimaxMove, MCTS_ITERATIONS, MINIMAX_DEPTH)

	end_time := time.Since(start_time).Seconds()
	fmt.Printf("%s\nTotal elapsed time: %.2f\n%s", MAGENTA_ANSI, end_time, RESET_ANSI)
}

func benchmarkGame(
	black_ai func(Othello, int) [2]int,
	white_ai func(Othello, int) [2]int,
	black_iterations int,
	white_iterations int,
) {
	var black_wins int
	var white_wins int
	var draws int

	start := time.Now()
	for i := 0; i < GAMES; i++ {
		fmt.Printf("  game: %d/%d, elapsed time: %.2fs\r", i+1, GAMES, time.Since(start).Seconds())

		game := NewOthello()
		for game.state == BLACK_TURN || game.state == WHITE_TURN {
			var move [2]int
			if game.state == BLACK_TURN {
				move = black_ai(game, black_iterations)
			} else {
				move = white_ai(game, white_iterations)
			}
			game.MakeMove(move)
		}
		if game.state == BLACK_WON {
			black_wins++
		} else if game.state == WHITE_WON {
			white_wins++
		} else {
			draws++
		}
	}
	fmt.Printf("  elapsed time: %.2fs                     \n", time.Since(start).Seconds())
	fmt.Printf("    BLACK wins: %d %.1f%%\n", black_wins, (float32(black_wins)/float32(GAMES))*100)
	fmt.Printf("    WHITE wins: %d %.1f%%\n", white_wins, (float32(white_wins)/float32(GAMES))*100)
	fmt.Printf("         Draws: %d %.1f%%\n", draws, (float32(draws)/float32(GAMES))*100)
}

func random_move(game Othello, iterations int) [2]int {
	possible_moves := game.GetValidMoves()
	return possible_moves[rand.Intn(len(possible_moves))]
}
