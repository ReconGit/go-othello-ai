package main

import (
	"fmt"
	"github.com/ReconGit/go-othello-ai/pkg/ui"
)

func main() {
	for {
		ui.PlayGame()
		// ask user if he wants to play again
		var input string
		for input != "y" && input != "n" {
			fmt.Printf("Play again? (y/n): ")
			fmt.Scanln(&input)
			if input == "n" {
				fmt.Println("\nThanks for playing!")
				return
			}
		}
	}
}
