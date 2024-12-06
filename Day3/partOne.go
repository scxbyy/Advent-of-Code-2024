package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("memory.txt")
	if err != nil {
		log.Fatal(err)
	}

	corruptedMemory := string(data)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	totalSum := 0

	matches := re.FindAllStringSubmatch(corruptedMemory, -1)

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		totalSum += x * y
	}

	fmt.Println("Total Sum of Multiplications:", totalSum)
}
