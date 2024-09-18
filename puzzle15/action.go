package puzzle15

import (
	"fmt"
)

func applyArrowAction(a ArrowAction, field Field) (Field, error) {
	// Get the current blank spot coordinates
	x, y := field.BlankX, field.BlankY

	// Determine the new coordinates after the move
	newX, newY := x, y

	if a.Up {
		newX = x + 1
	} else if a.Down {
		newX = x - 1
	} else if a.Left {
		newY = y + 1
	} else if a.Right {
		newY = y - 1
	}

	// Check if the new coordinates are within bounds (0 to 3)
	if newX < 0 || newX > 3 || newY < 0 || newY > 3 {
		return field, fmt.Errorf("move out of bounds")
	}

	// Swap the blank spot with the new position
	field.F[x][y], field.F[newX][newY] = field.F[newX][newY], field.F[x][y]

	// Update the blank spot's coordinates
	field.BlankX, field.BlankY = newX, newY

	return field, nil
}

func applyPointAction(p PointAction, field Field) (Field, error) {
	// Check if the selected point is a neighbor of the blank spot
	// Two points are neighbors if their difference in one coordinate is 1 and the other is the same
	if (p.X == field.BlankX && abs(p.Y-field.BlankY) == 1) || (p.Y == field.BlankY && abs(p.X-field.BlankX) == 1) {
		// Swap the blank spot with the selected point
		field.F[field.BlankX][field.BlankY], field.F[p.X][p.Y] = field.F[p.X][p.Y], field.F[field.BlankX][field.BlankY]

		// Update the blank spot coordinates
		field.BlankX, field.BlankY = p.X, p.Y

		return field, nil
	} else {
		// If the selected point is not a neighbor, return an error
		return field, fmt.Errorf("selected point is not a neighbor to the blank spot")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
