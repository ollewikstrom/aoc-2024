package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var safeRoutesCount int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue // Skip blank lines
		}
		fmt.Println(line)
		vals := strings.Split(line, " ")
		var increasing bool = true
		for i, v := range vals {
			fmt.Println("vals len ", len(vals))
			if i+1 == len(vals) {
				fmt.Println("Line SAFE ", line)
				safeRoutesCount++
				fmt.Println("Safe Routes Count is now ", safeRoutesCount)
				break
			}
			curr, _ := strconv.Atoi(v)
			next, _ := strconv.Atoi(vals[i+1])

			if i == 0 && next < curr {
				increasing = false
			}
			fmt.Println(increasing)

			if increasing && next < curr {
				fmt.Println("Here when increasing fucks up")
				break
			} else if !increasing && curr < next {
				fmt.Println("Was decreasing, now increasing ", curr, next)
				break
			} else if 0 == int(math.Abs(float64(curr-next))) {
				fmt.Println("Same number twice %n, %n", curr, next)
				break
			} else if 3 < int(math.Abs(float64(curr-next))) {
				fmt.Println("Difference larger than 3 ", curr, next)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(safeRoutesCount)
}
