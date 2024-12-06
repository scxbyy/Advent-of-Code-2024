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

func main() {
	rules, updates := parseInput("manuals.txt")
	totalMiddle := 0
	for _, update := range updates {
		if isValidUpdate(rules, update) {
			middle := update[len(update)/2]
			totalMiddle += middle
		}
	}
	fmt.Println(totalMiddle)
}
