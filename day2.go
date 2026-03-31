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
	outer:
		for i := lower; i <= upper; i++ {
			max_pow := int(math.Log10(float64(i))) + 1
			for j := 1; j <= max_pow/2; j++ {
				for k := j; k <= max_pow-j; k += j {
					lhs := i % int(math.Pow10(k)) / int(math.Pow10(k-j))
					rhs := i % int(math.Pow10(k+j)) / int(math.Pow10(k))
					if lhs != rhs {
						break
					}
					if k == max_pow-j {
						sum += i
						continue outer
					}
				}
			}
		}

	}

	return
}
