package skyscrapers

import "testing"

func TestBoard(t *testing.T) {

}
func TestSolvePuzzleSixBySix(t *testing.T) {

	type testCase struct {
		name     string
		clues    []int
		expected [][]int
		starting [][][]int
	}

	cases := []testCase{
		{
			"full puzzle",
			[]int{3, 2, 2, 3, 2, 1,
				1, 2, 3, 3, 2, 2,
				5, 1, 2, 2, 4, 3,
				3, 2, 1, 2, 2, 4},
			[][]int{{2, 1, 4, 3, 5, 6},
				{1, 6, 3, 2, 4, 5},
				{4, 3, 6, 5, 1, 2},
				{6, 5, 2, 1, 3, 4},
				{5, 4, 1, 6, 2, 3},
				{3, 2, 5, 4, 6, 1},
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
		actual := SolvePuzzleSixBySixFrom(c.clues, c.starting)
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
