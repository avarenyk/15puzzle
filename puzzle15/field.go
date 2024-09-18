package puzzle15

import (
	"math/rand"
	"slices"
)

func createRandomField() Field {
	// Create a sorted array from 0 to 15
	nums := make([]int, 16)
	for i := 0; i < 16; i++ {
		nums[i] = i
	}

	// Shuffle the array
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	// Create a 4x4 field and fill it with shuffled values
	field := Field{
		F: make([][]int, 4),
	}

	for i := range field.F {
		field.F[i] = make([]int, 4)
	}

	// Fill the field with shuffled values and find the blank (0)
	for i := 0; i < 16; i++ {
		x := i / 4
		y := i % 4
		field.F[x][y] = nums[i]

		if nums[i] == 0 {
			field.BlankX = x
			field.BlankY = y
		}
	}

	return field
}

func isFieldSolved(field Field) bool {
	// Flatten the 2D field into a 1D array
	flattened := make([]int, 0, 16)
	for i := range field.F {
		flattened = append(flattened, field.F[i]...)
	}

	// Define the expected sorted array
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}

	// Compare the flattened array with the expected sorted version
	return slices.Equal(flattened, expected)
}
