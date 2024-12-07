package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to recursively evaluate possible results with operators
func evaluate(nums []int, idx int, current int, target int) bool {
	if idx == len(nums) {
		return current == target
	}

	// Try addition
	if evaluate(nums, idx+1, current+nums[idx], target) {
		return true
	}

	// Try multiplication
	if evaluate(nums, idx+1, current*nums[idx], target) {
		return true
	}

	// Try concatenation
	concatValue := concat(current, nums[idx])
	if evaluate(nums, idx+1, concatValue, target) {
		return true
	}

	return false
}

// Helper function to concatenate two integers
func concat(a, b int) int {
	return atoi(fmt.Sprintf("%d%d", a, b))
}

// Helper function to safely convert a string to an integer
func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	file, err := os.Open("calibration.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalCalibration int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		// Parse the target value
		target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Invalid target value:", parts[0])
			continue
		}

		// Parse the numbers
		numStrs := strings.Fields(parts[1])
		nums := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			nums[i], err = strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Invalid number:", numStr)
				continue
			}
		}

		// Check if the equation can be satisfied
		if evaluate(nums, 1, nums[0], target) {
			totalCalibration += target
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total Calibration Result:", totalCalibration)
}
