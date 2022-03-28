package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func res2int(res []string) int64 {
	r, err := strconv.ParseInt(strings.Join(res, ""), 2, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return r
}

func TestDay3Step1(t *testing.T) {
	ones := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	zeroes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	gamma := []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
	epsilon := []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

	for _, line := range fileLines("d3") {
		for y, token := range strings.Split(line, "") {
			if token == "1" {
				ones[y]++
			}
			if token == "0" {
				zeroes[y]++
			}
		}
	}
	for x := 0; x < 12; x++ {
		if ones[x] > zeroes[x] {
			gamma[x] = "1"
			epsilon[x] = "0"
		} else {
			gamma[x] = "0"
			epsilon[x] = "1"
		}
	}

	g := res2int(gamma)
	e := res2int(epsilon)
	fmt.Printf("d3s1 | %d X %d => %d\n", g, e, g*e)
}

func TestDay3Step2(t *testing.T) {
	fmt.Println("TODO")
}
