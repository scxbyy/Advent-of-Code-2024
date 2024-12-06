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

	// Function to simulate guard movement and check for a loop
	checkForLoop := func(tempGrid [][]rune, start Point, startDir int) bool {
		visited := map[Point]int{}
		guard := start
		dirIndex := startDir

		for {
			// If we've seen this position and direction before, there's a loop
			state := Point{guard.x*4 + dirIndex, guard.y*4 + dirIndex}
			if visited[state] > 0 {
				return true
			}
			visited[state]++

			// Calculate the next position
			next := Point{guard.x + directions[dirIndex].x, guard.y + directions[dirIndex].y}

			// Check if the guard is out of bounds
			if next.y < 0 || next.y >= len(tempGrid) || next.x < 0 || next.x >= len(tempGrid[0]) {
				break
			}

			// Check if the next position is blocked
			if tempGrid[next.y][next.x] == '#' {
				// Turn right (90 degrees clockwise)
				dirIndex = (dirIndex + 1) % 4
			} else {
				// Move forward
				guard = next
			}
		}

		return false
	}

	// Find all valid positions for new obstructions
	validPositions := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			// Skip positions that are already blocked or the guard's starting position
			if grid[y][x] != '.' || (x == guard.x && y == guard.y) {
				continue
			}

			// Temporarily place an obstruction
			grid[y][x] = '#'

			// Check if this causes the guard to get stuck in a loop
			if checkForLoop(grid, guard, dirIndex) {
				validPositions++
			}

			// Remove the obstruction
			grid[y][x] = '.'
		}
	}

	// Output the number of valid positions
	fmt.Println("Valid positions for obstruction:", validPositions)
}
