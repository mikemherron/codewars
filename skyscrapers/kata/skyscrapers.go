package kata

import (
	"errors"
	"fmt"
)

func SolvePuzzle(clues []int) [][]int {
	return SolvePuzzle7x7(clues, nil)
	//return SolvePuzzle4x4(clues, nil)
	//return SolvePuzzle4x4(clues, nil)
}

func SolvePuzzle4x4(clues []int, starting [][][]int) [][]int {
	return solve(
		NewBoardFrom(4, starting),
		[]*CellRange{
			{cell(0, 0), Down},
			{cell(0, 1), Down},
			{cell(0, 2), Down},
			{cell(0, 3), Down},

			{cell(0, 3), Left},
			{cell(1, 3), Left},
			{cell(2, 3), Left},
			{cell(3, 3), Left},

			{cell(3, 3), Up},
			{cell(3, 2), Up},
			{cell(3, 1), Up},
			{cell(3, 0), Up},

			{cell(3, 0), Right},
			{cell(2, 0), Right},
			{cell(1, 0), Right},
			{cell(0, 0), Right},
		},
		func(clue int, cells []*Cell, b *Board) {
			if clue == 1 {
				_ = b.Set(cells[0], 4)
			} else if clue == 2 {
				_ = b.Remove(cells[0], 4)
			} else if clue == 3 {
				_ = b.Remove(cells[0], 4)
				_ = b.Remove(cells[0], 3)
				_ = b.Remove(cells[1], 4)
			} else if clue == 4 {
				_ = b.Set(cells[0], 1)
				_ = b.Set(cells[1], 2)
				_ = b.Set(cells[2], 3)
				_ = b.Set(cells[3], 4)
			}
		},
		clues,
	)
}

func SolvePuzzle6x6(clues []int, starting [][][]int) [][]int {
	return solve(
		NewBoardFrom(6, starting),
		[]*CellRange{
			{cell(0, 0), Down},
			{cell(0, 1), Down},
			{cell(0, 2), Down},
			{cell(0, 3), Down},
			{cell(0, 4), Down},
			{cell(0, 5), Down},

			{cell(0, 5), Left},
			{cell(1, 5), Left},
			{cell(2, 5), Left},
			{cell(3, 5), Left},
			{cell(4, 5), Left},
			{cell(5, 5), Left},

			{cell(5, 5), Up},
			{cell(5, 4), Up},
			{cell(5, 3), Up},
			{cell(5, 2), Up},
			{cell(5, 1), Up},
			{cell(5, 0), Up},

			{cell(5, 0), Right},
			{cell(4, 0), Right},
			{cell(3, 0), Right},
			{cell(2, 0), Right},
			{cell(1, 0), Right},
			{cell(0, 0), Right},
		},
		func(clue int, cells []*Cell, b *Board) {
			if clue == 1 {
				_ = b.Set(cells[0], 6)
			} else if clue == 2 {
				_ = b.Remove(cells[0], 6)
			} else if clue == 3 {
				_ = b.Remove(cells[0], 6)
				_ = b.Remove(cells[1], 6)

				_ = b.Remove(cells[0], 5)
			} else if clue == 4 {
				_ = b.Remove(cells[0], 6)
				_ = b.Remove(cells[1], 6)
				_ = b.Remove(cells[2], 6)

				_ = b.Remove(cells[0], 5)
				_ = b.Remove(cells[1], 5)

				_ = b.Remove(cells[0], 4)
			} else if clue == 5 {
				_ = b.Remove(cells[0], 6)
				_ = b.Remove(cells[1], 6)
				_ = b.Remove(cells[2], 6)
				_ = b.Remove(cells[3], 6)

				_ = b.Remove(cells[0], 5)
				_ = b.Remove(cells[1], 5)
				_ = b.Remove(cells[2], 5)

				_ = b.Remove(cells[0], 4)
				_ = b.Remove(cells[1], 4)

				_ = b.Remove(cells[0], 3)

			} else if clue == 6 {
				_ = b.Set(cells[0], 1)
				_ = b.Set(cells[1], 2)
				_ = b.Set(cells[2], 3)
				_ = b.Set(cells[3], 4)
				_ = b.Set(cells[4], 5)
				_ = b.Set(cells[5], 6)
			}
		},
		clues,
	)
}

func SolvePuzzle7x7(clues []int, starting [][][]int) [][]int {
	return solve(
		NewBoardFrom(7, starting),
		[]*CellRange{
			{cell(0, 0), Down},
			{cell(0, 1), Down},
			{cell(0, 2), Down},
			{cell(0, 3), Down},
			{cell(0, 4), Down},
			{cell(0, 5), Down},
			{cell(0, 6), Down},

			{cell(0, 6), Left},
			{cell(1, 6), Left},
			{cell(2, 6), Left},
			{cell(3, 6), Left},
			{cell(4, 6), Left},
			{cell(5, 6), Left},
			{cell(6, 6), Left},

			{cell(6, 6), Up},
			{cell(6, 5), Up},
			{cell(6, 4), Up},
			{cell(6, 3), Up},
			{cell(6, 2), Up},
			{cell(6, 1), Up},
			{cell(6, 0), Up},

			{cell(6, 0), Right},
			{cell(5, 0), Right},
			{cell(4, 0), Right},
			{cell(3, 0), Right},
			{cell(2, 0), Right},
			{cell(1, 0), Right},
			{cell(0, 0), Right},
		},
		func(clue int, cells []*Cell, b *Board) {
			if clue == 1 {
				_ = b.Set(cells[0], 7)
			} else if clue == 2 {
				_ = b.Remove(cells[0], 7)
			} else if clue == 3 {
				_ = b.Remove(cells[0], 7)
				_ = b.Remove(cells[1], 7)

				_ = b.Remove(cells[0], 6)
			} else if clue == 4 {
				_ = b.Remove(cells[0], 7)
				_ = b.Remove(cells[1], 7)
				_ = b.Remove(cells[2], 7)

				_ = b.Remove(cells[0], 6)
				_ = b.Remove(cells[1], 6)

				_ = b.Remove(cells[0], 5)
			} else if clue == 5 {
				_ = b.Remove(cells[0], 7)
				_ = b.Remove(cells[1], 7)
				_ = b.Remove(cells[2], 7)
				_ = b.Remove(cells[3], 7)

				_ = b.Remove(cells[0], 6)
				_ = b.Remove(cells[1], 6)
				_ = b.Remove(cells[2], 6)

				_ = b.Remove(cells[0], 5)
				_ = b.Remove(cells[1], 5)

				_ = b.Remove(cells[0], 4)

			} else if clue == 6 {
				_ = b.Remove(cells[0], 7)
				_ = b.Remove(cells[1], 7)
				_ = b.Remove(cells[2], 7)
				_ = b.Remove(cells[3], 7)
				_ = b.Remove(cells[4], 7)

				_ = b.Remove(cells[0], 6)
				_ = b.Remove(cells[1], 6)
				_ = b.Remove(cells[2], 6)
				_ = b.Remove(cells[3], 6)

				_ = b.Remove(cells[0], 5)
				_ = b.Remove(cells[1], 5)
				_ = b.Remove(cells[2], 5)

				_ = b.Remove(cells[0], 4)
				_ = b.Remove(cells[1], 4)

				_ = b.Remove(cells[0], 3)

			} else if clue == 7 {
				_ = b.Set(cells[0], 1)
				_ = b.Set(cells[1], 2)
				_ = b.Set(cells[2], 3)
				_ = b.Set(cells[3], 4)
				_ = b.Set(cells[4], 5)
				_ = b.Set(cells[5], 6)
				_ = b.Set(cells[6], 7)
			}
		},
		clues,
	)
}

type Deducer func(clue int, cells []*Cell, b *Board)

type Solver func(clueMap []*CellRange, clues []int, b *Board) *Board

func solve(b *Board, clueMap []*CellRange, deduce Deducer, clueValues []int) [][]int {
	clues := make([]*Clue, 0)
	for i, clueValue := range clueValues {
		if clueValue == 0 {
			continue
		}

		cells := b.GetCells(clueMap[i])
		deduce(clueValue, cells, b)
		clues = append(clues, &Clue{clueValue, cells, make([][]int, 0)})
	}

	sol := findSolutionStack(b, clues)
	if sol == nil {
		panic("unable to find solution")
	}

	return sol.Collapse()
}

func findSolutionStack(b *Board, clues []*Clue) *Board {

	visible := 0
	maxHeight := 0

	clue := clues[0]

	//fmt.Printf("Checking clue %v, %d left\n", clue, len(clues))

	for _, cell := range clue.cells {

		//fmt.Printf("-Checking cell %v\n", cell)
		numberPossible, height := b.NumPossibles(cell)
		if numberPossible > 1 {

			var s *Board
			b.PossiblesCb(cell, func(v int) {
				if s != nil {
					return
				}
				boardCopy := b.Copy()
				err := boardCopy.Set(cell, v)
				if err != nil {
					return
				}

				//TODO: Copy clues so don't need to traverse again
				//in recurions
				solution := findSolutionStack(&boardCopy, clues)
				if solution != nil {
					s = solution
				}
			})

			//We've filled in all possible values for this clue
			//returna solution if we have it or nil if not
			return s
		} else {
			//fmt.Printf("-cell resolved %v\n", height)
			if height > maxHeight {
				visible++
				maxHeight = height
			}
		}
	}

	//If we get here, all numbers have been filled in for this clue
	if clue.visible == visible {
		//this clue has been satisfied
		if len(clues) == 1 {
			//this was the last clue, this is our solution
			return b
		}
		return findSolutionStack(b, clues[1:])

	} else {
		//fmt.Printf("Clue falied got %d expected %d for %v", visible, clue.visible, b)

		//if we are here - the solution doesn't match the clue
		//return nil
		return nil
	}

}

type Cell struct {
	Row int
	Col int
}

func (c *Cell) String() string {
	return fmt.Sprintf("[%d, %d]", c.Row, c.Col)
}

var cellPool [][]*Cell

func cell(row, col int) *Cell {
	return cellPool[row][col]
}

func initCellPool(size int) {
	cellPool = make([][]*Cell, size)
	for r := 0; r < size; r++ {
		cellPool[r] = make([]*Cell, size)
		for c := 0; c < size; c++ {
			cellPool[r][c] = &Cell{r, c}
		}
	}
}

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

type CellRange struct {
	Start *Cell
	Dir   Direction
}

var errUnsolvableCell = errors.New("cell has no remaining values")

type Board struct {
	bits []bool
	size int
}

func (b Board) cellIndex(c *Cell) int {
	return (c.Row * b.size * b.size) + (c.Col * b.size)
}

func (b Board) valueIndex(c *Cell, v int) int {
	return b.cellIndex(c) + (v - 1)
}

func (b Board) Includes(c *Cell, v int) bool {
	return b.bits[b.valueIndex(c, v)]
}

func (b Board) Only(c *Cell) int {
	possibles, only := b.NumPossibles(c)
	if possibles != 1 {
		panic(fmt.Sprintf("invalid only call %d", c))
	}

	return only
}

func (b Board) NumPossibles(c *Cell) (int, int) {
	n := 0
	p := 0
	cellIndex := b.cellIndex(c)
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

func (b Board) Possibles(c *Cell) []int {
	p := make([]int, 0, b.size)
	for v := 1; v <= b.size; v++ {
		if b.Includes(c, v) {
			p = append(p, v)
		}
	}

	return p
}

func (b Board) PossiblesCb(c *Cell, cb func(v int)) {
	for v := 1; v <= b.size; v++ {
		if b.Includes(c, v) {
			cb(v)
		}
	}
}

func (b Board) EachCell(cb func(c *Cell)) {
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			cb(cell(r, c))
		}
	}
}

func (b Board) GetCells(a *CellRange) []*Cell {
	cells := make([]*Cell, b.size)
	for i := 0; i < b.size; i++ {
		switch a.Dir {
		case Up:
			cells[i] = cell(a.Start.Row-i, a.Start.Col)
		case Down:
			cells[i] = cell(a.Start.Row+i, a.Start.Col)
		case Left:
			cells[i] = cell(a.Start.Row, a.Start.Col-i)
		case Right:
			cells[i] = cell(a.Start.Row, a.Start.Col+i)
		}
	}

	return cells
}

func (b Board) Remove(c *Cell, v int) error {
	b.bits[b.valueIndex(c, v)] = false

	possibles, firstPos := b.NumPossibles(c)
	if possibles == 0 {
		return errUnsolvableCell
	}

	if possibles == 1 {
		//fmt.Printf("%v can now only be %d\n", c, firstPos)
		err := b.Set(c, firstPos)
		if err != nil {
			return errUnsolvableCell
		}
	}

	//TODO: Check are there any values for which in this Row or column
	//there is now only a single candidate Cell?

	return nil
}

func (b Board) Init(c *Cell, v int) {
	b.bits[b.valueIndex(c, v)] = true
}

func (b Board) Set(c *Cell, v int) error {
	//fmt.Printf("Setting %v to %d\n", c, v)
	cellIndex := b.cellIndex(c)
	for s := 0; s < b.size; s++ {
		b.bits[cellIndex+s] = s == v-1
	}

	//Update every other Cell in this Row - cannot be v
	for row := 0; row < b.size; row++ {
		target := cell(row, c.Col)
		if target != c && b.Includes(target, v) {
			//fmt.Printf("Removing %v from %d\n", v, target)
			err := b.Remove(target, v)
			if err != nil {
				return errUnsolvableCell
			}
		}
	}

	//Update every other Cell in this column - cannot be v
	for col := 0; col < b.size; col++ {
		target := cell(c.Row, col)
		if target != c && b.Includes(target, v) {
			//fmt.Printf("Removing %v from %d\n", v, target)
			err := b.Remove(target, v)
			if err != nil {
				return errUnsolvableCell
			}
		}
	}

	return nil
}

func (b Board) Copy() Board {
	bitsCopy := append(b.bits[:0:0], b.bits...)
	bCopy := Board{bits: bitsCopy, size: b.size}
	return bCopy
}

func (b Board) Collapse() [][]int {
	result := make([][]int, b.size)
	for r := 0; r < b.size; r++ {
		result[r] = make([]int, b.size)
		for c := 0; c < b.size; c++ {
			_, v := b.NumPossibles(cell(r, c))
			result[r][c] = v
		}
	}

	return result
}

func NewBoard(size int) Board {
	initCellPool(size)
	b := Board{bits: make([]bool, size*size*size), size: size}
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
				b.Init(cell(r, c), s)
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
				if b.Includes(cell(r, c), v+1) {
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

type ClueSolution []map[int]bool

func (cs ClueSolution) Copy() ClueSolution {
	var c ClueSolution = make([]map[int]bool, len(cs))
	for i, existing := range cs {
		c[i] = make(map[int]bool)
		for k, v := range existing {
			c[i][k] = v
		}
	}

	return c
}

func (cs ClueSolution) SetSolution(i int, v int) error {

	cs[i] = map[int]bool{v: true}

	//Update indexes ahead so they can't have this value
	for j, possibles := range cs[i+1:] {
		if possibles[v] {
			delete(possibles, v)
			if len(possibles) == 0 {
				return fmt.Errorf("no possible soluitions left")
			}

			//Removing a value might mean this value *has* to be something
			//if so - remove that option for future possibles
			if len(possibles) == 1 {
				e := cs.SetSolution(i+1+j, firstKey(possibles))
				if e != nil {
					return e
				}
			}
		}
	}

	return nil
}

func (cs ClueSolution) Flatten() ([]int, int) {
	flattened := make([]int, len(cs))

	set := make(map[int]bool)
	visible := 0
	maxHeight := 0
	for i, s := range cs {
		height := firstKey(s)
		if _, ok := set[height]; ok {
			panic("Already have")
		}
		set[height] = true
		flattened[i] = height
		if height > maxHeight {
			visible++
			maxHeight = height
		}
	}

	return flattened, visible
}

type Clue struct {
	visible   int
	cells     []*Cell
	solutions [][]int
}

func resolveClues(b *Board, clues []*Clue) {
	//fmt.Printf("Resolving clues from %v:", b)
	for _, clue := range clues {
		var starting ClueSolution = make([]map[int]bool, 0)

		for _, c := range clue.cells {
			possibles := make(map[int]bool)
			for _, p := range b.Possibles(c) {
				possibles[p] = true
			}

			starting = append(starting, possibles)
		}

		resolveClue(clue, starting)
		if len(clue.solutions) == 0 {
			panic("No solutions found for clue")
		}
	}
	//see clue solution from initial board state
	//iterate through clue and fill in possibles
}

func firstKey(m map[int]bool) int {
	for k := range m {
		return k
	}
	panic("nothing in map")
}

func resolveClue(clue *Clue, solution ClueSolution) {

	//TODO: Could also bail out early here - calculate total
	//visible, if greater than clue then it can never go down
	complete := true
	for i, possibles := range solution {
		if len(possibles) == 1 {
			continue
		}

		complete = false
		for v := range possibles {
			solutionCopy := solution.Copy()
			err := solutionCopy.SetSolution(i, v)
			if err != nil {
				continue
			}

			//TODO: Could pass in index to save some iteration
			resolveClue(clue, solutionCopy)
		}

		break
	}

	if complete {
		flattened, visible := solution.Flatten()
		if visible == clue.visible {
			clue.solutions = append(clue.solutions, flattened)
		} else {
			//fmt.Printf("Discarding solution %v for %v\n", solution, clue)
		}
	}
}

func (c *Clue) String() string {
	s := ""
	for _, cell := range c.cells {
		s += fmt.Sprintf(" %v ", cell)
	}

	return fmt.Sprintf("%d visible in %s", c.visible, s)
}

func findSolutionClues(b *Board, clues []*Clue) *Board {

	clue := clues[0]
	for _, solution := range clue.solutions {
		bCopy := b.Copy()

		appliedClean := true
		for cellIndex, value := range solution {
			c := clue.cells[cellIndex]

			if !bCopy.Includes(c, value) {
				appliedClean = false
				break
			}

			err := bCopy.Set(c, value)
			if err != nil {
				appliedClean = false
				break
			}
		}

		if appliedClean {
			if len(clues) == 1 {
				return &bCopy
			}

			sol := findSolutionClues(&bCopy, clues[1:])
			if sol != nil {
				return sol
			}
		}
	}

	return nil
}
