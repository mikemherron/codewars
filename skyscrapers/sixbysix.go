//https://www.codewars.com/kata/5671d975d81d6c1c87000022
package skyscrapers

var clueMapSixBySix = []*CellRange{
	{Cell{0, 0}, Down},
	{Cell{0, 1}, Down},
	{Cell{0, 2}, Down},
	{Cell{0, 3}, Down},
	{Cell{0, 4}, Down},
	{Cell{0, 5}, Down},

	{Cell{0, 5}, Left},
	{Cell{1, 5}, Left},
	{Cell{2, 5}, Left},
	{Cell{3, 5}, Left},
	{Cell{4, 5}, Left},
	{Cell{5, 5}, Left},

	{Cell{5, 5}, Up},
	{Cell{5, 4}, Up},
	{Cell{5, 3}, Up},
	{Cell{5, 2}, Up},
	{Cell{5, 1}, Up},
	{Cell{5, 0}, Up},

	{Cell{5, 0}, Right},
	{Cell{4, 0}, Right},
	{Cell{3, 0}, Right},
	{Cell{2, 0}, Right},
	{Cell{1, 0}, Right},
	{Cell{0, 0}, Right},
}

func SolvePuzzleSixBySix(clues []int) [][]int {
	return SolvePuzzleSixBySixFrom(clues, nil)
}

func SolvePuzzleSixBySixFrom(clues []int, starting [][][]int) [][]int {

	b := NewBoardFrom(6, starting)

	clues2 := make([]*Clue, 0)

	for i, clue := range clues {
		if clue == 0 {
			continue
		}

		cells := b.GetCells(clueMapSixBySix[i])
		clues2 = append(clues2, &Clue{clue, cells})
		//Ignore errors here - assume that applying
		//the initial deductions from the clues
		//will not result in an invalid state :)
		if clue == 1 {
			_ = b.Set(*cells[0], 6)
		} else if clue == 2 {
			_ = b.Remove(*cells[0], 6)
		} else if clue == 3 {
			_ = b.Remove(*cells[0], 6)
			_ = b.Remove(*cells[1], 6)

			_ = b.Remove(*cells[0], 5)
		} else if clue == 4 {
			_ = b.Remove(*cells[0], 6)
			_ = b.Remove(*cells[1], 6)
			_ = b.Remove(*cells[2], 6)

			_ = b.Remove(*cells[0], 5)
			_ = b.Remove(*cells[1], 5)

			_ = b.Remove(*cells[0], 4)
		} else if clue == 5 {
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

		} else if clue == 6 {
			_ = b.Set(*cells[0], 1)
			_ = b.Set(*cells[1], 2)
			_ = b.Set(*cells[2], 3)
			_ = b.Set(*cells[3], 4)
			_ = b.Set(*cells[4], 5)
			_ = b.Set(*cells[5], 6)
		}
	}

	sol := findSolutionStack(&b, clues2)
	if sol == nil {
		panic("unable to find solution")
	}

	return sol.Collapse()
}
