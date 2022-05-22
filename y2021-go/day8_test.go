package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestDay8Step1(t *testing.T) {
	fieldCount := 0
	for _, line := range fileLines("d8") {
		for _, field := range strings.Fields(strings.Split(line, "|")[1]) {
			if len(field) == 7 || len(field) == 4 || len(field) == 2 || len(field) == 3 {
				fieldCount++
			}
		}
	}
	fmt.Println("d08s1 |", fieldCount)

}
func TestDay8Step2(t *testing.T) {
	digit := make(map[string]int)
	digit["abcdefg"] = 8
	digit["bcdef"] = 5
	digit["acdfg"] = 2
	digit["abcdf"] = 3
	digit["abd"] = 7
	digit["abcdef"] = 9
	digit["bcdefg"] = 6
	digit["abef"] = 4
	digit["abcdeg"] = 0
	digit["ab"] = 1

	//for _, line := range fileLines("d8") {
	//	entry := strings.Split(line, "|")
	//	head := entry[0]
	//	tail := entry[1]
	//	for _, field := range strings.Fields() {
	//	}
	//}
	//fmt.Println("acedgfb: 8 >", sortString("acedgfb"), " | ", word2weight("acedgfb", letterWeight))
	//fmt.Println("cdfbe: 5 >", sortString("cdfbe"), " | ", word2weight("cdfbe", letterWeight))
	//fmt.Println("gcdfa: 2 >", sortString("gcdfa"), " | ", word2weight("gcdfa", letterWeight))
	//fmt.Println("fbcad: 3 >", sortString("fbcad"), " | ", word2weight("fbcad", letterWeight))
	//fmt.Println("dab: 7 >", sortString("dab"), " | ", word2weight("dab", letterWeight))
	//fmt.Println("cefabd: 9 >", sortString("cefabd"), " | ", word2weight("cefabd", letterWeight))
	//fmt.Println("cdfgeb: 6 >", sortString("cdfgeb"), " | ", word2weight("cdfgeb", letterWeight))
	//fmt.Println("eafb: 4 >", sortString("eafb"), " | ", word2weight("eafb", letterWeight))
	//fmt.Println("cagedb: 0 >", sortString("cagedb"), " | ", word2weight("cagedb", letterWeight))
	//fmt.Println("ab: 1 >", sortString("ab"), " | ", word2weight("ab", letterWeight))

}

func word2weight(word string, letterWeight map[string]int) int {
	weight := 0
	for _, token := range word {
		weight += letterWeight[string(token)]
	}
	return weight
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
