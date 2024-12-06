package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([][]string, [][]int) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rules [][]string
	var updates [][]int
	isRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRules = false
			continue
		}
		if isRules {
			rules = append(rules, strings.Split(line, "|"))
		} else {
			update := []int{}
			for _, num := range strings.Split(line, ",") {
				val, _ := strconv.Atoi(num)
				update = append(update, val)
			}
			updates = append(updates, update)
		}
	}
	return rules, updates
}

func isValidUpdate(rules [][]string, update []int) bool {
	pageIndex := map[int]int{}
	for i, page := range update {
		pageIndex[page] = i
	}
	for _, rule := range rules {
		x, _ := strconv.Atoi(rule[0])
		y, _ := strconv.Atoi(rule[1])
		xIdx, xExists := pageIndex[x]
		yIdx, yExists := pageIndex[y]
		if xExists && yExists && xIdx > yIdx {
			return false
		}
	}
	return true
}

func reorderUpdate(rules [][]string, update []int) []int {
	pageIndex := map[int]int{}
	for i, page := range update {
		pageIndex[page] = i
	}

	// Build dependency graph
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, rule := range rules {
		x, _ := strconv.Atoi(rule[0])
		y, _ := strconv.Atoi(rule[1])
		if _, xExists := pageIndex[x]; xExists {
			if _, yExists := pageIndex[y]; yExists {
				graph[x] = append(graph[x], y)
				inDegree[y]++
			}
		}
	}

	// Topological sort using Kahn's Algorithm
	var sorted []int
	queue := []int{}
	for _, page := range update {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		sorted = append(sorted, curr)
		for _, neighbor := range graph[curr] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func main() {
	rules, updates := parseInput("manuals.txt")
	totalMiddle := 0
	for _, update := range updates {
		if !isValidUpdate(rules, update) {
			ordered := reorderUpdate(rules, update)
			middle := ordered[len(ordered)/2]
			totalMiddle += middle
		}
	}
	fmt.Println(totalMiddle)
}
