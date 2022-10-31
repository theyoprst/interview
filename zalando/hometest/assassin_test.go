package hometest

import (
	"math/rand"
	"strings"
	"testing"

	"interviews/zalando/hometest/solution"
)

func TestAssassin(t *testing.T) {
	cases := []struct {
		board []string
		expected bool
	} {
		{
			[]string{"A"},
			true,
		},
		{
			[]string{
				"...X",
				".X.>",
				"AX..",
			},
			true,
		},
		{
			[]string {
				"....",
				"..v.",
				".>..",
				"A...",
			},
			false,
		},
		{
			[]string{">A"},
			false,
		},
	}
	for _, test := range cases {
		t.Run(strings.Join(test.board, "/"), func(t *testing.T) {
			res := solution.CanAssassinEscape(test.board)
			if res != test.expected {
				t.Fatalf("got %t, want %t", res, test.expected)
			}
		})
	}
}

type CellPicker struct {
	wallChance float64
	guardChance float64
}

func NewCellPicker(wallChance, guardChance float64) CellPicker {
	if wallChance + guardChance * 4 > 1.0 {
		panic("chanced are not configured properly")
	}
	return CellPicker {
		wallChance: wallChance,
		guardChance: guardChance,
	}
}

var guards = []byte{
	solution.GuardLeft,
	solution.GuardRight,
	solution.GuardUp,
	solution.GuardDown,
}

func (p CellPicker) Pick() byte {
	r := rand.Float64()
	if r < p.wallChance {
		return solution.Wall
	}
	for _, guard := range guards {
		r -= p.wallChance
		if r < p.guardChance {
			return guard
		}
	}
	return solution.Empty
}

func GenRandomField(n int) []string {
	rand.Seed(1)
	picker := NewCellPicker(0.2, 0.01)
	var field []string
	for i := 0; i < n; i++ {
		var row []byte
		for j := 0; j < n; j++ {
			cell := picker.Pick()
			if i == 0 && j == 0 {
				cell = solution.Assassin
			} else if i == n-1 && j == n-1 {
				cell = solution.Empty
			}
			row = append(row, cell)
		}
		field = append(field, string(row))
	}
	return field
}

func GetGuardsOnBorders(n int) []string {
	var field [][]byte
	for i := 0; i < n; i++ {
		field = append(field, []byte{})
		for j := 0; j < n; j++ {
			field[i] = append(field[i], solution.Empty)
		}
	}
	for i := 0; i < n; i++ {
		field[0][i] = solution.GuardDown
		field[n-1][i] = solution.GuardUp
		field[i][0] = solution.GuardRight
		field[i][n-1] = solution.GuardLeft
	}
	var f []string
	for i := 0; i < n; i++ {
		f = append(f, string(field[i]))
	}
	return f
}

func BenchmarkAssassin(b *testing.B) {
	field := GetGuardsOnBorders(500)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = solution.CanAssassinEscape(field)
	}
}