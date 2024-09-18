package puzzle15

import (
	"fmt"
)

type GameStatus int

const (
	OnGoing GameStatus = iota
	Solved
)

type PuzzleGamer interface {
	GetGame() PuzzleGame
	RestartGame() PuzzleGame
	PointAction(PointAction) (PuzzleGame, error)
	ArrowAction(ArrowAction) (PuzzleGame, error)
}

type (
	PuzzleGame struct {
		Status     GameStatus
		Field      Field
		MovesCount int
	}

	Field struct {
		F [][]int

		BlankX, BlankY int
	}

	PointAction struct {
		X, Y int
	}

	ArrowAction struct {
		Up, Down, Left, Right bool
	}
)

func (p PointAction) IsValid() (bool, error) {
	// Check if X and Y coordinates are within the bounds of the 4x4 grid
	if p.X < 0 || p.X > 3 || p.Y < 0 || p.Y > 3 {
		return false, fmt.Errorf("point action is out of field 4x4 grid: %+v", p)
	}
	return true, nil
}

func (p ArrowAction) IsValid() (bool, error) {
	boolToInt := func(a bool) int {
		if a {
			return 1
		}
		return 0
	}
	sum := boolToInt(p.Up) + boolToInt(p.Down) + boolToInt(p.Right) + boolToInt(p.Left)
	if sum == 1 {
		return true, nil
	}
	return false, fmt.Errorf("arrow action should consist of exactly one direction, got: %+v", p)
}
