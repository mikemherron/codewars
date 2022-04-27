package skyscrapers

import "testing"

func TestNewBoardFrom(t *testing.T) {

	startingState := [][][]int{
		{{1, 2}, {1, 2}},
		{{1}, {2}},
	}

	b := NewBoardFrom(2, startingState)
	AssertBoardState(t, b, startingState)
}

func TestBoard_Bitset(t *testing.T) {
	bs := NewBitSet()
	Assert(t, bs.has(23), "should be set on by defult")
	Assert(t, bs.has(242), "should be set on by defult")
	bs.clear(23)
	Assert(t, !bs.has(23), "should be off")
	bs.set(23)
	Assert(t, bs.has(23), "should be on")
	bs.clear(10)
	Assert(t, !bs.has(10), "should be off")
	Assert(t, bs.has(10+64), "should be on")
	Assert(t, bs.has(10+64+64), "should be on")
	Assert(t, bs.has(10+64+64+64), "should be on")
}

func Assert(t *testing.T, v bool, m string) {
	if !v {
		t.Errorf(m)
	}
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
	AssertIncludes(t, b, 1, 2, []int{2})
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

func AssertIncludes(t *testing.T, b Board, r, c int, v []int) {
	for _, expected := range v {
		cell := cell(r, c)
		if !b.Includes(cell, expected) {
			t.Errorf("%v should contain %v", cell, v)
		}
	}
}

func AssertBoardState(t *testing.T, b Board, expected [][][]int) {
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
