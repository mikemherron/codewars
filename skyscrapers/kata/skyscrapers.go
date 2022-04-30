package kata

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
)

func SolvePuzzle(clues []int) [][]int {
	//makeClueSolutions(7)

	return SolvePuzzle7x7(clues, nil)
	//return SolvePuzzle4x4(clues, nil)
	//return SolvePuzzle4x4(clues, nil)
}

func SolvePuzzle4x4(clues []int, starting [][][]int) [][]int {
	return solve(
		NewBoardFrom(4, starting),
		[]*CellRange{
			{0, 0, Down},
			{0, 1, Down},
			{0, 2, Down},
			{0, 3, Down},

			{0, 3, Left},
			{1, 3, Left},
			{2, 3, Left},
			{3, 3, Left},

			{3, 3, Up},
			{3, 2, Up},
			{3, 1, Up},
			{3, 0, Up},

			{3, 0, Right},
			{2, 0, Right},
			{1, 0, Right},
			{0, 0, Right},
		},
		func(clue int, cells [][]int, b *Board) {
			if clue == 1 {
				_ = b.Set(cells[0][0], cells[0][1], 4)
			} else if clue == 2 {
				_ = b.Remove(cells[0][0], cells[0][1], 4)
			} else if clue == 3 {
				_ = b.Remove(cells[0][0], cells[0][1], 4)
				_ = b.Remove(cells[0][0], cells[0][1], 3)
				_ = b.Remove(cells[1][0], cells[1][0], 4)
			} else if clue == 4 {
				_ = b.Set(cells[0][0], cells[0][1], 1)
				_ = b.Set(cells[1][0], cells[1][1], 2)
				_ = b.Set(cells[2][0], cells[2][1], 3)
				_ = b.Set(cells[3][0], cells[3][1], 4)
			}
		},
		clues,
	)
}

//.func SolvePuzzle6x6(clues []int, starting [][][]int) [][]int {
//return solve(
//	NewBoardFrom(6, starting),
//	[]*CellRange{
//		{cell(0, 0), Down},
//		{cell(0, 1), Down},
//		{cell(0, 2), Down},
//		{cell(0, 3), Down},
//		{cell(0, 4), Down},
//		{cell(0, 5), Down},
//
//		{cell(0, 5), Left},
//		{cell(1, 5), Left},
//		{cell(2, 5), Left},
//		{cell(3, 5), Left},
//		{cell(4, 5), Left},
//		{cell(5, 5), Left},
//
//		{cell(5, 5), Up},
//		{cell(5, 4), Up},
//		{cell(5, 3), Up},
//		{cell(5, 2), Up},
//		{cell(5, 1), Up},
//		{cell(5, 0), Up},
//
//		{cell(5, 0), Right},
//		{cell(4, 0), Right},
//		{cell(3, 0), Right},
//		{cell(2, 0), Right},
//		{cell(1, 0), Right},
//		{cell(0, 0), Right},
//	},
//	func(clue int, cells []*Cell, b *Board) {
//		if clue == 1 {
//			_ = b.Set(cells[0], 6)
//		} else if clue == 2 {
//			_ = b.Remove(cells[0], 6)
//		} else if clue == 3 {
//			_ = b.Remove(cells[0], 6)
//			_ = b.Remove(cells[1], 6)
//
//			_ = b.Remove(cells[0], 5)
//		} else if clue == 4 {
//			_ = b.Remove(cells[0], 6)
//			_ = b.Remove(cells[1], 6)
//			_ = b.Remove(cells[2], 6)
//
//			_ = b.Remove(cells[0], 5)
//			_ = b.Remove(cells[1], 5)
//
//			_ = b.Remove(cells[0], 4)
//		} else if clue == 5 {
//			_ = b.Remove(cells[0], 6)
//			_ = b.Remove(cells[1], 6)
//			_ = b.Remove(cells[2], 6)
//			_ = b.Remove(cells[3], 6)
//
//			_ = b.Remove(cells[0], 5)
//			_ = b.Remove(cells[1], 5)
//			_ = b.Remove(cells[2], 5)
//
//			_ = b.Remove(cells[0], 4)
//			_ = b.Remove(cells[1], 4)
//
//			_ = b.Remove(cells[0], 3)
//
//		} else if clue == 6 {
//			_ = b.Set(cells[0], 1)
//			_ = b.Set(cells[1], 2)
//			_ = b.Set(cells[2], 3)
//			_ = b.Set(cells[3], 4)
//			_ = b.Set(cells[4], 5)
//			_ = b.Set(cells[5], 6)
//		}
//	},
//	clues,
//)
//}

func SolvePuzzle7x7(clues []int, starting [][][]int) [][]int {
	return solve(
		NewBoardFrom(7, starting),
		[]*CellRange{
			{0, 0, Down},
			{0, 1, Down},
			{0, 2, Down},
			{0, 3, Down},
			{0, 4, Down},
			{0, 5, Down},
			{0, 6, Down},

			{0, 6, Left},
			{1, 6, Left},
			{2, 6, Left},
			{3, 6, Left},
			{4, 6, Left},
			{5, 6, Left},
			{6, 6, Left},

			{6, 6, Up},
			{6, 5, Up},
			{6, 4, Up},
			{6, 3, Up},
			{6, 2, Up},
			{6, 1, Up},
			{6, 0, Up},

			{6, 0, Right},
			{5, 0, Right},
			{4, 0, Right},
			{3, 0, Right},
			{2, 0, Right},
			{1, 0, Right},
			{0, 0, Right},
		},
		func(clue int, cells [][]int, b *Board) {
			if clue == 1 {
				_ = b.Set(cells[0][0], cells[0][1], 7)
			} else if clue == 2 {
				_ = b.Remove(cells[0][0], cells[0][1], 7)
			} else if clue == 3 {
				_ = b.Remove(cells[0][0], cells[0][1], 7)
				_ = b.Remove(cells[1][0], cells[1][1], 7)

				_ = b.Remove(cells[0][0], cells[0][1], 6)
			} else if clue == 4 {
				_ = b.Remove(cells[0][0], cells[0][1], 7)
				_ = b.Remove(cells[1][0], cells[1][1], 7)
				_ = b.Remove(cells[2][0], cells[2][1], 7)

				_ = b.Remove(cells[0][0], cells[0][1], 6)
				_ = b.Remove(cells[1][0], cells[1][1], 6)

				_ = b.Remove(cells[0][0], cells[0][1], 5)
			} else if clue == 5 {
				_ = b.Remove(cells[0][0], cells[0][1], 7)
				_ = b.Remove(cells[1][0], cells[1][1], 7)
				_ = b.Remove(cells[2][0], cells[2][1], 7)
				_ = b.Remove(cells[3][0], cells[3][1], 7)

				_ = b.Remove(cells[0][0], cells[0][1], 6)
				_ = b.Remove(cells[1][0], cells[1][1], 6)
				_ = b.Remove(cells[2][0], cells[2][1], 6)

				_ = b.Remove(cells[0][0], cells[0][1], 5)
				_ = b.Remove(cells[1][0], cells[1][1], 5)

				_ = b.Remove(cells[0][0], cells[0][1], 4)

			} else if clue == 6 {
				_ = b.Remove(cells[0][0], cells[0][1], 7)
				_ = b.Remove(cells[1][0], cells[1][1], 7)
				_ = b.Remove(cells[2][0], cells[2][1], 7)
				_ = b.Remove(cells[3][0], cells[3][1], 7)
				_ = b.Remove(cells[4][0], cells[4][1], 7)

				_ = b.Remove(cells[0][0], cells[0][1], 6)
				_ = b.Remove(cells[1][0], cells[1][1], 6)
				_ = b.Remove(cells[2][0], cells[2][1], 6)
				_ = b.Remove(cells[3][0], cells[3][1], 6)

				_ = b.Remove(cells[0][0], cells[0][1], 5)
				_ = b.Remove(cells[1][0], cells[1][1], 5)
				_ = b.Remove(cells[2][0], cells[2][1], 5)

				_ = b.Remove(cells[0][0], cells[0][1], 4)
				_ = b.Remove(cells[1][0], cells[1][1], 4)

				_ = b.Remove(cells[0][0], cells[0][1], 3)

			} else if clue == 7 {
				_ = b.Set(cells[0][0], cells[0][1], 1)
				_ = b.Set(cells[1][0], cells[1][1], 2)
				_ = b.Set(cells[2][0], cells[2][1], 3)
				_ = b.Set(cells[3][0], cells[3][1], 4)
				_ = b.Set(cells[4][0], cells[4][1], 5)
				_ = b.Set(cells[5][0], cells[5][1], 6)
				_ = b.Set(cells[6][0], cells[6][1], 7)
			}
		},
		clues,
	)
}

type Deducer func(clue int, cells [][]int, b *Board)

var profile = flag.String("profile", "", "")

func solve(b *Board, clueMap []*CellRange, deduce Deducer, clueValues []int) [][]int {
	flag.Parse()
	if *profile != "" {
		f, err := os.Create(*profile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	clues := make([]*Clue, 0)
	for i, clueValue := range clueValues {
		if clueValue == 0 {
			continue
		}

		cells := b.GetCells(clueMap[i])
		clues = append(clues, &Clue{clueValue, cells, make([][]int, 0)})

		//Apply the initial deductions to the board state
		deduce(clueValue, cells, b)
	}

	//full clue approach:

	//makeClueSolutions(b.size)
	//populateSolutions(clues, b)
	//sol, err := checkSolutions(clues, newClueTester(b.size))
	//if err != nil {
	//	fmt.Printf("Chgeckjed: %d\n", checked)
	//	panic(err)
	//}
	//	return sol

	//	Inital recusive

	sort.Sort(ByVisible(clues))
	sol := findSolutionStack(b, clues, 0, 0, 0)
	if sol == nil {
		panic("unable to find solution")
	}

	return sol.Collapse()
}

func findSolutionStack(b *Board, clues []*Clue, visible int, maxHeight int, cellStart int) *Board {

	nextClue := clues[0]

	for i := cellStart; i < len(nextClue.cells); i++ {
		c := nextClue.cells[i]
		possible := b.Possibles(c[0], c[1])
		if len(possible) > 1 {
			for _, v := range possible {
				if v > maxHeight && visible+1 > nextClue.visible {
					continue
				}

				boardCopy := b.Copy()
				err := boardCopy.Set(c[0], c[1], v)
				if err != nil {
					continue
				}

				solution := findSolutionStack(boardCopy, clues, visible, maxHeight, i)
				if solution != nil {
					//solution.Finalise()
					return solution
				}
			}

			return nil
		}

		height := possible[0]
		if height > maxHeight {
			visible++
			maxHeight = height
		}

	}

	if nextClue.visible == visible {
		//this nextClue has been satisfied
		if len(clues) == 1 {
			//this was the last nextClue, this is our solution
			return b
		}
		return findSolutionStack(b, clues[1:], 0, 0, 0)
	}

	//The final solution for this clue does not match the value the clue
	//expects, so return nil
	return nil
}

type Clue struct {
	visible   int
	cells     [][]int
	solutions [][]int
}

type ByVisible []*Clue

func (c ByVisible) Len() int           { return len(c) }
func (c ByVisible) Less(i, j int) bool { return c[i].visible > c[j].visible }
func (c ByVisible) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

type ClueTester struct {
	rowVals       []int
	colVals       []int
	cellVals      [][]int
	cellValsCount [][]int
	size          int
}

func newClueTester(size int) *ClueTester {
	t := ClueTester{
		rowVals:       make([]int, size),
		colVals:       make([]int, size),
		cellVals:      make([][]int, size),
		cellValsCount: make([][]int, size),
		size:          size,
	}

	for r := 0; r < size; r++ {
		t.cellVals[r] = make([]int, size)
		t.cellValsCount[r] = make([]int, size)
	}

	return &t
}

func (ct *ClueTester) Copy() *ClueTester {
	ctCopy := newClueTester(ct.size)
	for r := 0; r < ct.size; r++ {
		copy(ctCopy.cellVals[r], ct.cellVals[r])
	}

	copy(ctCopy.rowVals, ct.rowVals)
	copy(ctCopy.colVals, ct.colVals)

	return ctCopy
}

var errorInvalid = errors.New("cant apply")

//
//func (ct *ClueTester) apply(cells [][]int, values []int) error {
//	for i, cell := range cells {
//		v, r, c := values[i], cell[0], cell[1]
//		if hasBit(ct.rowVals[r], v) {
//			fmt.Printf("Can't set [%d,%d] to %v, row already has that value\n", r, c, v)
//			return errorInvalid
//		}
//		ct.rowVals[r] |= 1 << v
//
//		if hasBit(ct.colVals[c], v) {
//			fmt.Printf("Can't set [%d,%d] to %v, col already has that value\n", r, c, v)
//			return errorInvalid
//		}
//		ct.colVals[c] |= 1 << v
//
//		if ct.cellVals[r][c] != 0 && ct.cellVals[r][c] != v {
//			fmt.Printf("Can't set [%d,%d] to %v, cell already has a different value (%d) \n", r, c, v, ct.cellVals[r][c])
//			return errorInvalid
//		}
//
//		ct.cellVals[r][c] = v
//		ct.cellValsCount[r][c]++
//	}
//
//	return nil
//}
//
//func (ct *ClueTester) undo(cells [][]int, values []int) {
//	for i, cell := range cells {
//		v, r, c := values[i], cell[0], cell[1]
//
//		mask := ^(1 << v)
//		ct.rowVals[r] &= mask
//		ct.colVals[c] &= mask
//
//		ct.cellValsCount[r][c]--
//		if ct.cellValsCount[r][c] == 0 {
//			ct.cellVals[r][c] = 0
//		}
//	}
//}
//
//func (ct *ClueTester) test(cells [][]int, values []int) bool {
//	for i, cell := range cells {
//		v, r, c := values[i], cell[0], cell[1]
//		if hasBit(ct.rowVals[r], v) {
//			return false
//		}
//
//		if hasBit(ct.colVals[c], v) {
//			return false
//		}
//
//		if ct.cellVals[r][c] != 0 && ct.cellVals[r][c] != v {
//			return false
//		}
//	}
//
//	return true
//}
//
//func hasBit(n int, pos int) bool {
//	val := n & (1 << pos)
//	return val > 0
//}

//Next steps
// -for each clue, go through clueSolutions and find candidates
// - once we have candidates, we need to rule them out
// - rather than doing this on a full board, make a new struct
/*

 rows[idx] = a bit map with values  set on each row
 cols[oidx] same as obve
 cells[odx]pdox] - valie for each cell

 apply method
  - can chec kfor cella nd ools
  DONT recurse

On

*/
//
//var checked int
//
//func checkSolutions(clues []*Clue, ct *ClueTester) ([][]int, error) {
//	checked++
//	if len(clues) == 0 {
//		b := NewBoard(ct.size)
//		for r := 0; r < b.size; r++ {
//			for c := 0; c < b.size; c++ {
//				e := b.Set(r, c, ct.cellVals[r][c])
//				if e != nil {
//					panic("That shouldn't jave ha[[pend!")
//				}
//			}
//		}
//		return b.Collapse(), nil
//	}
//
//	nextClue := clues[0]
//	for _, solution := range nextClue.solutions {
//
//		err := ct.apply(nextClue.cells, solution)
//		if err != nil {
//			ct.undo(nextClue.cells, solution)
//			continue
//		}
//
//		sol, err := checkSolutions(clues[1:], ct)
//		if err != nil {
//			ct.undo(nextClue.cells, solution)
//			continue
//		}
//
//		return sol, nil
//	}
//
//	return nil, errorInvalid
//}
//
//var clueSolutions map[int][][]int
//
//func populateSolutions(clues []*Clue, b *Board) {
//	for _, clue := range clues {
//		for _, solution := range clueSolutions[clue.visible] {
//			match := true
//			for cellIndex, solutionValue := range solution {
//				if !b.Includes(clue.cells[cellIndex][0], clue.cells[cellIndex][1], solutionValue) {
//					match = false
//					break
//				}
//			}
//			if match {
//				clue.solutions = append(clue.solutions, solution)
//			}
//		}
//	}
//}
//
//func remove(s map[int]bool, v int) map[int]bool {
//	c := make(map[int]bool)
//	for k := range s {
//		if k == v {
//			continue
//		}
//		c[k] = true
//	}
//
//	return c
//}
//
//func copyWith(s []int, v int) []int {
//	c := make([]int, len(s))
//	copy(c, s)
//	c = append(c, v)
//	return c
//}
//
//func makeClueSolutions(size int) {
//	clueSolutions = make(map[int][][]int)
//
//	initSet := make(map[int]bool)
//	for i := 1; i <= size; i++ {
//		initSet[i] = true
//	}
//
//	makeClueSolution(initSet, make([]int, 0))
//}
//
//func makeClueSolution(remaining map[int]bool, current []int) {
//	if len(remaining) == 0 {
//		max, visible := 0, 0
//		for _, h := range current {
//			if h > max {
//				visible++
//				max = h
//			}
//		}
//		clueSolutions[visible] = append(clueSolutions[visible], current)
//		return
//	}
//
//	for r := range remaining {
//		makeClueSolution(remove(remaining, r), copyWith(current, r))
//	}
//}

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

type CellRange struct {
	StartRow int
	StartCol int
	Dir      Direction
}

var errUnsolvableCell = errors.New("cell has no remaining values")

type Board struct {
	bits    []bool
	size    int
	rowVals []int
	colVals []int
}

func (b Board) CellIndex(r int, c int) int {
	return (r * b.size * b.size) + (c * b.size)
}

func (b Board) valueIndex(r int, c int, v int) int {
	return b.CellIndex(r, c) + (v - 1)
}

func (b Board) Includes(r int, c int, v int) bool {
	return b.bits[b.valueIndex(r, c, v)]
}

func (b Board) Only(r int, c int) int {
	possibles, only := b.NumPossibles(r, c)
	if possibles != 1 {
		panic(fmt.Sprintf("invalid only call %d", c))
	}

	return only
}

func (b Board) NumPossibles(r int, c int) (int, int) {
	n := 0
	p := 0
	cellIndex := b.CellIndex(r, c)
	for idx, isSet := range b.bits[cellIndex : cellIndex+b.size] {
		if isSet {
			n++
			if n == 1 {
				p = idx + 1
			}
		}
	}

	return n, p
}

func (b Board) Possibles(r int, c int) []int {
	p := make([]int, 0, b.size)
	for v := 1; v <= b.size; v++ {
		if b.Includes(r, c, v) {
			p = append(p, v)
		}
	}

	return p
}

func (b Board) PossiblesCb(r int, c int, cb func(v int)) {
	for v := 1; v <= b.size; v++ {
		if b.Includes(r, c, v) {
			cb(v)
		}
	}
}

func (b Board) EachCell(cb func(r int, c int)) {
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			cb(r, c)
		}
	}
}

func (b Board) GetCells(a *CellRange) [][]int {
	cells := make([][]int, b.size)
	for i := 0; i < b.size; i++ {
		switch a.Dir {
		case Up:
			cells[i] = []int{a.StartRow - i, a.StartCol}
		case Down:
			cells[i] = []int{a.StartRow + i, a.StartCol}
		case Left:
			cells[i] = []int{a.StartRow, a.StartCol - i}
		case Right:
			cells[i] = []int{a.StartRow, a.StartCol + i}
		}
	}

	return cells
}

func (b Board) Remove(r int, c int, v int) error {
	b.bits[b.valueIndex(r, c, v)] = false

	possibles, firstPos := b.NumPossibles(r, c)
	if possibles == 0 {
		return errUnsolvableCell
	}

	if possibles == 1 {
		err := b.Set(r, c, firstPos)
		if err != nil {
			return errUnsolvableCell
		}
	}

	return nil
}

func (b Board) Init(r int, c int, v int) {
	b.bits[b.valueIndex(r, c, v)] = true
}

//
//func (b Board) SetQuick(r int, c int, v int) error {
//
//	if hasBit(b.rowVals[r], v) {
//		return errorInvalid
//	}
//
//	if hasBit(b.colVals[c], v) {
//		return errorInvalid
//	}
//
//	num, val := b.NumPossibles(r, c)
//	if num == 1 && val != v {
//		return errorInvalid
//	}
//
//	b.rowVals[r] |= 1 << v
//	b.colVals[c] |= 1 << v
//
//	cellIndex := b.CellIndex(r, c)
//	for s := 0; s < b.size; s++ {
//		b.bits[cellIndex+s] = s == v-1
//	}
//
//	return nil
//}
//
//func (b Board) UndoSetQuick(r int, c int, v int) error {
//
//	if hasBit(b.rowVals[r], v) {
//		return errorInvalid
//	}
//
//	if hasBit(b.colVals[c], v) {
//		return errorInvalid
//	}
//
//	num, val := b.NumPossibles(r, c)
//	if num == 1 && val != v {
//		return errorInvalid
//	}
//
//	b.rowVals[r] |= 1 << v
//	b.colVals[c] |= 1 << v
//
//	cellIndex := b.CellIndex(r, c)
//	for s := 0; s < b.size; s++ {
//		b.bits[cellIndex+s] = s == v-1
//	}
//
//	return nil
//}
//
//func (b Board) Finalise() {
//	for r := 0; r < b.size; r++ {
//		for c := 0; c < b.size; c++ {
//			num, val := b.NumPossibles(r, c)
//			if num == 1 {
//				err := b.Set(r, c, val)
//				if err != nil {
//					panic("Erro!")
//				}
//			}
//		}
//	}
//}

func (b Board) Set(r int, c int, v int) error {

	cellIndex := b.CellIndex(r, c)
	for s := 0; s < b.size; s++ {
		b.bits[cellIndex+s] = s == v-1
	}

	for oR := 0; oR < b.size; oR++ {
		if oR != r && b.Includes(oR, c, v) {
			err := b.Remove(oR, c, v)
			if err != nil {
				return errUnsolvableCell
			}
		}
	}

	for oC := 0; oC < b.size; oC++ {
		if oC != c && b.Includes(r, oC, v) {
			err := b.Remove(r, oC, v)
			if err != nil {
				return errUnsolvableCell
			}
		}
	}

	return nil
}

func (b Board) Copy() *Board {
	bitsCopy := append(b.bits[:0:0], b.bits...)
	//rowsCopy := append(b.rowVals[:0:0], b.rowVals...)
	//colsCopy := append(b.colVals[:0:0], b.colVals...)
	bCopy := Board{bits: bitsCopy, size: b.size}
	return &bCopy
}

func (b Board) Collapse() [][]int {
	result := make([][]int, b.size)
	for r := 0; r < b.size; r++ {
		result[r] = make([]int, b.size)
		for c := 0; c < b.size; c++ {
			_, v := b.NumPossibles(r, c)
			result[r][c] = v
		}
	}

	return result
}

func NewBoard(size int) Board {
	b := Board{bits: make([]bool, size*size*size), size: size, rowVals: make([]int, size), colVals: make([]int, size)}
	for i := range b.bits {
		b.bits[i] = true
	}

	return b
}

func NewBoardFrom(size int, initial [][][]int) *Board {
	b := NewBoard(size)
	if initial == nil {
		return &b
	}

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			for _, s := range initial[r][c] {
				b.Init(r, c, s)
			}
		}
	}

	return &b
}

func (b Board) String() string {
	out := "\n"
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			s := "|"
			for v := 0; v < b.size; v++ {
				if b.Includes(r, c, v+1) {
					s += fmt.Sprintf("%d", v+1)
				} else {
					s += "-"
				}
			}
			s += "|"
			out += s
		}
		out += "\n"
	}
	return out
}
