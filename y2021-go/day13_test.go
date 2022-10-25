package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

type Fold struct {
	dimension string
	value     int
}

func TestDay13Step1(t *testing.T) {
	points, maxX, maxY, folds := readData()
	points = mergeHorizontal(folds[0].value, maxX, maxY, points)
	fmt.Println("d11s1 |", len(points))
}

func TestDay13Step2(t *testing.T) {
	points, maxX, maxY, folds := readData()
	for _, fold := range folds {
		if fold.dimension == "x" {
			points = mergeHorizontal(fold.value, maxX, maxY, points)
			maxX = fold.value
		}
		if fold.dimension == "y" {
			points = mergeVertical(fold.value, maxX, maxY, points)
			maxY = fold.value
		}
	}
	fmt.Println("d11s1 |")
	display(points, maxX, maxY)
}

func mergeVertical(index int, maxX int, maxY int, allPoints map[string]bool) map[string]bool {
	var resultPoints = make(map[string]bool)
	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			if y < index {
				pointTop := fmt.Sprintf("%d,%d", x, y)
				if allPoints[pointTop] {
					resultPoints[pointTop] = true
				}
			}
			if y > index {
				pointBottom := fmt.Sprintf("%d,%d", x, y)
				pointBottomMove := fmt.Sprintf("%d,%d", x, int(math.Abs(float64(y-2*index))))
				if allPoints[pointBottom] {
					resultPoints[pointBottomMove] = true
				}
			}
		}
	}
	return resultPoints
}

func mergeHorizontal(index int, maxX int, maxY int, allPoints map[string]bool) map[string]bool {
	var resultPoints = make(map[string]bool)
	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			if x < index {
				leftPoint := fmt.Sprintf("%d,%d", x, y)
				if allPoints[leftPoint] {
					resultPoints[leftPoint] = true
				}
			}
			if x > index {
				pointRight := fmt.Sprintf("%d,%d", x, y)
				pointRightMove := fmt.Sprintf("%d,%d", int(math.Abs(float64(x-2*index))), y)
				if allPoints[pointRight] {
					resultPoints[pointRightMove] = true
				}
			}
		}
	}
	return resultPoints
}

func display(points map[string]bool, maxX int, maxY int) {
	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX+1; x++ {
			point := fmt.Sprintf("%d,%d", x, y)
			if points[point] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func readData() (map[string]bool, int, int, []Fold) {
	var folds = make([]Fold, 0)
	var points = make(map[string]bool)
	maxX := 0
	maxY := 0
	for _, line := range fileLines("d13") {
		if strings.TrimSpace(line) == "" {
			continue
		} else if strings.HasPrefix(line, "fold") {
			foldInfo := strings.Split(strings.ReplaceAll(line, "fold along ", ""), "=")
			value, _ := strconv.Atoi(foldInfo[1])
			folds = append(folds, Fold{dimension: foldInfo[0], value: value})
		} else {
			points[line] = true
			coords := strings.Split(line, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	return points, maxX, maxY, folds
}
