package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

type Octopus struct {
	Energy     int
	HasFlashed bool
}

func loadFloor() [][]Octopus {
	var floor [][]Octopus
	for _, line := range fileLines("d11") {
		var octopusLine []Octopus
		for i := 0; i < len(line); i++ {
			octopusEnergy, err := strconv.Atoi(string(line[i]))
			if err != nil {
				log.Fatalln(err)
			}
			octopusLine = append(octopusLine, Octopus{octopusEnergy, false})
		}
		floor = append(floor, octopusLine)
	}
	return floor

}

func TestDay11Step1(t *testing.T) {
	floor := loadFloor()
	STEPS := 99
	flashes := 0
	printBoard(floor, "init")

	for s := 0; s < STEPS; s++ {
		if DEBUG {
			fmt.Println("\n ---> new step", s)
		}
		for y, line := range floor {
			for x := range line {
				if floor[y][x].HasFlashed {
					floor[y][x].Energy = 0
					floor[y][x].HasFlashed = false
				} else {
					floor[y][x].Energy++
				}
			}
		}

		printBoard(floor, "after +1")

		for {
			change := false
			for y, line := range floor {
				for x := range line {
					if floor[y][x].Energy >= 9 && !floor[y][x].HasFlashed {
						flashes++
						floor[y][x].HasFlashed = true
						change = true
						floor = flash(floor, x, y)
					}
				}
			}
			if !change {
				break
			}
		}

	}
	fmt.Println("d11s1 |", flashes)
}

func TestDay11Step2(t *testing.T) {
	floor := loadFloor()
	step := 1
	syncStep := 0
	for {
		for y, line := range floor {
			for x := range line {
				if floor[y][x].HasFlashed {
					floor[y][x].Energy = 0
					floor[y][x].HasFlashed = false
				} else {
					floor[y][x].Energy++
				}
			}
		}

		printBoard(floor, fmt.Sprintf("step %d", step))

		if checkSync(floor) {
			syncStep = step
			break
		}

		for {
			change := false
			for y, line := range floor {
				for x := range line {
					if floor[y][x].Energy >= 9 && !floor[y][x].HasFlashed {
						floor[y][x].HasFlashed = true
						change = true
						floor = flash(floor, x, y)
					}
				}
			}
			if !change {
				break
			}
		}
		step++
	}
	fmt.Println("d11s2 |", syncStep)
}

func checkSync(floor [][]Octopus) bool {
	for y, line := range floor {
		for x := range line {
			if floor[y][x].Energy != 0 {
				return false
			}
		}
	}
	return true
}

func flash(floor [][]Octopus, x int, y int) [][]Octopus {
	lx := len(floor[0]) - 1
	ly := len(floor) - 1
	if y > 0 {
		floor[y-1][x].Energy++
	}
	//fmt.Println("y,x,lx", y, x, lx)
	if y > 0 && x < lx {
		floor[y-1][x+1].Energy++
	}
	if x < lx {
		floor[y][x+1].Energy++
	}
	if y < ly && x < lx {
		floor[y+1][x+1].Energy++
	}
	if y < ly {
		floor[y+1][x].Energy++
	}
	if y < ly && x > 0 {
		floor[y+1][x-1].Energy++
	}
	if x > 0 {
		floor[y][x-1].Energy++
	}
	if y > 0 && x > 0 {
		floor[y-1][x-1].Energy++
	}
	return floor
}

func printBoard(floor [][]Octopus, title string) {
	if DEBUG {
		fmt.Println("\n----------\n", title, "\n---vvvv----")
		for y := range floor {
			for x := range floor[y] {
				fmt.Print(floor[y][x].Energy)
			}
			fmt.Println("")
		}
	}
}
