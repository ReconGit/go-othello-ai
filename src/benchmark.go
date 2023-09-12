package src

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	GAMES = 10
	MINIMAX_DEPTH = 2
	MAGENTA_ANSI = "\x1b[35m"
	BLUE_ANSI = "\x1b[34m"
)

func RunBenchmarks() {
	fmt.Printf("%sRunning benchmarks...%s\n", MAGENTA_ANSI, RESET_ANSI)
	start_time := time.Now()

	fmt.Printf("%sRandom vs. Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmarkGame(random_move, random_move, 0, 0)

	fmt.Printf("%sBLACK Minimax vs. WHITE Random%s\n", BLUE_ANSI, RESET_ANSI)
	benchmarkGame(MinimaxMove, random_move, MINIMAX_DEPTH, 0)

	fmt.Printf("%sWHITE Minimax vs. BLACK Minimax%s\n", BLUE_ANSI, RESET_ANSI)
	benchmarkGame(random_move, MinimaxMove, 0, MINIMAX_DEPTH)

	fmt.Printf("%sMinimax vs. Minimax%s\n", BLUE_ANSI, RESET_ANSI)
	benchmarkGame(MinimaxMove, MinimaxMove, MINIMAX_DEPTH, MINIMAX_DEPTH)

	end_time := time.Since(start_time)
	fmt.Printf("Benchmarks completed in %s\n", end_time)
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
        if game.state == BLACK_WON{
            black_wins++
        } else if game.state == WHITE_WON {
            white_wins++
        } else {
            draws++
        }
    }
    fmt.Printf("  elapsed time: %.2fs                     \n", time.Since(start).Seconds())
    fmt.Printf("    BLACK wins: %d %.1f%%\n", black_wins, (float32(black_wins) / float32(GAMES)) * 100)
    fmt.Printf("    WHITE wins: %d %.1f%%\n", white_wins, (float32(white_wins) / float32(GAMES)) * 100)
    fmt.Printf("         Draws: %d %.1f%%\n", draws, (float32(draws) / float32(GAMES)) * 100)
}

func random_move(game Othello, iterations int) [2]int {
	possible_moves := game.GetValidMoves()
	return possible_moves[rand.Intn(len(possible_moves))]
}