package mathutil

import "testing"

func TestAdd(t *testing.T) {
	// Happy path cases
	tests := []struct {
		x, y, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, -1, -2},
		{100, 200, 300},
		{-100, 100, 0},
	}

	for _, test := range tests {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", test.x, test.y, result, test.expected)
		}
	}

	// Edge cases
	edgeTests := []struct {
		x, y, expected int
	}{
		{int(^uint(0) >> 1), 1, int(^uint(0) >> 1) + 1}, // Max int value
		{int(^uint(0) >> 1), -1, int(^uint(0) >> 1)},    // Max int value - 1
		{int(^uint(0) >> 1), 0, int(^uint(0) >> 1)},     // Max int value + 0
		{int(^uint(0) >> 1) - 1, 1, int(^uint(0) >> 1)}, // Max int value - 1 + 1
	}

	for _, test := range edgeTests {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", test.x, test.y, result, test.expected)
		}
	}
}