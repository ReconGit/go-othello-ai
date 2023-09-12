package main

import (
	"fmt"
	"github.com/ReconGit/go-othello-ai/pkg/game"
)

func main() {
	for {
		game.PlayGame()
		// ask user if he wants to play again
		var input string
		for input != "y" && input != "n" {
			fmt.Printf("Do you want to play again? (y/n): ")
			fmt.Scanln(&input)
			if input == "n" {
				fmt.Println("\nThanks for playing!")
				return
			}
		}
	}
}
