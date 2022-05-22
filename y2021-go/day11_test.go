package main

import (
	"fmt"
	"testing"
)

func TestDay11Step1(t *testing.T) {

	for _, line := range fileLines("d11") {
		if DEBUG {
			fmt.Println(line)
		}
	}
	fmt.Println("d11s1 |", 1)
}
