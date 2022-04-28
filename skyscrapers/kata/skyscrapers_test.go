package kata

import "testing"

func TestSolvePuzzle(t *testing.T) {

	type testCase struct {
		name     string
		solve    func(clues []int, starting [][][]int) [][]int
		clues    []int
		starting [][][]int
		expected [][]int
	}

	cases := []testCase{
		{
			"4x4 full puzzle",
			SolvePuzzle4x4,
			[]int{
				0, 0, 1, 2,
				0, 2, 0, 0,
				0, 3, 0, 0,
				0, 1, 0, 0},
			nil,
			[][]int{
				{2, 1, 4, 3},
				{3, 4, 1, 2},
				{4, 2, 3, 1},
				{1, 3, 2, 4},
			},
		},
		{
			"4x4 almost completed puzzle",
			SolvePuzzle4x4,
			[]int{
				0, 0, 1, 2,
				0, 2, 0, 0,
				0, 3, 0, 0,
				0, 1, 0, 0},
			[][][]int{
				{{2, 3}, {1}, {4}, {2, 3}},
				{{1, 3}, {4}, {1}, {2}},
				{{4}, {2}, {3}, {1}},
				{{1, 3}, {3, 1}, {2}, {4}},
			},
			[][]int{
				{2, 1, 4, 3},
				{3, 4, 1, 2},
				{4, 2, 3, 1},
				{1, 3, 2, 4},
			},
		},
		{
			"6x6 full puzzle",
			SolvePuzzle6x6,
			[]int{3, 2, 2, 3, 2, 1,
				1, 2, 3, 3, 2, 2,
				5, 1, 2, 2, 4, 3,
				3, 2, 1, 2, 2, 4},
			nil,
			[][]int{{2, 1, 4, 3, 5, 6},
				{1, 6, 3, 2, 4, 5},
				{4, 3, 6, 5, 1, 2},
				{6, 5, 2, 1, 3, 4},
				{5, 4, 1, 6, 2, 3},
				{3, 2, 5, 4, 6, 1},
			},
		},
		{
			"7x7 full puzzle",
			SolvePuzzle7x7,
			[]int{7, 0, 0, 0, 2, 2, 3, 0, 0, 3, 0, 0, 0, 0, 3, 0, 3, 0, 0, 5, 0, 0, 0, 0, 0, 5, 0, 4},
			nil,
			[][]int{
				{1, 5, 6, 7, 4, 3, 2},
				{2, 7, 4, 5, 3, 1, 6},
				{3, 4, 5, 6, 7, 2, 1},
				{4, 6, 3, 1, 2, 7, 5},
				{5, 3, 1, 2, 6, 4, 7},
				{6, 2, 7, 3, 1, 5, 4},
				{7, 1, 2, 4, 5, 6, 3},
			},
		},
	}

	for _, c := range cases {
		actual := c.solve(c.clues, c.starting)
		for rIndex, row := range c.expected {
			for cIndex := range row {
				if len(actual) < rIndex+1 || len(actual[rIndex]) < cIndex+1 {
					t.Errorf("No value at %d, %d", rIndex, cIndex)
				} else if actual[rIndex][cIndex] != c.expected[rIndex][cIndex] {
					t.Errorf("Expected %d at [%d,%d], got %d", c.expected[rIndex][cIndex], rIndex, cIndex, actual[rIndex][cIndex])
				}
			}
		}
	}
}

func TestNewBoardFrom(t *testing.T) {

	startingState := [][][]int{
		{{1, 2}, {1, 2}},
		{{1}, {2}},
	}

	b := NewBoardFrom(2, startingState)
	AssertBoardState(t, b, startingState)
}

func TestBoard_Includes(t *testing.T) {

	startingState := [][][]int{
		{{2}, {1, 2}},
		{{1}, {2}},
	}

	b := NewBoardFrom(2, startingState)

	AssertIncludes(t, b, 0, 0, []int{2})
	AssertIncludes(t, b, 0, 1, []int{1, 2})
	AssertIncludes(t, b, 1, 0, []int{1})
	AssertIncludes(t, b, 1, 1, []int{2})
}

func TestBoard_Set(t *testing.T) {

	startingState := [][][]int{
		{{1, 2}, {1, 2}},
		{{1, 2}, {1, 2}},
	}

	b := NewBoardFrom(2, startingState)
	err := b.Set(cell(0, 0), 1)
	if err != nil {
		t.Fatalf("unable to set %v", err)
	}

	AssertBoardState(t, b, [][][]int{
		{{1}, {2}},
		{{2}, {1}},
	})
}

func TestBoard_cellIndex(t *testing.T) {

	cases := map[Cell]int{
		Cell{0, 0}: 0,
		Cell{0, 1}: 2,
		Cell{1, 0}: 4,
		Cell{1, 1}: 6,
	}

	startingState := [][][]int{
		{{1, 2}, {1, 2}},
		{{1, 2}, {1, 2}},
	}

	b := NewBoardFrom(2, startingState)

	for c, expected := range cases {
		actual := b.cellIndex(&c)
		if actual != expected {
			t.Errorf("Expected %d for %v, got %d", expected, c, actual)
		}

	}
}

func AssertIncludes(t *testing.T, b *Board, r, c int, v []int) {
	for _, expected := range v {
		cell := cell(r, c)
		if !b.Includes(cell, expected) {
			t.Errorf("%v should contain %v", cell, v)
		}
	}
}

func AssertBoardState(t *testing.T, b *Board, expected [][][]int) {
	hasFailed := false
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			cell := cell(r, c)
			for _, cellValue := range expected[r][c] {
				if !b.Includes(cell, cellValue) {
					hasFailed = true
					t.Errorf("%v does not include %v", cell, cellValue)
				}
			}
		}
	}

	if hasFailed {
		t.Errorf("Board state: %v", b)
	}
}
