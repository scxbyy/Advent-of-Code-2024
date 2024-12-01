package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftList, rightList, nil
}

func calculateDistance(leftList, rightList []int) int {
	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}

	return totalDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	leftList, rightList, err := readLists("lists.txt")
	if err != nil {
		fmt.Println("Error reading lists:", err)
		return
	}

	totalDistance := calculateDistance(leftList, rightList)
	fmt.Println("Total distance:", totalDistance)
}
