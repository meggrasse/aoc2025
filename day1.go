package main

import (
	"bufio"
	"os"
	"strconv"
)

func day1() (count int) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	position := 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		distance, _ := strconv.Atoi(line[1:])
		if line[0] == 'R' {
			position += distance

			for position > 99 {
				position -= 100
				count++
			}
		} else if line[0] == 'L' {
			if position == 0 {
				position = 100
			}
			position -= distance
			for position < 0 {
				position += 100
				count++
			}
			if position == 0 {
				count++
			}
		}
	}
	return
}
