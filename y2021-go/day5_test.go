package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

var DEBUG = false

func initData() ([][]int, [][]int) {
	var vents [][]int
	maxX := 0
	maxY := 0
	for x, line := range fileLines("d5") {
		tokens := strings.Split(line, " -> ")
		x1y1 := strings.Split(tokens[0], ",")
		x2y2 := strings.Split(tokens[1], ",")
		x1, _ := strconv.Atoi(x1y1[0])
		y1, _ := strconv.Atoi(x1y1[1])
		x2, _ := strconv.Atoi(x2y2[0])
		y2, _ := strconv.Atoi(x2y2[1])
		vents = append(vents, []int{x1, y1, x2, y2})
		if DEBUG {
			fmt.Println(x, ">", x1, y1, x2, y2)
		}
		if x1 > maxX {
			maxX = x1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y1 > maxY {
			maxY = y1
		}
		if y2 > maxY {
			maxY = y2
		}
	}
	if DEBUG {
		fmt.Println("maxX", maxX, "maxY", maxY)
		fmt.Println("-------")
	}
	ventsVectors := make([][]int, maxY+1)
	for i := 0; i < maxY+1; i++ {
		ventsVectors[i] = make([]int, maxX+1)
	}

	return vents, ventsVectors
}

func printVectors(vectors [][]int) {
	for _, ventLine := range vectors {
		for _, vent := range ventLine {
			if vent == 0 {
				fmt.Print(".")
			}
			if vent != 0 {
				fmt.Print(vent)

			}
		}
		fmt.Print("\n")
	}
}

func fromTo(a int, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

func computeVector(entry []int, vectors [][]int) [][]int {
	x1, y1, x2, y2 := entry[0], entry[1], entry[2], entry[3]
	//compute Y
	if x1 == x2 {
		if DEBUG {
			print("\\Y/ ", y1 < y2, " y1 ", y1, " y2 ", y2, " x ", x1, "\n")
		}
		from, to := fromTo(y1, y2)
		for y := from; y <= to; y++ {
			vectors[y][x1]++
		}
	}
	//compute X
	if y1 == y2 {
		if DEBUG {
			print("\\X/ ", x1 < x2, " x1 ", x1, " x2 ", x2, " y ", y1, "\n")
		}
		from, to := fromTo(x1, x2)
		for x := from; x <= to; x++ {
			vectors[y1][x]++
		}
	}
	return vectors
}

func TestDay5Step1(t *testing.T) {
	vents, vectors := initData()
	for s, entry := range vents {
		if DEBUG {
			fmt.Printf("\n--- step %d ---\n", s)
		}
		vectors = computeVector(entry, vectors)
		if DEBUG {
			printVectors(vectors)
		}
	}
	totalOverlap := 0
	for _, vector := range vectors {
		for _, point := range vector {
			if point > 1 {
				totalOverlap++
			}
		}
	}
	fmt.Println("d5s1 |", totalOverlap)

}

func computeVectorWithDiagonals(entry []int, vectors [][]int) [][]int {
	x1, y1, x2, y2 := entry[0], entry[1], entry[2], entry[3]
	//compute Y
	if x1 == x2 {
		if DEBUG {
			print("\\Y/ ", y1 < y2, " y1 ", y1, " y2 ", y2, " x ", x1, "\n")
		}
		from, to := fromTo(y1, y2)
		for y := from; y <= to; y++ {
			vectors[y][x1]++
		}
	}
	//compute X
	if y1 == y2 {
		if DEBUG {
			print("\\X/ ", x1 < x2, " x1 ", x1, " x2 ", x2, " y ", y1, "\n")
		}
		from, to := fromTo(x1, x2)
		for x := from; x <= to; x++ {
			vectors[y1][x]++
		}
	}

	diffX := math.Abs(math.Abs(float64(x1)) - math.Abs(float64(x2)))
	diffY := math.Abs(math.Abs(float64(y1)) - math.Abs(float64(y2)))

	if diffX == diffY {

		deltaX := x1
		deltaY := y1

		for d := 0; d < int(diffX)+1; d++ {
			vectors[deltaY][deltaX]++
			if x1 > x2 {
				deltaX--
			} else {
				deltaX++
			}
			if y1 > y2 {
				deltaY--
			} else {
				deltaY++
			}
		}
	}

	return vectors
}

func TestDay5Step2(t *testing.T) {
	vents, vectors := initData()
	for s, entry := range vents {
		if DEBUG {
			fmt.Printf("\n--- step %d ---\n", s)
		}
		vectors = computeVectorWithDiagonals(entry, vectors)
		if DEBUG {
			printVectors(vectors)
		}
	}
	totalOverlap := 0
	for _, vector := range vectors {
		for _, point := range vector {
			if point > 1 {
				totalOverlap++
			}
		}
	}
	fmt.Println("d5s2 |", totalOverlap)
}
