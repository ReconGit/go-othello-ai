package game

import (
	"fmt"
)

const (
	BLACK_ANSI  = "\x1b[30m"
	WHITE_ANSI  = "\x1b[37m"
	YELLOW_ANSI = "\x1b[33m"
	GREEN_ANSI  = "\x1b[32m"
	RESET_ANSI  = "\x1b[0m"
)

func PlayGame() {
	fmt.Println("Welcome to Othello!")

	game := NewOthello()
	round := 0
	for game.State == BLACK_TURN || game.State == WHITE_TURN {
		round++
		fmt.Printf("\n%s      Round %d%s\n", GREEN_ANSI, round, RESET_ANSI)
		print_board(game.Board)
		print_score(game.BlackScore, game.WhiteScore)
		print_state(game.State)

		var move [2]int
		if game.State == BLACK_TURN {
			move = MinimaxMove(game, 1)
		} else {
			move = MctsMove(game, 100)
		}
		fmt.Printf("      Move: %c%d\n", move[0]+65, move[1]+1)
		game.MakeMove(move)
	}
	fmt.Println("\n     Game Over!")
	print_board(game.Board)
	print_score(game.BlackScore, game.WhiteScore)
	print_state(game.State)
	fmt.Println()
}

func print_board(game_board [8][8]Cell) {
	fmt.Println("   A B C D E F G H")
	for y := 0; y < 8; y++ {
		fmt.Printf("%d |", y+1)
		for x := 0; x < 8; x++ {
			switch game_board[x][y] {
			case EMPTY:
				fmt.Printf(" ")
			case BLACK:
				fmt.Printf("%s●%s", BLACK_ANSI, RESET_ANSI)
			case WHITE:
				fmt.Printf("%s●%s", WHITE_ANSI, RESET_ANSI)
			case VALID:
				fmt.Printf("%s*%s", YELLOW_ANSI, RESET_ANSI)
			}
			fmt.Printf("|")
		}
		fmt.Printf("\n")
	}
}

func print_score(BlackScore, WhiteScore int) {
	fmt.Printf("%sBlack: %d %s| White: %d\n", BLACK_ANSI, BlackScore, RESET_ANSI, WhiteScore)
}

func print_state(State State) {
	switch State {
	case BLACK_TURN:
		fmt.Printf("%s     BLACK turn%s\n", BLACK_ANSI, RESET_ANSI)
	case WHITE_TURN:
		fmt.Printf("%s     WHITE turn%s\n", WHITE_ANSI, RESET_ANSI)
	case BLACK_WON:
		fmt.Printf("%s     BLACK won%s\n", BLACK_ANSI, RESET_ANSI)
	case WHITE_WON:
		fmt.Printf("%s     WHITE won%s\n", WHITE_ANSI, RESET_ANSI)
	case DRAW:
		fmt.Println("        DRAW")
	}
}
