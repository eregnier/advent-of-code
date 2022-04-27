package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type Point struct {
	x int
	y int
}

func d9InitData() [][]int {
	var cave [][]int
	for _, line := range fileLines("d9") {
		var caveLine []int
		for _, point := range strings.Split(line, "") {
			value, err := strconv.Atoi(point)
			if err != nil {
				log.Fatalln("could not parse file")
			}
			caveLine = append(caveLine, value)
		}
		cave = append(cave, caveLine)
	}
	if DEBUG {
		for _, line := range cave {
			fmt.Println(line)
		}
	}
	return cave
}

func computeLowLevel(cave [][]int, h int, w int) []int {
	var lowPoints []int
	for y, line := range cave {
		for x, point := range line {
			isLowPoint := true
			//left
			if x > 0 && cave[y][x-1] <= point {
				isLowPoint = false
			}
			//top
			if y > 0 && cave[y-1][x] <= point {
				isLowPoint = false
			}
			//right
			if x+1 < w && cave[y][x+1] <= point {
				isLowPoint = false
			}
			//bottom
			if y+1 < h && cave[y+1][x] <= point {
				isLowPoint = false
			}
			if isLowPoint {
				lowPoints = append(lowPoints, point)
			}
		}

	}
	if DEBUG {
		fmt.Println("low points", lowPoints)
	}
	return lowPoints
}

func TestDay9Step1(t *testing.T) {
	cave := d9InitData()
	h, w := len(cave), len(cave[0])
	if DEBUG {
		fmt.Println("h|w", h, w)
	}
	lowPoints := computeLowLevel(cave, h, w)
	riskLevel := 0
	for _, lowPoint := range lowPoints {
		riskLevel += 1 + lowPoint
	}
	fmt.Println("d9s1 |", riskLevel)
}

func lowLevelCoords(cave [][]int, h int, w int) []Point {
	var lowPoints []Point
	for y, line := range cave {
		for x, point := range line {
			isLowPoint := true
			//left
			if x > 0 && cave[y][x-1] <= point {
				isLowPoint = false
			}
			//top
			if y > 0 && cave[y-1][x] <= point {
				isLowPoint = false
			}
			//right
			if x+1 < w && cave[y][x+1] <= point {
				isLowPoint = false
			}
			//bottom
			if y+1 < h && cave[y+1][x] <= point {
				isLowPoint = false
			}
			if isLowPoint {
				lowPoints = append(lowPoints, Point{x, y})
			}
		}

	}
	if DEBUG {
		fmt.Println("low points", lowPoints)
	}
	return lowPoints
}

func findBasin(point Point, affectedPoints *map[string]bool, cave [][]int, w int, h int) []Point {
	pointKey := fmt.Sprintf("%d-%d", point.x, point.y)
	var points []Point
	if cave[point.y][point.x] < 9 && !(*affectedPoints)[pointKey] {
		(*affectedPoints)[pointKey] = true
		points = append(points, point)
		//left
		if point.x > 0 {
			neighbors := findBasin(Point{point.x - 1, point.y}, affectedPoints, cave, w, h)
			points = append(points, neighbors...)
		}
		//top
		if point.y > 0 {
			neighbors := findBasin(Point{point.x, point.y - 1}, affectedPoints, cave, w, h)
			points = append(points, neighbors...)
		}
		//right
		if point.x+1 < w {
			neighbors := findBasin(Point{point.x + 1, point.y}, affectedPoints, cave, w, h)
			points = append(points, neighbors...)
		}
		////bottom
		if point.y+1 < h {
			neighbors := findBasin(Point{point.x, point.y + 1}, affectedPoints, cave, w, h)
			points = append(points, neighbors...)
		}
	}
	return points
}

func TestDay9Step2(t *testing.T) {
	cave := d9InitData()
	h, w := len(cave), len(cave[0])
	lowPoints := lowLevelCoords(cave, h, w)
	affectedPoints := make(map[string]bool)
	var basinsSize []int
	for _, lowPoint := range lowPoints {
		basinsSize = append(basinsSize, len(findBasin(lowPoint, &affectedPoints, cave, w, h)))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinsSize)))
	fmt.Println("d9s2 |", basinsSize[0] * basinsSize[1] * basinsSize[2] )

}
