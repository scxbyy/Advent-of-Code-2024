package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkPattern(grid [][]rune, r1, c1, r2, c2, r3, c3 int) bool {
	pattern := string([]rune{grid[r1][c1], grid[r2][c2], grid[r3][c3]})
	if pattern == "MAS" || pattern == "SAM" {
		return true
	}

	pattern = string([]rune{grid[r1][c1], grid[r2][c2], grid[r3][c3]})
	if pattern == "SAM" || pattern == "MAS" {
		return true
	}

	return false
}

func isXMas(grid [][]rune, r, c int) bool {
	if r-1 < 0 || r+1 >= len(grid) || c-1 < 0 || c+1 >= len(grid[r]) {
		return false
	}

	if !checkPattern(grid, r-1, c-1, r, c, r+1, c+1) {
		return false
	}

	if !checkPattern(grid, r-1, c+1, r, c, r+1, c-1) {
		return false
	}

	return true
}

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	count := 0
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[r])-1; c++ {
			if isXMas(grid, r, c) {
				count++
			}
		}
	}

	fmt.Printf("The number of X-MAS patterns is: %d\n", count)
}
