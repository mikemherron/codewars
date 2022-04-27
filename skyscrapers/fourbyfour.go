//https://www.codewars.com/kata/5671d975d81d6c1c87000022
package skyscrapers

import (
	"fmt"
)

//Other option - pregenerate all permuations with their visible.
//then try each - disadvnatag eis you don't have the specific
//clues the initial board dedication provides

//func generate(visible []int, size int) map[int][][]int {
//	result := make(map[int][][]int)
//	for _, v := range visible {
//		result[v] = make([][]int, 0)
//	}
//
//	pool := make([]int, size)
//	for i := 1; i <= size; i++ {
//		pool = append(pool, i)
//	}
//
//	return result
//}
//
//func gen()

//var errUnsolvableCell = errors.New("cell has no remaining values")

var clueMap []*CellRange

func SolvePuzzle(clues []int) [][]int {
	return SolvePuzzleFrom(clues, nil)
}

func SolvePuzzleFrom(clueValues []int, starting [][][]int) [][]int {

	initCellPool(4)

	clueMap = []*CellRange{
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
	}

	b := NewBoardFrom(4, starting)

	clues := make([]*Clue, 0)

	for i, clue := range clueValues {
		if clue == 0 {
			continue
		}

		cells := b.GetCells(clueMap[i])

		clues = append(clues, &Clue{clue, cells, make([][]int, 0)})

		//Ignore errors here - assume that applying
		//the initial deductions from the clueValues
		//will not result in an invalid state :)
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
	}

	resolveClues(&b, clues)

	sol := findSolutionClues(&b, clues)
	if sol == nil {
		panic("unable to find solution")
	}
	/*sol := findSolutionStack(&b, clues)
	if sol == nil {
		panic("unable to find solution")
	}*/

	return sol.Collapse()
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
