package puzzle15

import (
	"testing"
)

const testRunForRandFunc = 100

// Test if the field contains all values from 0 to 15
func TestCreateRandomField_ContainsAllValues(t *testing.T) {
	for i := 0; i < testRunForRandFunc; i++ {
		field := createRandomField()

		// Create a map to track occurrences of each value
		valueMap := make(map[int]bool)

		// Loop through the field and mark each value as found
		for i := range field.F {
			for j := range field.F[i] {
				val := field.F[i][j]
				if val < 0 || val > 15 {
					t.Errorf("Field contains out of range value: %d", val)
				}
				valueMap[val] = true
			}
		}

		// Check that all values from 0 to 15 are present
		for val, found := range valueMap {
			if !found {
				t.Errorf("Field does not contain value: %d", val)
			}
		}
	}
}

// Test if the blank coordinates (0) are correctly set
func TestCreateRandomField_BlankCoordinates(t *testing.T) {
	for i := 0; i < testRunForRandFunc; i++ {
		field := createRandomField()

		if field.F[field.BlankX][field.BlankY] != 0 {
			t.Errorf("Expected blank at (%d, %d), but got %d", field.BlankX, field.BlankY, field.F[field.BlankX][field.BlankY])
		}
	}
}

// Test if the field dimensions are correct (4x4 grid)
func TestCreateRandomField_FieldDimensions(t *testing.T) {
	for i := 0; i < testRunForRandFunc; i++ {
		field := createRandomField()

		if len(field.F) != 4 {
			t.Errorf("Expected field to have 4 rows, but got %d", len(field.F))
		}

		for i := range field.F {
			if len(field.F[i]) != 4 {
				t.Errorf("Expected row %d to have 4 columns, but got %d", i, len(field.F[i]))
			}
		}
	}
}

func TestIsFieldSolved(t *testing.T) {
	// Create a solved field
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

	if !isFieldSolved(solvedField) {
		t.Errorf("Expected field to be solved, but it's not")
	}

	// Create an unsolved field
	unsolvedField := createRandomField()
	if isFieldSolved(unsolvedField) {
		t.Errorf("Expected field to be unsolved, but it's solved")
	}
}
