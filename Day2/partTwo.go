package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if isIncreasing && diff < 0 || !isIncreasing && diff > 0 {
			return false
		}
	}

	return true
}

func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		modifiedReport := append([]int{}, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)
		if isSafe(modifiedReport) {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("reports.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		numStrings := strings.Fields(line)
		var report []int

		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error parsing number:", numStr)
				return
			}
			report = append(report, num)
		}

		if isSafeWithDampener(report) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Number of safe reports:", safeCount)
}
