package main

import (
	"bufio"
	"fmt"
	"os"
)

const targetWord = "XMAS"

var directions = []struct {
	dx, dy int
}{
	{1, 0},
	{-1, 0},
	{0, 1},
	{1, 1},
	{-1, 1},
	{1, -1},
	{-1, -1},
}

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, dir := range directions {
				if found := searchWord(grid, row, col, dir.dx, dir.dy); found {
					count++
				}
			}
		}
	}

	fmt.Println("Total occurrences of 'XMAS':", count)
}

func searchWord(grid []string, row, col, dx, dy int) bool {
	for i := 0; i < len(targetWord); i++ {
		newRow := row + i*dy
		newCol := col + i*dx

		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[newRow]) || grid[newRow][newCol] != targetWord[i] {
			return false
		}
	}
	return true
}
