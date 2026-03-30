package main

import (
	"bufio"
	"bytes"
	"math"
	"os"
	"strconv"
)

func day2() (sum int) {
	file, _ := os.Open("day2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		i := bytes.Index(data, []byte(","))
		if i == -1 {
			if atEOF {
				return 0, data, bufio.ErrFinalToken
			}
			return 0, nil, nil
		} else {
			return i + 1, data[0:i], nil
		}
		// no error handling for aoc
	}

	scanner.Split(split)

	for scanner.Scan() {
		b := scanner.Bytes()
		l, u, _ := bytes.Cut(b, []byte("-"))
		upper, _ := strconv.Atoi(string(u))
		lower, _ := strconv.Atoi(string(l))
		for i := lower; i <= upper; i++ {
			pow := int(math.Ceil(math.Log10(float64(i)))) / 2
			base := int(math.Pow10(pow))
			lhs := i - i/base        // intentionally using integer division
			rhs := (i / base) * base // intentionally using integer division
			if lhs == rhs {
				sum += i
			}
		}

	}

	return
}
