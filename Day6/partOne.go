package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	// Open the input file
	file, err := os.Open("maps.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the map into a 2D grid
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

	// Find the starting position and initial direction
	var guard Point
	directions := []Point{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}
	var dirIndex int // Initial direction is up

	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' || cell == '>' || cell == 'v' || cell == '<' {
				guard = Point{x, y}
				switch cell {
				case '^':
					dirIndex = 0
				case '>':
					dirIndex = 1
				case 'v':
					dirIndex = 2
				case '<':
					dirIndex = 3
				}
				grid[y][x] = '.' // Clear the starting position
			}
		}
	}

	// Track visited positions
	visited := map[Point]bool{}
	visited[guard] = true

	// Simulate the guard's movement
	for {
		next := Point{guard.x + directions[dirIndex].x, guard.y + directions[dirIndex].y}

		// Check if the guard is out of bounds
		if next.y < 0 || next.y >= len(grid) || next.x < 0 || next.x >= len(grid[0]) {
			break
		}

		// Check if the next position is blocked
		if grid[next.y][next.x] == '#' {
			// Turn right (90 degrees clockwise)
			dirIndex = (dirIndex + 1) % 4
		} else {
			// Move forward
			guard = next
			visited[guard] = true
		}
	}

	// Count the number of distinct visited positions
	fmt.Println("Distinct positions visited:", len(visited))
}
