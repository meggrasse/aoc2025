package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	file, open_err := os.Open("input.txt")
	if open_err != nil {
		fmt.Println("Error opening file: %v", open_err)
	  }
	defer file.Close()

	position := 50
	var count int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		distance, _ := strconv.Atoi(line[1:])
		if line[0] == 'R' {
			distance *= 1
		} else if line[0] == 'L' {
			distance *= -1
		}
		position += distance
		position = position % 100
		if position == 0 {
			count++
		}
	}

	fmt.Println("Password:", count)
}