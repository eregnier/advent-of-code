package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay14Step1(t *testing.T) {
	transformations, polymer := readData()
	for x := 0; x < 10; x++ {
		polymer = transform(polymer, transformations)
		//fmt.Println(polymer)
	}
	min, max := getMinMaxPolymerElement(polymer)
	fmt.Println("d14s1 |", max-min)

}

func TestDay14Step2(t *testing.T) {
	//https://www.reddit.com/r/adventofcode/comments/rfzq6f/comment/hoib78w/?utm_source=share&utm_medium=web2x&context=3
	transformations, polymer := readData()
	fmt.Println(transformations)
	polymerMap := polymerToMap(polymer)
	fmt.Println(polymerMap)
	for x := 0; x < 40; x++ {
		polymerMap = transform2(polymerMap, transformations)
	}
	min, max := getMinMaxPolymerElement2(polymerMap)
	fmt.Println("d14s2 |", max-min)

}

func getMinMaxPolymerElement2(polymerMap map[string]int) (int, int) {
	elementCounts := make(map[string]int)
	for item, count := range polymerMap {
		head, tail := string(item[0]), string(item[1])
		elementCounts[head] += count
		elementCounts[tail] += count
	}
	min := 0
	max := 0
	for _, count := range elementCounts {
		if min == 0 || count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	return min, max

}

func transform2(polymerMap map[string]int, transformations map[string]string) map[string]int {
	newPolymerMap := make(map[string]int)
	for key, value := range polymerMap {
		head, tail := string(key[0]), string(key[1])
		newItem := transformations[head+tail]
		newPolymerMap[head+newItem] += value
		newPolymerMap[newItem+tail] += value
	}
	return newPolymerMap
}

func polymerToMap(polymer string) map[string]int {
	polymerMap := make(map[string]int)
	for x := 0; x < len(polymer)-1; x++ {
		head := string(polymer[x])
		tail := string(polymer[x+1])
		polymerMap[head+tail] += 1
	}
	return polymerMap
}

func getMinMaxPolymerElement(polymer string) (int, int) {
	elementCounts := make(map[string]int)
	for x := 0; x < len(polymer); x++ {
		elementCounts[string(polymer[x])] += 1
	}
	min := 0
	max := 0
	for _, count := range elementCounts {
		if min == 0 || count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	return min, max
}

func transform(polymer string, transformations map[string]string) string {
	newPolymer := ""
	var lastTail string
	for x := 0; x < len(polymer)-1; x++ {
		head := string(polymer[x])
		tail := string(polymer[x+1])
		newPolymer += head + transformations[head+tail]
		lastTail = tail
	}
	return newPolymer + lastTail
}

func readData() (map[string]string, string) {
	var axiom string
	var transformations = make(map[string]string)
	for x, line := range fileLines("d14") {
		if x == 0 {
			axiom = strings.TrimSpace(line)
		} else if line != "" {
			tokens := strings.Split(line, " -> ")
			transformations[tokens[0]] = tokens[1]
		}
	}
	return transformations, axiom

}
