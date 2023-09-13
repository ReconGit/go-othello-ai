package game

import (
	"fmt"
)

const (
	BLACK_ANSI  = "\x1b[30m"
	WHITE_ANSI  = "\x1b[37m"
	YELLOW_ANSI = "\x1b[33m"
	GREEN_ANSI  = "\x1b[32m"
	RED_ANSI    = "\x1b[31m"
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

		var position [2]int
		if game.State == BLACK_TURN {
			position = user_move(game)
		} else {
			position = MctsMove(game, 1000)
		}
		fmt.Printf("      Move: %c%d\n", position[0]+65, position[1]+1)
		game.MakeMove(position)
	}
	fmt.Println("\n     Game Over!")
	print_board(game.Board)
	print_score(game.BlackScore, game.WhiteScore)
	print_state(game.State)
	fmt.Println()
}

func user_move(game Othello) [2]int {
	var input string
	for {
		fmt.Printf("Enter move (eg. A1): ")
		_, err := fmt.Scanln(&input)
		if err != nil || len(input) != 2 {
			fmt.Printf("%sInvalid input.%s\n", RED_ANSI, RESET_ANSI)
			continue
		}
		x := int(input[0]) - 65
		y := int(input[1]) - 49
		// if the move is in valid move list, return it
		for _, move := range game.GetValidMoves() {
			if move[0] == x && move[1] == y {
				return [2]int{x, y}
			}
		}
		fmt.Printf("%sInvalid move.%s\n", RED_ANSI, RESET_ANSI)
	}
}

func print_board(game_board [8][8]Cell) {
	fmt.Println("   A B C D E F G H")
	for y := 0; y < 8; y++ {
		fmt.Printf("%d |", y+1)
		for x := 0; x < 8; x++ {
			switch game_board[y][x] {
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
