package main

import (
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

// Инициализация
func Hirschberg(first, second *Sequence, sf ScoringFunc, GapValue int) *HirschbergSolver {
	nw := &HirschbergSolver{
		TopSequence:  second,
		LeftSequence: first,
		SF:           sf,
		GapValue:     GapValue,
	}

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

// Вход в рекурсию. Построение score по списку направлений движения
func (hs *HirschbergSolver) Solve() (string, string, int) {
	directions := hs.Director(0, 0, len(hs.LeftSequence.Value), len(hs.TopSequence.Value), true)

	firstRes, secondRes := string(rune(hs.LeftSequence.Value[0])), string(rune(hs.TopSequence.Value[0]))
	i, j := 0, 0
	score := hs.SF[hs.LeftSequence.Value[0]][hs.TopSequence.Value[0]]
	for _, direction := range directions {
		switch direction {
		case TopDirection:
			i++
			firstRes = firstRes + string(rune(hs.LeftSequence.Value[i]))
			secondRes = secondRes + GapSymbol
			score += hs.GapValue
		case LeftDirection:
			j++
			score += hs.GapValue
			secondRes = secondRes + string(rune(hs.TopSequence.Value[j]))
			firstRes = firstRes + GapSymbol
		case DiagonalDirection:
			i++
			j++
			score += hs.SF[hs.LeftSequence.Value[i]][hs.TopSequence.Value[j]]
			firstRes = firstRes + string(rune(hs.LeftSequence.Value[i]))
			secondRes = secondRes + string(rune(hs.TopSequence.Value[j]))
		}
	}
	return firstRes, secondRes, score
}

// Рекурсивное построение направлений "движения"
func (hs *HirschbergSolver) Director(leftFrom, topFrom, leftTo, topTo int, downshift bool) []Direction {
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

	return directions
}

// Построение центральной строки воображаемой матрицы снизу вверх или сверху вниз
func (hs *HirschbergSolver) Downshift(topSeq, leftSeq string) []int {
	N := len(topSeq) + 1
	count := len(leftSeq)

	first, second := make([]int, N), make([]int, N)

	for i := 0; i <= len(topSeq); i++ {
		first[i] = i * hs.GapValue
	}

	for k := 0; k < count; k++ {
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
	return second
}

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
