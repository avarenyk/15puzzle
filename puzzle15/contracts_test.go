package puzzle15

import (
	"testing"
)

func TestPointAction_IsValid(t *testing.T) {
	testCases := []struct {
		name        string
		pointAction PointAction
		expected    bool
	}{
		{
			name:        "Valid PointAction within bounds 1",
			pointAction: PointAction{X: 0, Y: 0},
			expected:    true,
		},
		{
			name:        "Valid PointAction within bounds 2",
			pointAction: PointAction{X: 1, Y: 2},
			expected:    true,
		},
		{
			name:        "Valid PointAction within bounds 3",
			pointAction: PointAction{X: 2, Y: 3},
			expected:    true,
		},
		{
			name:        "Invalid PointAction with X out of bounds",
			pointAction: PointAction{X: 4, Y: 1},
			expected:    false,
		},
		{
			name:        "Invalid PointAction with Y out of bounds",
			pointAction: PointAction{X: 2, Y: -1},
			expected:    false,
		},
		{
			name:        "Invalid PointAction with both X and Y out of bounds",
			pointAction: PointAction{X: 5, Y: 5},
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := tc.pointAction.IsValid()
			if isValid != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, isValid)
			}
			if !isValid && err == nil {
				t.Errorf("expected an error but got none")
			}
		})
	}
}

func TestArrowAction_IsValid(t *testing.T) {
	testCases := []struct {
		name        string
		arrowAction ArrowAction
		expected    bool
	}{
		{
			name:        "Valid ArrowAction with Up",
			arrowAction: ArrowAction{Up: true},
			expected:    true,
		},
		{
			name:        "Valid ArrowAction with Right",
			arrowAction: ArrowAction{Right: true},
			expected:    true,
		},
		{
			name:        "Valid ArrowAction with Down",
			arrowAction: ArrowAction{Down: true},
			expected:    true,
		},
		{
			name:        "Valid ArrowAction with Left",
			arrowAction: ArrowAction{Left: true},
			expected:    true,
		},
		{
			name:        "Valid ArrowAction with multiple directions",
			arrowAction: ArrowAction{Left: true, Down: true},
			expected:    false,
		},
		{
			name:        "Invalid ArrowAction with no direction",
			arrowAction: ArrowAction{},
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := tc.arrowAction.IsValid()
			if isValid != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, isValid)
			}
			if !isValid && err == nil {
				t.Errorf("expected an error but got none")
			}
		})
	}
}
