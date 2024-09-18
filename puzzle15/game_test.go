package puzzle15

import (
	"sync"
	"testing"
)

func TestArrowAction(t *testing.T) {
	// Initialize some example fields and actions
	initialField := Field{
		F: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 0, 12},
			{13, 14, 11, 15},
		},
		BlankX: 2,
		BlankY: 2,
	}

	solvedField := Field{
		F: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 0},
		},
		BlankX: 3,
		BlankY: 3,
	}

	testCases := []struct {
		name          string
		initialGame   PuzzleGame
		action        ArrowAction
		expectedGame  PuzzleGame
		expectedError bool
	}{
		{
			name: "Puzzle is solved",
			initialGame: PuzzleGame{
				Field:      solvedField,
				Status:     Solved,
				MovesCount: 5,
			},
			action: ArrowAction{Up: true},
			expectedGame: PuzzleGame{
				Field:      solvedField,
				Status:     Solved,
				MovesCount: 5,
			},
			expectedError: false,
		},
		{
			name: "Valid action updates field",
			initialGame: PuzzleGame{
				Field: Field{
					F: [][]int{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 0, 12},
						{13, 14, 11, 15},
					},
					BlankX: 2,
					BlankY: 2,
				},
				Status:     OnGoing,
				MovesCount: 5,
			},
			action: ArrowAction{Left: true},
			expectedGame: PuzzleGame{
				Field: Field{
					F: [][]int{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 12, 0},
						{13, 14, 11, 15},
					},
					BlankX: 2,
					BlankY: 3,
				},
				Status:     OnGoing,
				MovesCount: 6,
			},
			expectedError: false,
		},
		{
			name: "Valid action with no change",
			initialGame: PuzzleGame{
				Field: Field{
					F: [][]int{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 12, 0},
						{13, 14, 11, 15},
					},
					BlankX: 2,
					BlankY: 3,
				},
				Status:     OnGoing,
				MovesCount: 5,
			},
			action: ArrowAction{Left: true}, // No change for this action
			expectedGame: PuzzleGame{
				Field: Field{
					F: [][]int{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 12, 0},
						{13, 14, 11, 15},
					},
					BlankX: 2,
					BlankY: 3,
				},
				Status:     OnGoing,
				MovesCount: 5,
			},
			expectedError: false,
		},
		{
			name: "Invalid action",
			initialGame: PuzzleGame{
				Field:      initialField,
				Status:     OnGoing,
				MovesCount: 5,
			},
			action: ArrowAction{Up: true, Down: true}, // Invalid action
			expectedGame: PuzzleGame{
				Field:      initialField,
				Status:     OnGoing,
				MovesCount: 5,
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			engine := puzzleGameEngine{
				puzzleGame: tc.initialGame,
				l:          &sync.Mutex{},
			}
			result, err := engine.ArrowAction(tc.action)

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			if !fieldsEqual(result.Field, tc.expectedGame.Field) {
				t.Errorf("Expected field %+v, got %+v", tc.expectedGame.Field, result.Field)
			}

			if result.Status != tc.expectedGame.Status {
				t.Errorf("Expected status %v, got %v", tc.expectedGame.Status, result.Status)
			}

			if result.MovesCount != tc.expectedGame.MovesCount {
				t.Errorf("Expected MovesCount %d, got %d", tc.expectedGame.MovesCount, result.MovesCount)
			}
		})
	}
}

func TestPointAction(t *testing.T) {
	// Initialize some example fields and actions
	initialField := Field{
		F: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 0, 12},
			{13, 14, 11, 15},
		},
		BlankX: 2,
		BlankY: 2,
	}

	solvedField := Field{
		F: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 0},
		},
		BlankX: 3,
		BlankY: 3,
	}

	testCases := []struct {
		name          string
		initialGame   PuzzleGame
		action        PointAction
		expectedGame  PuzzleGame
		expectedError bool
	}{
		{
			name: "Puzzle is solved",
			initialGame: PuzzleGame{
				Field:      solvedField,
				Status:     Solved,
				MovesCount: 5,
			},
			action: PointAction{X: 2, Y: 3}, // Should have no effect
			expectedGame: PuzzleGame{
				Field:      solvedField,
				Status:     Solved,
				MovesCount: 5,
			},
			expectedError: false,
		},
		{
			name: "Valid action updates field",
			initialGame: PuzzleGame{
				Field:      initialField,
				Status:     OnGoing,
				MovesCount: 5,
			},
			action: PointAction{X: 2, Y: 3}, // Valid swap with 0
			expectedGame: PuzzleGame{
				Field: Field{
					F: [][]int{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 12, 0},
						{13, 14, 11, 15},
					},
					BlankX: 2,
					BlankY: 3,
				},
				Status:     OnGoing,
				MovesCount: 6,
			},
			expectedError: false,
		},
		{
			name: "Valid action with no change",
			initialGame: PuzzleGame{
				Field:      initialField,
				Status:     OnGoing,
				MovesCount: 5,
			},
			action: PointAction{X: 1, Y: 1}, // No change
			expectedGame: PuzzleGame{
				Field:      initialField,
				Status:     OnGoing,
				MovesCount: 5,
			},
			expectedError: false,
		},
		{
			name: "Valid action with no change, pointed to blank spot",
			initialGame: PuzzleGame{
				Field:      initialField,
				Status:     OnGoing,
				MovesCount: 5,
			},
			action: PointAction{X: 2, Y: 2}, // No change
			expectedGame: PuzzleGame{
				Field:      initialField,
				Status:     OnGoing,
				MovesCount: 5,
			},
			expectedError: false,
		},
		{
			name: "Invalid action",
			initialGame: PuzzleGame{
				Field:      initialField,
				Status:     OnGoing,
				MovesCount: 5,
			},
			action: PointAction{X: 4, Y: 4}, // Out of bounds
			expectedGame: PuzzleGame{
				Field:      initialField,
				Status:     OnGoing,
				MovesCount: 5,
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			engine := puzzleGameEngine{
				l:          &sync.Mutex{},
				puzzleGame: tc.initialGame,
			}
			result, err := engine.PointAction(tc.action)

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			if !fieldsEqual(result.Field, tc.expectedGame.Field) {
				t.Errorf("Expected field %+v, got %+v", tc.expectedGame.Field, result.Field)
			}

			if result.Status != tc.expectedGame.Status {
				t.Errorf("Expected status %v, got %v", tc.expectedGame.Status, result.Status)
			}

			if result.MovesCount != tc.expectedGame.MovesCount {
				t.Errorf("Expected MovesCount %d, got %d", tc.expectedGame.MovesCount, result.MovesCount)
			}
		})
	}
}
