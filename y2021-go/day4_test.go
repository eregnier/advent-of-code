package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

type Cell struct {
	val   int
	found bool
}

func markNumber(number int, board [][]Cell) [][]Cell {
	for x, lineBoard := range board {
		for y, cell := range lineBoard {
			if cell.val == number {
				board[x][y].found = true
			}
		}
	}
	return board
}

func parseFile() ([]int, [][][]Cell) {
	var runs []int
	var boards [][][]Cell
	var board [][]Cell
	for x, line := range fileLines("d4") {
		var lineBoard []Cell
		if x == 0 {
			for _, run := range strings.Split(line, ",") {
				pick, err := strconv.Atoi(run)
				if err != nil {
					log.Fatalln(err)
				}
				runs = append(runs, pick)
			}
		} else {
			if line == "" {
				if len(board) > 0 {
					boards = append(boards, board)
					board = make([][]Cell, 0)
				}
			} else {
				for _, token := range strings.Fields(line) {
					value, err := strconv.Atoi(strings.TrimSpace(token))
					if err != nil {
						log.Fatalln(err)
					}
					lineBoard = append(lineBoard, Cell{value, false})
				}
				board = append(board, lineBoard)
				lineBoard = make([]Cell, 0)

			}
		}
	}
	return runs, boards
}

func checkWin(board [][]Cell) (int, bool) {
	total := 0
	var foundByColumn = []int{0, 0, 0, 0, 0}
	foundLine := false
	foundColumn := false
	for _, lineBoard := range board {
		found := 0
		for y, cell := range lineBoard {
			if cell.found {
				found++
				foundByColumn[y]++
			}
		}
		if found == 5 {
			foundLine = true
		}
	}

	for _, column := range foundByColumn {
		if column == 5 {
			foundColumn = true
		}
	}
	if foundLine || foundColumn {
		for _, line := range board {
			for _, cell := range line {
				if !cell.found {
					total += cell.val
				}
			}
		}
		return total, true
	}
	return total, false
}

func TestDay4Step1(t *testing.T) {
	runs, boards := parseFile()
	for _, pick := range runs {
		for _, board := range boards {
			markNumber(pick, board)
			value, found := checkWin(board)
			if found {
				fmt.Println("d04s1 |", value*pick)
				return
			}
		}

	}
}

func TestDay4Step2(t *testing.T) {
	runs, boards := parseFile()
	lastWinValue := 0
	lastPickValue := 0
	pickedBoards := make(map[int]bool)
	for _, pick := range runs {
		for x, board := range boards {
			if !pickedBoards[x] {
				markNumber(pick, board)
				value, found := checkWin(board)
				if found {
					pickedBoards[x] = true
					lastWinValue = value
					lastPickValue = pick
				}
			}
		}
	}
	fmt.Println("d04s2 |", lastWinValue*lastPickValue)
}
