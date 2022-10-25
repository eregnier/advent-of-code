package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var DEBUG = os.Getenv("DEBUG") == "1"

func fileLines(file string) []string {
	content, err := ioutil.ReadFile("inputs/" + file)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(string(content), "\n")
}
