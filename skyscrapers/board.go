package skyscrapers

import (
	"errors"
	"fmt"
	"math"
)

var errUnsolvableCell = errors.New("cell has no remaining values")

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

type BitSet struct {
	bits1 uint64 //64
	bits2 uint64 //128
	bits3 uint64 //192
	bits4 uint64 //256
	bits5 uint64 //320
	bits6 uint64 //384
}

func NewBitSet() *BitSet {
	return &BitSet{
		bits1: math.MaxUint64,
		bits2: math.MaxUint64,
		bits3: math.MaxUint64,
		bits4: math.MaxUint64,
		bits5: math.MaxUint64,
		bits6: math.MaxUint64,
	}
}

func (b *BitSet) bitsForPos(pos int) *uint64 {
	if pos > 320 {
		return &b.bits6
	} else if pos > 256 {
		return &b.bits5
	} else if pos > 192 {
		return &b.bits4
	} else if pos > 128 {
		return &b.bits3
	} else if pos > 64 {
		return &b.bits2
	} else if pos > 0 {
		return &b.bits1
	}

	panic("invalid pos")
}

func (b *BitSet) set(pos int) {
	bits := b.bitsForPos(pos)
	relPos := pos % 64
	*bits |= 1 << relPos
}

func (b *BitSet) clear(pos int) {
	bits := b.bitsForPos(pos)
	relPos := pos % 64
	*bits &= ^(1 << relPos)
}

func (b *BitSet) has(pos int) bool {
	bits := b.bitsForPos(pos)
	relPos := pos % 64
	val := *bits & (1 << relPos)
	return val > 0
}

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
	//copy(bCopy.bits, b.bits)
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
	b := Board{bits: make([]bool, size*size*size), size: size}
	for i := range b.bits {
		b.bits[i] = true
	}

	return b
}

func NewBoardFrom(size int, initial [][][]int) Board {
	b := NewBoard(size)
	if initial == nil {
		return b
	}

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			for _, s := range initial[r][c] {
				b.Init(cell(r, c), s)
			}
		}
	}

	return b
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
