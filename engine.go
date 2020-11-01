package main

import (
	"fmt"
	"reflect"
)

const GapSymbol = "-"

type HirschbergSolver struct {
	TopSequence   *Sequence
	LeftSequence  *Sequence
	TopSequenceR  *Sequence
	LeftSequenceR *Sequence
	TopVector     []*Cell
	BottomVector  []*Cell
	SF            ScoringFunc
	GapValue      int
}

func (hs *HirschbergSolver) Director(leftFrom, topFrom, leftTo, topTo int, downshift bool) []Direction {
	//fmt.Println("LEFT: ", leftFrom,":",leftTo)
	//fmt.Println("TOP: ", topFrom,":",topTo)
	if topTo - topFrom == 0 || leftTo - leftFrom == 0 {
		return []Direction{}
	}
	if leftTo - leftFrom == 1 {
		directions := make([]Direction, topTo-topFrom-1)
		for i := 0; i < topTo-topFrom-1; i++ {
			directions[i] = LeftDirection
		}
		return directions
	}
	if topTo - topFrom == 1 {
		directions := make([]Direction, leftTo - leftFrom-1)
		for i := 0; i < leftTo - leftFrom-1; i++ {
			directions[i] = TopDirection
		}
		return directions
	}

	var leftStr, topStr string
	leftStr = hs.LeftSequence.Value[leftFrom:leftTo]
	topStr = hs.TopSequence.Value[topFrom:topTo]

	countDown := (leftTo - leftFrom + 1) / 2
	countUp := len(leftStr) - countDown
	var up, down []int
	//if !downshift {
	//	topStr = reverseString(topStr)
	//	leftStr = reverseString(leftStr)
	//}
	up = hs.Downshift(topStr, leftStr[:countDown])
	down = hs.Downshift(reverseString(topStr), reverseString(leftStr)[:countUp])
	reverseAny(down)

	maxVal := up[1] + down[0]
	maxI := 1
	action := TopDirection

	for i := 1; i < len(up); i++ {
		val := up[i] + down[i-1] + hs.GapValue
		if val > maxVal {
			maxVal = val
			maxI = i
			action = TopDirection
		}

		if i < len(up) - 1 {
			val = up[i] + down[i] + hs.SF[leftStr[countDown]][topStr[i]]
			if val > maxVal {
				maxVal = val
				maxI = i
				action = DiagonalDirection
			}
		}
	}

	directions := hs.Director(leftFrom, topFrom, leftFrom+countDown, topFrom+maxI, true)
	directions = append(directions, action)
	if action == DiagonalDirection {
		maxI++
	}
	directions = append(directions, hs.Director(leftFrom+countDown, topFrom+maxI-1, leftTo, topTo, false)...)

	//for _, direction := range directions {
	//	fmt.Println(printDirection(direction))
	//}

	return directions
}

func (hs *HirschbergSolver) Downshift(topSeq, leftSeq string) []int {
	N := len(topSeq) + 1
	count := len(leftSeq)

	first, second := make([]int, N), make([]int, N)

	for i := 0; i <= len(topSeq); i++ {
		first[i] = i * hs.GapValue
	}

	for k := 0; k < count; k++ {
		//for _, cell := range first {
		//	fmt.Print(cell, " ")
		//}
		//fmt.Println()

		for i := 0; i < N; i++ {
			top := first[i] + hs.GapValue
			second[i] = top
			if i > 0 {
				diag := first[i-1] + hs.SF[leftSeq[k]][topSeq[i-1]]
				left := second[i-1] + hs.GapValue
				if diag >= left && diag >= top {
					second[i] = diag
				} else if left > diag && left > top {
					second[i] = left
				}
			}
		}
		copy(first, second)
	}
	//for _, cell := range second {
	//	fmt.Print(cell, " ")
	//}
	//fmt.Println()
	return second
}

func Hirschberg(first, second *Sequence, sf ScoringFunc, GapValue int) *HirschbergSolver {
	nw := &HirschbergSolver{
		TopSequence:  second,
		LeftSequence: first,
		SF:           sf,
		GapValue:     GapValue,
	}

	//for i := 0; i <= len(nw.TopSequence.Value); i++ {
	//	dir := LeftDirection
	//	if i == 0 {
	//		dir = NullDirection
	//	}
	//	nw.TopVector[i] = &Cell{
	//		Distance: i * GapValue,
	//		Dir:      dir,
	//	}
	//	nw.BottomVector[i] = &Cell{
	//		Distance: i * GapValue,
	//		Dir:      dir,
	//	}
	//}

	nw.TopSequenceR = &Sequence{
		ID:          nw.TopSequence.ID,
		Description: nw.TopSequence.Description,
		Value:       reverseString(nw.TopSequence.Value),
	}

	nw.LeftSequenceR = &Sequence{
		ID:          nw.LeftSequence.ID,
		Description: nw.LeftSequence.Description,
		Value:       reverseString(nw.LeftSequence.Value),
	}

	return nw
}

func (hs *HirschbergSolver) Solve() {
	//countDown := (len(hs.LeftSequence.Value) + 1) / 2
	//countUp := len(hs.LeftSequence.Value) - countDown
	//up := hs.Downshift(hs.TopSequence.Value, hs.LeftSequence.Value[:countDown])
	//down := hs.Downshift(hs.TopSequenceR.Value, hs.LeftSequenceR.Value[:countUp])
	//reverseAny(down)
	//
	//maxVal := up[1] + down[0]
	//maxI := 1
	//action := TopDirection
	//
	//for i := 1; i < len(up); i++ {
	//	val := up[i] + down[i-1] + hs.GapValue
	//	if val > maxVal {
	//		maxVal = val
	//		maxI = i
	//		action = TopDirection
	//	}
	//
	//	if i < len(up) - 1 {
	//		val = up[i] + down[i] + hs.SF[hs.LeftSequence.Value[countDown]][hs.TopSequence.Value[i]]
	//		fmt.Println(i, up[i], down[i], string(rune(hs.LeftSequence.Value[countDown])), string(rune(hs.TopSequence.Value[i])))
	//		if val > maxVal {
	//			maxVal = val
	//			maxI = i
	//			action = DiagonalDirection
	//		}
	//	}
	//}
	//
	//fmt.Println(up[1:])
	//fmt.Println()
	//
	//fmt.Println(down[:len(down)-1])
	//fmt.Println()
	//
	//fmt.Println(maxVal)
	//fmt.Println(printDirection(action))
	//fmt.Println(maxI)
	//
	//directions := hs.Director(0, 0, countDown, maxI, true)
	//directions = append(directions, action)
	//if action == DiagonalDirection {
	//	maxI++
	//}
	//directions = append(directions, hs.Director(countDown, maxI-1, len(hs.LeftSequence.Value), len(hs.TopSequence.Value), false)...)
	//
	//for _, direction := range directions {
	//	fmt.Println(printDirection(direction))
	//}
	directions := hs.Director(0, 0, len(hs.LeftSequence.Value), len(hs.TopSequence.Value), true)
	for _, direction := range directions {
		fmt.Println(printDirection(direction))
	}
}

// Функция вывода таблицы для отладки
//func (nw *HirschbergSolver) Print() {
//	for i := 0; i <= len(nw.LeftSequence.Value); i++ {
//		for j := 0; j <= len(nw.TopSequence.Value); j++ {
//			fmt.Print(nw.Table[i][j].Distance, ", ", nw.Table[i][j].Dir)
//			fmt.Print("   | ")
//		}
//		fmt.Println()
//	}
//}
//
//func (nw *HirschbergSolver) Solve() (string, string, int) {
//	// Рекурсивно определяем значения матрицы, одновременно определяя и направления
//	nw.determine(len(nw.LeftSequence.Value), len(nw.TopSequence.Value))
//
//	cell := nw.Table[len(nw.LeftSequence.Value)][len(nw.TopSequence.Value)]
//
//	score := cell.Distance
//	firstRes, secondRes := "", ""
//
//	// Двигаемся от правой нижней ячейки матрицы к левой верхней, и строим с конца строки-выравнивания
//	fp, sp := len(nw.LeftSequence.Value)-1, len(nw.TopSequence.Value)-1
//	for cell.Dir != NullDirection {
//		if cell.Dir == DiagonalDirection {
//			firstRes = string(rune(nw.LeftSequence.Value[fp])) + firstRes
//			secondRes = string(rune(nw.TopSequence.Value[sp])) + secondRes
//			sp--
//			fp--
//		} else if cell.Dir == TopDirection {
//			firstRes = string(rune(nw.LeftSequence.Value[fp])) + firstRes
//			secondRes = GapSymbol + secondRes
//			fp--
//		} else if cell.Dir == LeftDirection {
//			firstRes = GapSymbol + firstRes
//			secondRes = string(rune(nw.TopSequence.Value[sp])) + secondRes
//			sp--
//		}
//		cell = nw.Table[fp+1][sp+1]
//	}
//
//	return firstRes, secondRes, score
//}
//
//// Рекурсивное заполнение матрицы
//func (nw *HirschbergSolver) determine(i, j int) {
//	if nw.Table[i][j] != nil {
//		return
//	}
//	leftCell, topCell, diagCell := nw.Table[i][j-1], nw.Table[i-1][j], nw.Table[i-1][j-1]
//	if leftCell == nil {
//		nw.determine(i, j-1)
//		leftCell = nw.Table[i][j-1]
//	}
//	if diagCell == nil {
//		nw.determine(i-1, j-1)
//		diagCell = nw.Table[i-1][j-1]
//	}
//	if topCell == nil {
//		nw.determine(i-1, j)
//		topCell = nw.Table[i-1][j]
//	}
//
//	maxVal, maxNum := max3(
//		nw.Table[i-1][j-1].Distance+nw.SF[nw.LeftSequence.Value[i-1]][nw.TopSequence.Value[j-1]],
//		nw.Table[i][j-1].Distance+nw.GapValue,
//		nw.Table[i-1][j].Distance+nw.GapValue,
//	)
//
//	nw.Table[i][j] = &Cell{
//		Distance: maxVal,
//	}
//	curCell := nw.Table[i][j]
//
//	switch maxNum {
//	case 1:
//		curCell.Dir = DiagonalDirection
//	case 2:
//		curCell.Dir = LeftDirection
//	case 3:
//		curCell.Dir = TopDirection
//	}
//}
//
//// Достаточно примитивная функция выбора максимума из трех с указанием номера максимального
//func max3(a, b, c int) (int, int) {
//	if a >= b {
//		if a >= c {
//			return a, 1
//		}
//		return c, 3
//	}
//	if b >= c {
//		return b, 2
//	}
//	return c, 3
//}
//
//func max2(a, b int) (int, bool) {
//	if a >= b {
//		return a, true
//	}
//	return b, false
//}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func printDirection(i Direction) string {
	switch i {
	case TopDirection:
		return "top"
	case LeftDirection:
		return "left"
	case DiagonalDirection:
		return "diagonal"
	default:
		return "nill"
	}
}
