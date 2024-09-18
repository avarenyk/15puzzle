package puzzle15

import (
	"testing"
)

var arrowActionTestCases = []struct {
	startField    Field
	arrowAction   ArrowAction
	expectedField Field
}{
	// Move blank up
	{
		startField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 0},
			},
			BlankX: 3,
			BlankY: 3,
		},
		arrowAction: ArrowAction{Down: true},
		expectedField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 0},
				{13, 14, 15, 12},
			},
			BlankX: 2,
			BlankY: 3,
		},
	},
	// Move blank right
	{
		startField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 0, 7, 8},
				{9, 6, 10, 12},
				{13, 14, 15, 11},
			},
			BlankX: 1,
			BlankY: 1,
		},
		arrowAction: ArrowAction{Left: true},
		expectedField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 7, 0, 8},
				{9, 6, 10, 12},
				{13, 14, 15, 11},
			},
			BlankX: 1,
			BlankY: 2,
		},
	},
	// Move blank left
	{
		startField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 0, 12},
				{13, 14, 15, 11},
			},
			BlankX: 2,
			BlankY: 2,
		},
		arrowAction: ArrowAction{Right: true},
		expectedField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 0, 10, 12},
				{13, 14, 15, 11},
			},
			BlankX: 2,
			BlankY: 1,
		},
	},
	// Move blank right again
	{
		startField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 0, 7, 8},
				{9, 6, 10, 12},
				{13, 14, 15, 11},
			},
			BlankX: 1,
			BlankY: 1,
		},
		arrowAction: ArrowAction{Up: true},
		expectedField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 0, 10, 12},
				{13, 14, 15, 11},
			},
			BlankX: 2,
			BlankY: 1,
		},
	},
}

var testCases = []struct {
	startField    Field
	pointAction   PointAction
	expectedField Field
}{
	// Test case 1: Swap blank with 15 at (3,3)
	{
		startField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 0, 15},
			},
			BlankX: 3,
			BlankY: 2,
		},
		pointAction: PointAction{X: 3, Y: 3},
		expectedField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 0},
			},
			BlankX: 3,
			BlankY: 3,
		},
	},
	// Test case 2: Swap blank with 11 at (2,2)
	{
		startField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 0, 15},
			},
			BlankX: 3,
			BlankY: 2,
		},
		pointAction: PointAction{X: 2, Y: 2},
		expectedField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 0, 12},
				{13, 14, 11, 15},
			},
			BlankX: 2,
			BlankY: 2,
		},
	},
	// Test case 3: Swap blank with 14 at (3,1)
	{
		startField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 0, 14, 15},
			},
			BlankX: 3,
			BlankY: 1,
		},
		pointAction: PointAction{X: 3, Y: 2},
		expectedField: Field{
			F: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 0, 15},
			},
			BlankX: 3,
			BlankY: 2,
		},
	},
}

func TestApplyPointAction_SuccessfulSwaps(t *testing.T) {
	// Iterate through test cases
	for i, tc := range testCases {
		newField, err := applyPointAction(tc.pointAction, tc.startField)
		if err != nil {
			t.Errorf("Test case %d: Unexpected error: %v", i, err)
		}

		// Compare the result with the expected field
		if !fieldsEqual(newField, tc.expectedField) {
			t.Errorf("Test case %d: Expected field %+v, but got %+v", i, tc.expectedField, newField)
		}
	}
}

func TestApplyArrowAction_Directions(t *testing.T) {
	// Test cases for different movements

	// Run each test case
	for i, tc := range arrowActionTestCases {
		newField, err := applyArrowAction(tc.arrowAction, tc.startField)
		if err != nil {
			t.Errorf("Test case %d: Unexpected error: %v", i, err)
		}

		// Compare the result with the expected field
		if !fieldsEqual(newField, tc.expectedField) {
			t.Errorf("Test case %d: Expected field %+v, but got %+v", i, tc.expectedField, newField)
		}
	}
}

func TestApplyArrowAction_MoveDown_OutOfBounds(t *testing.T) {
	// Field with blank at (3,3)
	initialField := Field{
		F: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 0},
		},
		BlankX: 3,
		BlankY: 3,
	}

	// Move blank down (out of bounds)
	action := ArrowAction{Up: true}

	_, err := applyArrowAction(action, initialField)
	if err == nil {
		t.Fatalf("Expected out-of-bounds error, but got no error")
	}
}

func TestApplyPointAction_NotNeighbor(t *testing.T) {
	// Field with blank at (3,2)
	initialField := Field{
		F: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 0, 15},
		},
		BlankX: 3,
		BlankY: 2,
	}

	// Try to swap blank with a non-neighbor point (2,0)
	action := PointAction{X: 2, Y: 0}

	_, err := applyPointAction(action, initialField)
	if err == nil {
		t.Fatalf("Expected non-neighbor error, but got no error")
	}
}

// Helper function to compare two Field structs
func fieldsEqual(f1, f2 Field) bool {
	if f1.BlankX != f2.BlankX || f1.BlankY != f2.BlankY {
		return false
	}
	for i := range f1.F {
		for j := range f1.F[i] {
			if f1.F[i][j] != f2.F[i][j] {
				return false
			}
		}
	}
	return true
}
