package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestDay10Step1(t *testing.T) {

	totalError := 0
	for l, line := range fileLines("d10") {
		var stack []rune
		for x := 0; x < len(line); x++ {
			if line[x] == '(' {
				stack = append(stack, ')')
			}
			if line[x] == '<' {
				stack = append(stack, '>')
			}
			if line[x] == '{' {
				stack = append(stack, '}')
			}
			if line[x] == '[' {
				stack = append(stack, ']')
			}
			if line[x] == ')' {
				if stack[len(stack)-1] != ')' {
					if DEBUG {
						fmt.Printf("%s => match error on token %c at line %d::%d", line, line[x], l, x)
					}
					totalError += 3
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if line[x] == '}' {
				if stack[len(stack)-1] != '}' {
					if DEBUG {
						fmt.Printf("%s => match error on token %c at line %d::%d", line, line[x], l, x)
					}
					totalError += 1197
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if line[x] == '>' {
				if stack[len(stack)-1] != '>' {
					if DEBUG {
						fmt.Printf("%s => match error on token %c at line %d::%d", line, line[x], l, x)
					}
					totalError += 25137
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if line[x] == ']' {
				if stack[len(stack)-1] != ']' {
					if DEBUG {
						fmt.Printf("%s => match error on token %c at line %d::%d", line, line[x], l, x)
					}
					totalError += 57
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		if DEBUG {
			fmt.Println("")
		}
	}
	fmt.Println("d10s1 |", totalError)

}

func TestDay10Step2(t *testing.T) {

	var scores []int
	for l, line := range fileLines("d10") {
		var stack []rune
		error := false
		for x := 0; x < len(line); x++ {
			if line[x] == '(' {
				stack = append(stack, ')')
			}
			if line[x] == '<' {
				stack = append(stack, '>')
			}
			if line[x] == '{' {
				stack = append(stack, '}')
			}
			if line[x] == '[' {
				stack = append(stack, ']')
			}
			if line[x] == ')' {
				if stack[len(stack)-1] != ')' {
					if DEBUG {
						fmt.Printf("%s => match error on token %c at line %d::%d", line, line[x], l, x)
					}
					error = true
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if line[x] == '}' {
				if stack[len(stack)-1] != '}' {
					if DEBUG {
						fmt.Printf("%s => match error on token %c at line %d::%d", line, line[x], l, x)
					}
					error = true
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if line[x] == '>' {
				if stack[len(stack)-1] != '>' {
					if DEBUG {
						fmt.Printf("%s => match error on token %c at line %d::%d", line, line[x], l, x)
					}
					error = true

					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			if line[x] == ']' {
				if stack[len(stack)-1] != ']' {
					if DEBUG {
						fmt.Printf("%s => match error on token %c at line %d::%d", line, line[x], l, x)
					}
					error = true

					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		score := 0
		if !error && len(stack) > 0 {
			if DEBUG {
				fmt.Println(string(stack))
			}
			for x := len(stack) - 1; x >= 0; x-- {
				token := stack[x]
				if DEBUG {
					fmt.Print(string(token))
				}
				score *= 5
				if token == ')' {
					score += 1
				}
				if token == ']' {
					score += 2
				}
				if token == '}' {
					score += 3
				}
				if token == '>' {
					score += 4
				}
			}
			scores = append(scores, score)
			if DEBUG {
				fmt.Println(" > score ", score)
			}
		}

	}
	sort.Ints(scores[:])
	if DEBUG {
		fmt.Println(scores)
	}
	fmt.Println("d10s2 |", scores[len(scores)/2])

}
