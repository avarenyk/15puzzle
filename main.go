package main

import (
	"fmt"
	"puzzle15/puzzle15"
	"puzzle15/terminalGUI"
)

func main() {
	game := puzzle15.NewPuzzleGame()
	termGUI, err := terminalGUI.InitTerminalGui(game)
	if err != nil {
		fmt.Printf("could not start game UI : %v\n", err)
	}

	// hold execution of main goroutine until game is finished
	<-termGUI.Run()
}
