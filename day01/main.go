package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readInts() []int {
	byteSlice, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	intSlice := make([]int, len(byteSlice))
	for i, b := range byteSlice {
		intSlice[i], err = strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
	}
	return intSlice
}

func CountIncreases(ints []int) int {
	counter := 0
	for i := range ints {
		if i == 0 {
			continue
		}
		if ints[i-1] < ints[i] {
			counter++
		}
	}
	return counter
}

func SumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func CountWindowedIncreases(ints []int) int {
	counter := 0
	for i := 0; i < len(ints)-3; i++ {
		window1 := SumInts(ints[i : i+3])
		window2 := SumInts(ints[i+1 : i+4])

		if window2 > window1 {
			counter++
		}
	}
	return counter
}

func main() {
	ints := readInts()
	count := CountIncreases(ints)
	fmt.Println(count)

	windowedCount := CountWindowedIncreases(ints)
	fmt.Println(windowedCount)
}
