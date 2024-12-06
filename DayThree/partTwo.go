package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("memory.txt")
	if err != nil {
		fmt.Println("Error reading memory.txt:", err)
		return
	}

	content := string(data)

	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	isEnabled := true
	total := 0

	instructions := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(content, -1)

	for _, instruction := range instructions {
		if instruction == "do()" {
			isEnabled = true
		} else if instruction == "don't()" {
			isEnabled = false
		} else if isEnabled {
			matches := mulRegex.FindStringSubmatch(instruction)
			if len(matches) == 3 {
				num1, _ := strconv.Atoi(matches[1])
				num2, _ := strconv.Atoi(matches[2])
				total += num1 * num2
			}
		}
	}

	fmt.Println("Total sum of enabled multiplications:", total)
}
