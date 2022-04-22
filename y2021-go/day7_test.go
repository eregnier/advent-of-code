package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestDay7Step1(t *testing.T) {
	var crabPositions []int
	minPosition := 0
	maxPosition := 0
	for _, val := range strings.Split(fileLines("d7")[0], ",") {
		value, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln("could not parse input file")
		}
		crabPositions = append(crabPositions, value)
		if value > maxPosition {
			maxPosition = value
		}
		if value < minPosition {
			minPosition = value
		}
	}
	if DEBUG {
		fmt.Println("Initial state :", crabPositions, minPosition, maxPosition)
	}
	minFuel := -1
	for x := minPosition; x < maxPosition; x++ {
		fuel := 0
		for _, crabPosition := range crabPositions {
			if x < crabPosition {
				fuel += crabPosition - x
				if DEBUG {
					fmt.Println("Move from ", crabPosition, "to", x, ":", crabPosition-x, "fuel")
				}
			} else {
				fuel += x - crabPosition
				if DEBUG {
					fmt.Println("Move from ", crabPosition, "to", x, ":", x-crabPosition, "fuel")
				}

			}
		}
		if DEBUG {
			fmt.Println("fuel for x", fuel, x)
		}
		if fuel < minFuel || minFuel == -1 {
			minFuel = fuel
		}

	}
	fmt.Println("d7s1 |", minFuel)

}

func TestDay7Step2(t *testing.T) {
	var crabPositions []int
	costs := make(map[int]int)
	minPosition := 0
	maxPosition := 0
	for _, val := range strings.Split(fileLines("d7")[0], ",") {
		value, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln("could not parse input file")
		}
		crabPositions = append(crabPositions, value)
		if value > maxPosition {
			maxPosition = value
		}
		if value < minPosition {
			minPosition = value
		}
	}
	if DEBUG {
		fmt.Println("Initial state :", crabPositions, minPosition, maxPosition)
	}
	//cache moves costs
	for x := minPosition; x < maxPosition+1; x++ {
		if x == 0 {
			costs[x] = 0
		} else {
			costs[x] = costs[x-1] + x
		}
	}
	minFuel := -1
	for x := minPosition; x < maxPosition; x++ {
		fuel := 0
		for _, crabPosition := range crabPositions {
			if x < crabPosition {
				cost := costs[crabPosition-x]
				fuel += cost
				if DEBUG {
					fmt.Println("Move from ", crabPosition, "to", x, ":", cost, "fuel")
				}
			} else {
				cost := costs[x-crabPosition]
				fuel += cost
				if DEBUG {
					fmt.Println("Move from ", crabPosition, "to", x, ":", cost, "fuel")
				}

			}
		}
		if DEBUG {
			fmt.Println("fuel for x", fuel, x)
		}
		if fuel < minFuel || minFuel == -1 {
			minFuel = fuel
		}

	}
	fmt.Println("d7s2 |", minFuel)

}
