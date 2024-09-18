package puzzle15

import (
	"sync"
)

func NewPuzzleGame() PuzzleGamer {
	gameEngine := &puzzleGameEngine{
		l: &sync.Mutex{},
	}
	gameEngine.RestartGame()

	return gameEngine
}

type puzzleGameEngine struct {
	l *sync.Mutex

	puzzleGame PuzzleGame
}

func (pg *puzzleGameEngine) GetGame() PuzzleGame {
	pg.l.Lock()
	defer pg.l.Unlock()

	return pg.puzzleGame
}

func (pg *puzzleGameEngine) RestartGame() PuzzleGame {
	pg.l.Lock()
	defer pg.l.Unlock()

	pg.puzzleGame = PuzzleGame{
		Status:     OnGoing,
		Field:      createRandomField(),
		MovesCount: 0,
	}

	return pg.puzzleGame
}

func (pg *puzzleGameEngine) PointAction(action PointAction) (PuzzleGame, error) {
	pg.l.Lock()
	defer pg.l.Unlock()

	// when puzzle is solved we do nothing on action inputs
	if pg.puzzleGame.Status == Solved {
		return pg.puzzleGame, nil
	}

	// there is input problems, raise an error
	valid, err := action.IsValid()
	if !valid {
		return pg.puzzleGame, err
	}

	// we increase move count and check if game is solved only when action was applied and changed field
	pg.puzzleGame.Field, err = applyPointAction(action, pg.puzzleGame.Field)
	if err == nil {
		pg.puzzleGame.MovesCount++
		if isFieldSolved(pg.puzzleGame.Field) {
			pg.puzzleGame.Status = Solved
		}
	}

	return pg.puzzleGame, nil
}

func (pg *puzzleGameEngine) ArrowAction(action ArrowAction) (PuzzleGame, error) {
	pg.l.Lock()
	defer pg.l.Unlock()

	// when puzzle is solved we do nothing on action inputs
	if pg.puzzleGame.Status == Solved {
		return pg.puzzleGame, nil
	}

	// there is input problems, raise an error
	valid, err := action.IsValid()
	if !valid {
		return pg.puzzleGame, err
	}

	// we increase move count and check if game is solved only when action was applied and changed field
	pg.puzzleGame.Field, err = applyArrowAction(action, pg.puzzleGame.Field)
	if err == nil {
		pg.puzzleGame.MovesCount++
		if isFieldSolved(pg.puzzleGame.Field) {
			pg.puzzleGame.Status = Solved
		}
	}

	return pg.puzzleGame, nil
}
