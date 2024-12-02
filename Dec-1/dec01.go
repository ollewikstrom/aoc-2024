package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var listOne []int
	var listTwo []int

	//Split the list into the two arrays
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		var entry string = scanner.Text()
		val, _ := strconv.Atoi(entry)

		if count == 0 || count%2 == 0 {
			listOne = append(listOne, val)
		} else {
			listTwo = append(listTwo, val)
		}
		count++
	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	var totalDist int = 0
	defer fmt.Println(totalDist)

	for i, v := range listOne {
		// fmt.Println(v)
		// fmt.Println(listTwo[i])
		// fmt.Println(int(math.Abs(float64(v - listTwo[i]))))

		dist := int(math.Abs(float64(v - listTwo[i])))
		// fmt.Println(dist)
		totalDist += dist
		fmt.Println(totalDist)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
