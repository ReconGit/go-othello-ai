package ui

import (
	"fmt"
	"github.com/ReconGit/go-othello-ai/pkg/ai"
	"github.com/ReconGit/go-othello-ai/pkg/game"
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

	// initialize othello
	othello := game.NewOthello()
	round := 0
	for othello.State == game.BLACK_TURN || othello.State == game.WHITE_TURN {
		round++
		fmt.Printf("\n%s      Round %d%s\n", GREEN_ANSI, round, RESET_ANSI)
		print_board(othello.Board)
		print_score(othello.BlackScore, othello.WhiteScore)
		print_state(othello.State)

		// get move
		var move [2]int
		if othello.State == game.BLACK_TURN {
			move = ai.MinimaxMove(othello, 1)
		} else {
			move = ai.MctsMove(othello, 100)
		}
		fmt.Printf("      Move: %c%d\n", move[0]+65, move[1]+1)
		othello.MakeMove(move)
	}
	fmt.Println("\n     Game Over!")
	print_board(othello.Board)
	print_score(othello.BlackScore, othello.WhiteScore)
	print_state(othello.State)
	fmt.Println()
}

func print_board(game_board [8][8]game.Cell) {
	fmt.Println("   A B C D E F G H")
	for y := 0; y < 8; y++ {
		fmt.Printf("%d |", y+1)
		for x := 0; x < 8; x++ {
			switch game_board[x][y] {
			case game.EMPTY:
				fmt.Printf(" ")
			case game.BLACK:
				fmt.Printf("%s●%s", BLACK_ANSI, RESET_ANSI)
			case game.WHITE:
				fmt.Printf("%s●%s", WHITE_ANSI, RESET_ANSI)
			case game.VALID:
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

func print_state(State game.State) {
	switch State {
	case game.BLACK_TURN:
		fmt.Printf("%s     BLACK turn%s\n", BLACK_ANSI, RESET_ANSI)
	case game.WHITE_TURN:
		fmt.Printf("%s     WHITE turn%s\n", WHITE_ANSI, RESET_ANSI)
	case game.BLACK_WON:
		fmt.Printf("%s     BLACK won%s\n", BLACK_ANSI, RESET_ANSI)
	case game.WHITE_WON:
		fmt.Printf("%s     WHITE won%s\n", WHITE_ANSI, RESET_ANSI)
	case game.DRAW:
		fmt.Println("        DRAW")
	}
}
