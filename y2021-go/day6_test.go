package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

//var DEBUG = false

func TestDay6Step1(t *testing.T) {
	days := 80
	var fishes []int
	for _, fish := range strings.Split(fileLines("d6")[0], ",") {
		value, err := strconv.Atoi(fish)
		if err != nil {
			log.Fatalln("could not parse file")
		}
		fishes = append(fishes, value)
	}
	if DEBUG {
		fmt.Println("Initial state :", fishes)
	}

	for d := 0; d < days; d++ {
		appendCount := 0
		for x := 0; x < len(fishes); x++ {
			if fishes[x] == 0 {
				appendCount++
				fishes[x] = 6
			} else {
				fishes[x]--
			}
		}
		for x := 0; x < appendCount; x++ {
			fishes = append(fishes, 8)
		}
		if DEBUG {
			fmt.Println("After", d, "days: ", fishes)
		}
	}
	fmt.Println("d6s1 |", len(fishes))

}

func TestDay6Step2(t *testing.T) {
	days := 256
	fishes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, fish := range strings.Split(fileLines("d6")[0], ",") {
		value, err := strconv.Atoi(fish)
		if err != nil {
			log.Fatalln("could not parse file")
		}
		fishes[value]++
	}
	if DEBUG {
		fmt.Println("Initial state :", fishes)
	}

	resetFishes := 0
	for d := 0; d < days; d++ {
		newFishes := fishes[0]
		for x := 0; x < 8; x++ {
			fishes[x] = fishes[x+1]
		}
		fishes[6] += resetFishes
		resetFishes = fishes[0]
		fishes[8] = newFishes
		if DEBUG {
			fmt.Println("After", d, "days: ", fishes)
			//fmt.Println("DEBUG", "        ", debugger(d))
		}
	}
	sum := 0
	for _, fish := range fishes {
		sum += fish
	}
	fmt.Println("d6s2 |", sum)

}

func debugger(day int) []int {
	data :=
		[]string{
			"2,3,2,0,1",
			"1,2,1,6,0,8",
			"0,1,0,5,6,7,8",
			"6,0,6,4,5,6,7,8,8",
			"5,6,5,3,4,5,6,7,7,8",
			"4,5,4,2,3,4,5,6,6,7",
			"3,4,3,1,2,3,4,5,5,6",
			"2,3,2,0,1,2,3,4,4,5",
			"1,2,1,6,0,1,2,3,3,4,8",
			"0,1,0,5,6,0,1,2,2,3,7,8",
			"6,0,6,4,5,6,0,1,1,2,6,7,8,8,8",
			"5,6,5,3,4,5,6,0,0,1,5,6,7,7,7,8,8",
			"4,5,4,2,3,4,5,6,6,0,4,5,6,6,6,7,7,8,8",
			"3,4,3,1,2,3,4,5,5,6,3,4,5,5,5,6,6,7,7,8",
			"2,3,2,0,1,2,3,4,4,5,2,3,4,4,4,5,5,6,6,7",
			"1,2,1,6,0,1,2,3,3,4,1,2,3,3,3,4,4,5,5,6,8",
			"0,1,0,5,6,0,1,2,2,3,0,1,2,2,2,3,3,4,4,5,7,8",
			"6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8",
		}
	fishes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, v := range strings.Split(data[day], ",") {
		i, _ := strconv.Atoi(v)
		fishes[i]++
	}
	return fishes
}
