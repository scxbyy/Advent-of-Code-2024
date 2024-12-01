package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Helper function to read the two lists from a file
func readLists(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into the two numbers
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		// Convert the strings to integers
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		// Append to the respective lists
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftList, rightList, nil
}

// Function to calculate the similarity score
func calculateSimilarityScore(leftList, rightList []int) int {
	similarityScore := 0

	// For each number in the left list, calculate its contribution to the similarity score
	for _, left := range leftList {
		// Count how many times `left` appears in the right list
		count := 0
		for _, right := range rightList {
			if left == right {
				count++
			}
		}

		// Multiply `left` by the count and add to the total similarity score
		similarityScore += left * count
	}

	return similarityScore
}

func main() {
	// Read lists from the file
	leftList, rightList, err := readLists("lists.txt")
	if err != nil {
		fmt.Println("Error reading lists:", err)
		return
	}

	// Calculate the total similarity score
	similarityScore := calculateSimilarityScore(leftList, rightList)

	// Output the result
	fmt.Println("Total similarity score:", similarityScore)
}
