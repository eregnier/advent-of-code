package main

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

type Node struct {
	children []Node
	Value    string
}

func inArray(token string, items []string) bool {
	for _, item := range items {
		if item == token {
			return true
		}
	}
	return false
}

func loadGraph() map[string][]string {
	tree := make(map[string][]string)
	for _, line := range fileLines("d12") {
		fmt.Println(line)
		nodeLink := strings.Split(line, "-")
		//if !inArray(nodeLink[1], tree[nodeLink[0]]) {
		if nodeLink[1] != "start" {
			tree[nodeLink[0]] = append(tree[nodeLink[0]], nodeLink[1])
		}
		//}
		//if !inArray(nodeLink[0], tree[nodeLink[1]]) {
		if nodeLink[0] != "start" {
			tree[nodeLink[1]] = append(tree[nodeLink[1]], nodeLink[0])
		}
		//}
	}
	fmt.Println(tree)
	return tree
}

//
//func explore(node map[string][]string, children []string, path []string) int {
//	count := 0
//	for _, child := range children {
//		if child == "end" {
//			count += 1
//		} else {
//			if continueExplore(path, child) {
//
//			}
//			if unicode.IsUpper(rune(child[0])) {
//				fmt.Println("upper")
//			}
//		}
//	}
//	return count
//}

//func contains(){
//
//	for _, item := range items {
//	 if item == contain {
//
//		if item == contain {
//			return true
//		}
//	}
//	return false
//	}
//func continueExplore(items []string, contain string) bool {
//if unicode.IsUpper(rune(contain[0])){
//	return true
//}else {
//
//
//}

func walk(tree map[string][]string, step string, stack []string) []string {
	//fmt.Println("walking", step, stack)
	if step == "end" || step == "start" {
		return stack
	} else {
		stack = append(stack, step)
		isLower := unicode.IsLower([]rune(step)[0])
		//fmt.Println("is Lower", step, isLower)
		if isLower && inArray(step, stack) {
			//fmt.Println(step, "in", stack, " > return")
			return stack
		}
		for _, path := range tree[step] {
			fmt.Println("PATH", stack, walk(tree, path, stack))
		}
	}
	return stack
}

func TestDay12Step1(t *testing.T) {
	tree := loadGraph()
	fmt.Println(tree)
	for _, path := range tree["start"] {
		walk(tree, path, []string{})
	}

	//pathCount := explore(tree, tree["start"])
	//fmt.Println("pathCount", pathCount)
	//fmt.Println(nodeLinks)
	//root := Node{Value: "start"}
	//for {
	//	change := false
	//	for _, nodeLink := range nodeLinks {
	//		treeChange := false
	//		root, treeChange = appendNode(root, nodeLink)
	//		if treeChange {
	//			change = true
	//		}
	//	}
	//	if !change {
	//		break
	//	}
	//}
	//computePaths(root)
}

//func computePaths(node Node) {
//	fmt.Println(node)
//}
