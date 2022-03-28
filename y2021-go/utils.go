package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func fileLines(file string) []string{
	content, err := ioutil.ReadFile("inputs/" + file)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(string(content), "\n")
}
