package main

import "testing"

func TestAreBishopsAttackingSimple(t *testing.T) {
	tests := []struct {
		name     string
		start    [2]int
		end      [2]int
		expected bool
	}{
		{
			name:     "White bishop on A1, Black bishop on C3",
			start:    [2]int{0, 0},
			end:      [2]int{2, 2},
			expected: true,
		},
		{
			name:     "White bishop on C1, Black bishop on A3",
			start:    [2]int{2, 0},
			end:      [2]int{0, 2},
			expected: true,
		},
		{
			name:     "White bishop on A1, Black bishop on B1",
			start:    [2]int{0, 0},
			end:      [2]int{1, 0},
			expected: false,
		},
		{
			name:     "White bishop on A1, Black bishop on B2",
			start:    [2]int{0, 0},
			end:      [2]int{1, 1},
			expected: true,
		},
		{
			name:     "White bishop on B2, Black bishop on B2",
			start:    [2]int{1, 1},
			end:      [2]int{1, 1},
			expected: false,
		},
		{
			name:     "White bishop on C1, Black bishop on G5",
			start:    [2]int{2, 0}, // C1
			end:      [2]int{6, 4}, // G5
			expected: true,
		},
		{
			name:     "White bishop on C2, Black bishop on C4",
			start:    [2]int{2, 1}, // C2
			end:      [2]int{2, 3}, // C4
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := areBishopsAttacking(tt.start[0], tt.start[1], tt.end[0], tt.end[1])
			if actual != tt.expected {
				t.Errorf("areBishopsAttacking(%v, %v, %v, %v) got %v, want %v", tt.start[0], tt.start[1], tt.end[0], tt.end[1], actual, tt.expected)
			}
		})
	}
}
