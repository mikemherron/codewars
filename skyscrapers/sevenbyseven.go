//https://www.codewars.com/kata/5671d975d81d6c1c87000022
package skyscrapers

var clueMapSevenBySeven = []*CellRange{
	{Cell{0, 0}, Down},
	{Cell{0, 1}, Down},
	{Cell{0, 2}, Down},
	{Cell{0, 3}, Down},
	{Cell{0, 4}, Down},
	{Cell{0, 5}, Down},
	{Cell{0, 6}, Down},

	{Cell{0, 6}, Left},
	{Cell{1, 6}, Left},
	{Cell{2, 6}, Left},
	{Cell{3, 6}, Left},
	{Cell{4, 6}, Left},
	{Cell{5, 6}, Left},
	{Cell{6, 6}, Left},

	{Cell{6, 6}, Up},
	{Cell{6, 5}, Up},
	{Cell{6, 4}, Up},
	{Cell{6, 3}, Up},
	{Cell{6, 2}, Up},
	{Cell{6, 1}, Up},
	{Cell{6, 0}, Up},

	{Cell{6, 0}, Right},
	{Cell{5, 0}, Right},
	{Cell{4, 0}, Right},
	{Cell{3, 0}, Right},
	{Cell{2, 0}, Right},
	{Cell{1, 0}, Right},
	{Cell{0, 0}, Right},
}

func SolvePuzzlSevenBySeven(clues []int) [][]int {
	return SolvePuzzleSevenBySevenFrom(clues, nil)
}

func SolvePuzzleSevenBySevenFrom(clues []int, starting [][][]int) [][]int {

	b := NewBoardFrom(7, starting)

	clues2 := make([]*Clue, 0)

	for i, clue := range clues {
		if clue == 0 {
			continue
		}

		cells := b.GetCells(clueMapSevenBySeven[i])
		clues2 = append(clues2, &Clue{clue, cells})
		//Ignore errors here - assume that applying
		//the initial deductions from the clues
		//will not result in an invalid state :)
		if clue == 1 {
			_ = b.Set(*cells[0], 7)
		} else if clue == 2 {
			_ = b.Remove(*cells[0], 7)
		} else if clue == 3 {
			_ = b.Remove(*cells[0], 7)
			_ = b.Remove(*cells[1], 7)

			_ = b.Remove(*cells[0], 6)
		} else if clue == 4 {
			_ = b.Remove(*cells[0], 7)
			_ = b.Remove(*cells[1], 7)
			_ = b.Remove(*cells[2], 7)

			_ = b.Remove(*cells[0], 6)
			_ = b.Remove(*cells[1], 6)

			_ = b.Remove(*cells[0], 5)
		} else if clue == 5 {
			_ = b.Remove(*cells[0], 7)
			_ = b.Remove(*cells[1], 7)
			_ = b.Remove(*cells[2], 7)
			_ = b.Remove(*cells[3], 7)

			_ = b.Remove(*cells[0], 6)
			_ = b.Remove(*cells[1], 6)
			_ = b.Remove(*cells[2], 6)

			_ = b.Remove(*cells[0], 5)
			_ = b.Remove(*cells[1], 5)

			_ = b.Remove(*cells[0], 4)

		} else if clue == 6 {
			_ = b.Remove(*cells[0], 7)
			_ = b.Remove(*cells[1], 7)
			_ = b.Remove(*cells[2], 7)
			_ = b.Remove(*cells[3], 7)
			_ = b.Remove(*cells[4], 7)

			_ = b.Remove(*cells[0], 6)
			_ = b.Remove(*cells[1], 6)
			_ = b.Remove(*cells[2], 6)
			_ = b.Remove(*cells[3], 6)

			_ = b.Remove(*cells[0], 5)
			_ = b.Remove(*cells[1], 5)
			_ = b.Remove(*cells[2], 5)

			_ = b.Remove(*cells[0], 4)
			_ = b.Remove(*cells[1], 4)

			_ = b.Remove(*cells[0], 3)

		} else if clue == 7 {
			_ = b.Set(*cells[0], 1)
			_ = b.Set(*cells[1], 2)
			_ = b.Set(*cells[2], 3)
			_ = b.Set(*cells[3], 4)
			_ = b.Set(*cells[4], 5)
			_ = b.Set(*cells[5], 6)
			_ = b.Set(*cells[6], 7)
		}
	}

	sol := findSolutionStack(&b, clues2)
	if sol == nil {
		panic("unable to find solution")
	}

	return sol.Collapse()
}
