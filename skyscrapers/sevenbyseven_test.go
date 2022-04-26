package skyscrapers

import "testing"

func TestSolvePuzzleSevenBySeven(t *testing.T) {

	type testCase struct {
		name     string
		clues    []int
		expected [][]int
		starting [][][]int
	}

	cases := []testCase{
		{
			"full puzzle",
			[]int{7, 0, 0, 0, 2, 2, 3, 0, 0, 3, 0, 0, 0, 0, 3, 0, 3, 0, 0, 5, 0, 0, 0, 0, 0, 5, 0, 4},
			[][]int{
				{1, 5, 6, 7, 4, 3, 2},
				{2, 7, 4, 5, 3, 1, 6},
				{3, 4, 5, 6, 7, 2, 1},
				{4, 6, 3, 1, 2, 7, 5},
				{5, 3, 1, 2, 6, 4, 7},
				{6, 2, 7, 3, 1, 5, 4},
				{7, 1, 2, 4, 5, 6, 3},
			},
			nil,
		},
		//{
		//	"almost completed puzzle",
		//	[]int{
		//		0, 0, 1, 2,
		//		0, 2, 0, 0,
		//		0, 3, 0, 0,
		//		0, 1, 0, 0},
		//	[][]int{
		//		{2, 1, 4, 3},
		//		{3, 4, 1, 2},
		//		{4, 2, 3, 1},
		//		{1, 3, 2, 4},
		//	},
		//	[][][]int{
		//		{{2, 3}, {1}, {4}, {2, 3}},
		//		{{1, 3}, {4}, {1}, {2}},
		//		{{4}, {2}, {3}, {1}},
		//		{{1, 3}, {3, 1}, {2}, {4}},
		//	},
		//},
	}

	for _, c := range cases {
		actual := SolvePuzzleSevenBySevenFrom(c.clues, c.starting)
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
