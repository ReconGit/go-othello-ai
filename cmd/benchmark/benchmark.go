package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ReconGit/go-othello-ai/pkg/game"
)

const (
	GAMES           = 100
	MINIMAX_DEPTH   = 3
	MCTS_ITERATIONS = 100

	MAGENTA_ANSI = "\x1b[35m"
	BLUE_ANSI    = "\x1b[34m"
	RESET_ANSI   = "\x1b[0m"
)

func main() {
	run_benchmarks()
}

func run_benchmarks() {
	fmt.Printf("%sRunning benchmarks...%s\n", MAGENTA_ANSI, RESET_ANSI)
	start_time := time.Now()

	//
	fmt.Printf("%s\nRandom vs Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmark_game(random_move, random_move, 0, 0)

	fmt.Printf("%s\nBLACK Minimax vs WHITE Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmark_game(game.MinimaxMove, random_move, MINIMAX_DEPTH, 0)

	fmt.Printf("%s\nWHITE Minimax vs BLACK Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmark_game(random_move, game.MinimaxMove, 0, MINIMAX_DEPTH)

	fmt.Printf("%s\nMinimax vs Minimax%s\n", BLUE_ANSI, RESET_ANSI)
	benchmark_game(game.MinimaxMove, game.MinimaxMove, MINIMAX_DEPTH, MINIMAX_DEPTH)

	fmt.Printf("%s\nBLACK MCTS vs WHITE Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmark_game(game.MctsMove, random_move, MCTS_ITERATIONS, 0)

	fmt.Printf("%s\nWHITE MCTS vs BLACK Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmark_game(random_move, game.MctsMove, 0, MCTS_ITERATIONS)

	fmt.Printf("%s\nMCTS vs MCTS%s\n", BLUE_ANSI, RESET_ANSI)
	benchmark_game(game.MctsMove, game.MctsMove, MCTS_ITERATIONS, MCTS_ITERATIONS)

	fmt.Printf("%s\nBLACK Minimax vs WHITE MCTS%s\n", BLUE_ANSI, RESET_ANSI)
	benchmark_game(game.MinimaxMove, game.MctsMove, MINIMAX_DEPTH, MCTS_ITERATIONS)

	fmt.Printf("%s\nBLACK MCTS vs WHITE Minimax%s\n", BLUE_ANSI, RESET_ANSI)
	benchmark_game(game.MctsMove, game.MinimaxMove, MCTS_ITERATIONS, MINIMAX_DEPTH)
	//

	end_time := time.Since(start_time).Seconds()
	fmt.Printf("%s\nTotal elapsed time: %.2f\n%s", MAGENTA_ANSI, end_time, RESET_ANSI)
}

func benchmark_game(
	black_ai func(game.Othello, int) [2]int,
	white_ai func(game.Othello, int) [2]int,
	black_iterations int,
	white_iterations int,
) {
	var black_wins int
	var white_wins int
	var draws int

	start := time.Now()
	for i := 0; i < GAMES; i++ {
		fmt.Printf("  game: %d/%d, elapsed time: %.2fs\r", i+1, GAMES, time.Since(start).Seconds())

		othello := game.NewOthello()
		for othello.State == game.BLACK_TURN || othello.State == game.WHITE_TURN {
			var move [2]int
			if othello.State == game.BLACK_TURN {
				move = black_ai(othello, black_iterations)
			} else {
				move = white_ai(othello, white_iterations)
			}
			othello.MakeMove(move)
		}
		if othello.State == game.BLACK_WON {
			black_wins++
		} else if othello.State == game.WHITE_WON {
			white_wins++
		} else {
			draws++
		}
	}
	fmt.Printf("  elapsed time: %.2fs                     \n", time.Since(start).Seconds())
	fmt.Printf("    BLACK wins: %d %.0f%%\n", black_wins, (float32(black_wins)/float32(GAMES))*100)
	fmt.Printf("    WHITE wins: %d %.0f%%\n", white_wins, (float32(white_wins)/float32(GAMES))*100)
	fmt.Printf("         Draws: %d %.0f%%\n", draws, (float32(draws)/float32(GAMES))*100)
}

func random_move(game game.Othello, _dummy_iterations int) [2]int {
	possible_moves := game.GetValidMoves()
	return possible_moves[rand.Intn(len(possible_moves))]
}
