package kata

import (
	"errors"
	"fmt"
	"sort"
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
		func(clue *Clue, b *Board) {
			if clue.visible == 1 {
				_ = b.Set(clue.cells[0], 4)
			} else if clue.visible == 2 {
				_ = b.Remove(clue.cells[0], 4)
			} else if clue.visible == 3 {
				_ = b.Remove(clue.cells[0], 4)
				_ = b.Remove(clue.cells[0], 3)
				_ = b.Remove(clue.cells[1], 4)
			} else if clue.visible == 4 {
				_ = b.Set(clue.cells[0], 1)
				_ = b.Set(clue.cells[1], 2)
				_ = b.Set(clue.cells[2], 3)
				_ = b.Set(clue.cells[3], 4)
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
		func(clue *Clue, b *Board) {
			if clue.visible == 1 {
				_ = b.Set(clue.cells[0], 6)
			} else if clue.visible == 2 {
				_ = b.Remove(clue.cells[0], 6)
			} else if clue.visible == 3 {
				_ = b.Remove(clue.cells[0], 6)
				_ = b.Remove(clue.cells[1], 6)

				_ = b.Remove(clue.cells[0], 5)
			} else if clue.visible == 4 {
				_ = b.Remove(clue.cells[0], 6)
				_ = b.Remove(clue.cells[1], 6)
				_ = b.Remove(clue.cells[2], 6)

				_ = b.Remove(clue.cells[0], 5)
				_ = b.Remove(clue.cells[1], 5)

				_ = b.Remove(clue.cells[0], 4)
			} else if clue.visible == 5 {
				_ = b.Remove(clue.cells[0], 6)
				_ = b.Remove(clue.cells[1], 6)
				_ = b.Remove(clue.cells[2], 6)
				_ = b.Remove(clue.cells[3], 6)

				_ = b.Remove(clue.cells[0], 5)
				_ = b.Remove(clue.cells[1], 5)
				_ = b.Remove(clue.cells[2], 5)

				_ = b.Remove(clue.cells[0], 4)
				_ = b.Remove(clue.cells[1], 4)

				_ = b.Remove(clue.cells[0], 3)

			} else if clue.visible == 6 {
				_ = b.Set(clue.cells[0], 1)
				_ = b.Set(clue.cells[1], 2)
				_ = b.Set(clue.cells[2], 3)
				_ = b.Set(clue.cells[3], 4)
				_ = b.Set(clue.cells[4], 5)
				_ = b.Set(clue.cells[5], 6)
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
		func(clue *Clue, b *Board) {
			if clue.visible == 1 {
				_ = b.Set(clue.cells[0], 7)
			} else if clue.visible == 2 {
				_ = b.Remove(clue.cells[0], 7)
			} else if clue.visible == 3 {
				_ = b.Remove(clue.cells[0], 7)
				_ = b.Remove(clue.cells[1], 7)

				_ = b.Remove(clue.cells[0], 6)
			} else if clue.visible == 4 {
				_ = b.Remove(clue.cells[0], 7)
				_ = b.Remove(clue.cells[1], 7)
				_ = b.Remove(clue.cells[2], 7)

				_ = b.Remove(clue.cells[0], 6)
				_ = b.Remove(clue.cells[1], 6)

				_ = b.Remove(clue.cells[0], 5)
			} else if clue.visible == 5 {
				_ = b.Remove(clue.cells[0], 7)
				_ = b.Remove(clue.cells[1], 7)
				_ = b.Remove(clue.cells[2], 7)
				_ = b.Remove(clue.cells[3], 7)

				_ = b.Remove(clue.cells[0], 6)
				_ = b.Remove(clue.cells[1], 6)
				_ = b.Remove(clue.cells[2], 6)

				_ = b.Remove(clue.cells[0], 5)
				_ = b.Remove(clue.cells[1], 5)

				_ = b.Remove(clue.cells[0], 4)

			} else if clue.visible == 6 {
				_ = b.Remove(clue.cells[0], 7)
				_ = b.Remove(clue.cells[1], 7)
				_ = b.Remove(clue.cells[2], 7)
				_ = b.Remove(clue.cells[3], 7)
				_ = b.Remove(clue.cells[4], 7)

				_ = b.Remove(clue.cells[0], 6)
				_ = b.Remove(clue.cells[1], 6)
				_ = b.Remove(clue.cells[2], 6)
				_ = b.Remove(clue.cells[3], 6)

				_ = b.Remove(clue.cells[0], 5)
				_ = b.Remove(clue.cells[1], 5)
				_ = b.Remove(clue.cells[2], 5)

				_ = b.Remove(clue.cells[0], 4)
				_ = b.Remove(clue.cells[1], 4)

				_ = b.Remove(clue.cells[0], 3)

			} else if clue.visible == 7 {
				_ = b.Set(clue.cells[0], 1)
				_ = b.Set(clue.cells[1], 2)
				_ = b.Set(clue.cells[2], 3)
				_ = b.Set(clue.cells[3], 4)
				_ = b.Set(clue.cells[4], 5)
				_ = b.Set(clue.cells[5], 6)
				_ = b.Set(clue.cells[6], 7)
			}
		},
		clues,
	)
}

func solve(b *Board, clueMap []*CellRange, deduce Deducer, clueValues []int) [][]int {

	clues := NewClues(clueValues, clueMap, b)
	for _, clue := range clues {
		deduce(clue, b)
	}

	sol := solveNextClue(b, clues)
	if sol == nil {
		panic(errorNoSolution)
	}

	return sol.Flatten()
}

func solveNextClue(b *Board, clues []*Clue) *Board {

	visible, maxHeight, nextClue := 0, 0, clues[0]

	//Go through cell that is part of this clue
	for _, c := range nextClue.cells {
		numPossible, height := b.NumPossibles(c)

		//If the number of possible options is only 1,
		if numPossible == 1 {
			//Increase the number skyscrapers we can see from this
			//clue if the height is greater than the current max height
			if height > maxHeight {
				visible++
				maxHeight = height
				//If we can see more skyscrapers than we should,
				//just exit now - as it's not possible for this number
				//to go down no matter what we do next
				if visible > nextClue.visible {
					break
				}
			}
		} else {

			//If there is more than 1 possible option for each cell, then
			//iterate through each possibility, set that as the value for
			//the cell and then recurse to resolve the remaining cells
			for _, v := range b.Possibles(c) {

				boardCopy := b.Copy()
				err := boardCopy.Set(c, v)
				if err != nil {
					continue
				}

				solution := solveNextClue(boardCopy, clues)
				if solution != nil {
					return solution
				}
			}

			//We've filled in all possible values for this clue
			//but none of these resulted in a possible solution, so
			//return nil.
			return nil
		}

	}

	//We've checked all cells,  they all have only 1 possible
	//value. Check that the visible skyscrapers match what we
	//should have for this clue
	if nextClue.visible != visible {
		return nil
	}

	//If there are more clues to resolved, recurse and move on to the
	//next clue
	if len(clues) > 1 {
		return solveNextClue(b, clues[1:])
	}

	//This was the last clue, we have our solution
	return b
}

var errUnsolvableCell = errors.New("cell has no remaining values")
var errorNoSolution = errors.New("no solution")

type Deducer func(clue *Clue, b *Board)

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

type Clue struct {
	visible int
	cells   []*Cell
}

type Clues []*Clue

func NewClues(values []int, clueMap []*CellRange, b *Board) Clues {
	var clues Clues = make([]*Clue, 0)
	for i, value := range values {
		if value == 0 {
			continue
		}

		clues = append(clues, &Clue{value, b.GetCells(clueMap[i])})
	}

	sort.Sort(clues)

	return clues
}

func (c Clues) Len() int {
	return len(c)
}

func (c Clues) Less(i, j int) bool {
	return c[i].visible > c[j].visible

}
func (c Clues) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
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

type Board struct {
	bits []bool
	size int
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
	num, nextPossible := 0, 0
	cellIndex := b.cellIndex(c)
	for idx, isSet := range b.bits[cellIndex : cellIndex+b.size] {
		if !isSet {
			continue
		}

		num++
		if num == 1 {
			nextPossible = idx + 1
		}
	}

	return num, nextPossible
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
		err := b.Set(c, firstPos)
		if err != nil {
			return errUnsolvableCell
		}
	}

	return nil
}

func (b Board) Init(c *Cell, v int) {
	b.bits[b.valueIndex(c, v)] = true
}

func (b Board) Set(c *Cell, v int) error {
	cellIndex := b.cellIndex(c)
	for s := 0; s < b.size; s++ {
		b.bits[cellIndex+s] = s == v-1
	}

	for row := 0; row < b.size; row++ {
		target := cell(row, c.Col)
		if target != c && b.Includes(target, v) {
			err := b.Remove(target, v)
			if err != nil {
				return err
			}
		}
	}

	for col := 0; col < b.size; col++ {
		target := cell(c.Row, col)
		if target != c && b.Includes(target, v) {
			err := b.Remove(target, v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b Board) Copy() *Board {
	bitsCopy := append(b.bits[:0:0], b.bits...)
	bCopy := Board{bits: bitsCopy, size: b.size}
	return &bCopy
}

func (b Board) Flatten() [][]int {
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
