//https://www.codewars.com/kata/5671d975d81d6c1c87000022
package skyscrapers

import "fmt"

//var errUnsolvableCell = errors.New("cell has no remaining values")

var clueMap = []*CellRange{
	{Cell{0, 0}, Down},
	{Cell{0, 1}, Down},
	{Cell{0, 2}, Down},
	{Cell{0, 3}, Down},

	{Cell{0, 3}, Left},
	{Cell{1, 3}, Left},
	{Cell{2, 3}, Left},
	{Cell{3, 3}, Left},

	{Cell{3, 3}, Up},
	{Cell{3, 2}, Up},
	{Cell{3, 1}, Up},
	{Cell{3, 0}, Up},

	{Cell{3, 0}, Right},
	{Cell{2, 0}, Right},
	{Cell{1, 0}, Right},
	{Cell{0, 0}, Right},
}

func SolvePuzzle(clues []int) [][]int {
	return SolvePuzzleFrom(clues, nil)
}

func SolvePuzzleFrom(clueValues []int, starting [][][]int) [][]int {

	b := NewBoardFrom(4, starting)

	clues := make([]*Clue, 0)

	for i, clue := range clueValues {
		if clue == 0 {
			continue
		}

		cells := b.GetCells(clueMap[i])

		clues = append(clues, &Clue{clue, cells})

		//Ignore errors here - assume that applying
		//the initial deductions from the clueValues
		//will not result in an invalid state :)
		if clue == 1 {
			_ = b.Set(*cells[0], 4)
		} else if clue == 2 {
			_ = b.Remove(*cells[0], 4)
		} else if clue == 3 {
			_ = b.Remove(*cells[0], 4)
			_ = b.Remove(*cells[0], 3)
			_ = b.Remove(*cells[1], 4)
		} else if clue == 4 {
			_ = b.Set(*cells[0], 1)
			_ = b.Set(*cells[1], 2)
			_ = b.Set(*cells[2], 3)
			_ = b.Set(*cells[3], 4)
		}
	}

	sol := findSolutionStack(&b, clues)
	if sol == nil {
		panic("unable to find solution")
	}

	return sol.Collapse()
}

type Clue struct {
	visible int
	cells   []*Cell
}

func (c *Clue) String() string {
	s := ""
	for _, cell := range c.cells {
		s += fmt.Sprintf(" %v ", cell)
	}

	return fmt.Sprintf("%d visible in %s", c.visible, s)
}

func findSolutionStack(b *Board, clues []*Clue) *Board {

	visible := 0
	maxHeight := 0

	clue := clues[0]

	//fmt.Printf("Checking clue %v, %d left\n", clue, len(clues))

	for _, cell := range clue.cells {

		//fmt.Printf("-Checking cell %v\n", cell)
		numberPossible, height := b.NumPossibles(*cell)
		if numberPossible > 1 {

			var s *Board
			b.PossiblesCb(*cell, func(v int) {
				if s != nil {
					return
				}
				boardCopy := b.Copy()
				err := boardCopy.Set(*cell, v)
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
