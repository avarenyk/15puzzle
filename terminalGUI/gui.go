package terminalGUI

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"puzzle15/puzzle15"
)

type GameTerminalGUI struct {
	game puzzle15.PuzzleGamer

	screen tcell.Screen

	finished chan bool
}

func InitTerminalGui(game puzzle15.PuzzleGamer) (*GameTerminalGUI, error) {
	var err error

	gameGUI := &GameTerminalGUI{
		game: game,

		finished: make(chan bool),
	}

	gameGUI.screen, err = tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("error creating screen: %v", err)
	}
	if err = gameGUI.screen.Init(); err != nil {
		return nil, fmt.Errorf("error initializing screen: %v", err)
	}

	return gameGUI, nil
}

func (gui *GameTerminalGUI) Run() chan bool {
	pzGame := gui.game.GetGame()
	go func() {
		for {
			gui.renderGame(pzGame)

			event := gui.screen.PollEvent()

			switch ev := event.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					gui.screen.Fini()
					close(gui.finished)
					return
				case tcell.KeyRune:
					if ev.Rune() == ' ' {
						pzGame = gui.game.RestartGame()
					}
				case tcell.KeyUp:
					pzGame, _ = gui.game.ArrowAction(puzzle15.ArrowAction{Up: true})
				case tcell.KeyDown:
					pzGame, _ = gui.game.ArrowAction(puzzle15.ArrowAction{Down: true})
				case tcell.KeyLeft:
					pzGame, _ = gui.game.ArrowAction(puzzle15.ArrowAction{Left: true})
				case tcell.KeyRight:
					pzGame, _ = gui.game.ArrowAction(puzzle15.ArrowAction{Right: true})
				}
			}
		}
	}()
	return gui.finished
}

func (gui *GameTerminalGUI) renderGame(puzzle puzzle15.PuzzleGame) {
	gui.screen.Clear()
	gui.renderInstruction()

	for i, yArr := range puzzle.Field.F {
		for j, val := range yArr {
			gui.printTile(j*4, i*2+4, val)
		}
	}

	if puzzle.Status == puzzle15.Solved {
		gui.renderWonStatement(puzzle.MovesCount)
	}
	gui.screen.Show()
}

func (gui *GameTerminalGUI) printTile(x, y, val int) {
	st := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
	str := fmt.Sprintf("%2d", val)
	if val == 0 {
		str = "  "
		st = tcell.StyleDefault.Background(tcell.ColorBlack)
	}
	for i, r := range str {
		gui.screen.SetContent(x+i, y, r, nil, st)
	}
}

func (gui *GameTerminalGUI) renderInstruction() {
	st := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	instructions1 := "Use arrows (Up, Down, Right, Left) to slide tiles into blank space."
	instructions2 := "ESC to quit a game and SPACE to restart a game"
	for i, r := range instructions1 {
		gui.screen.SetContent(i, 0, r, nil, st)
	}
	for i, r := range instructions2 {
		gui.screen.SetContent(i, 2, r, nil, st)
	}
}

func (gui *GameTerminalGUI) renderWonStatement(moves int) {
	st := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen)
	message := fmt.Sprintf("You WON in %d moves! Congratulations! Hit SPACE to start a new game", moves)
	for i, r := range message {
		gui.screen.SetContent(i, 12, r, nil, st)
	}

}
