package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads input file and returns slice of int arrays, containing bits
func readLines() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bits := strings.Split(scanner.Text(), "")
		bitLen := len(bits)
		bitsInt := make([]int, bitLen)
		for i, bit := range bits {
			bitInt, err := strconv.Atoi(bit)
			if err != nil {
				return nil, err
			}
			bitsInt[i] = bitInt
		}
		lines = append(lines, bitsInt)

	}
	return lines, scanner.Err()
}

func calculateSums(lines [][]int) []int {
	sum := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, bit := range line {
			sum[i] += bit
		}
	}
	return sum
}

func determineMostCommonBit(sums []int, maxLen int) []int {
	commonBits := make([]int, len(sums))
	for i, onesCount := range sums {
		if onesCount > maxLen/2 {
			commonBits[i] = 1
		} else {
			commonBits[i] = 0
		}
	}
	return commonBits
}

func inverseBitArray(bits []int) []int {
	inversedBits := make([]int, len(bits))
	for i, bit := range bits {
		if bit == 1 {
			inversedBits[i] = 0
		} else {
			inversedBits[i] = 1
		}
	}
	return inversedBits
}

func bitIntArrayAsInt(bits []int) (int, error) {
	chars := make([]string, len(bits))
	for i, bit := range bits {
		chars[i] = strconv.Itoa(bit)
	}
	string := strings.Join(chars, "")
	number, err := strconv.ParseInt(string, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(number), nil
}

func reduceByMostCommon(lines [][]int) []int {
	lineLen := len(lines[0])
	for i := 0; i < lineLen; i++ {
		leftLines := [][]int{}
		currLinesLen := len(lines)
		sums := calculateSums(lines)
		mcb := 0
		if sums[i] >= currLinesLen/2 {
			mcb = 1
		}
		for _, bits := range lines {
			if bits[i] == mcb {
				leftLines = append(leftLines, bits)
			}
		}
		if len(leftLines) == 1 {
			return leftLines[0]
		}
		lines = leftLines
	}
	return nil
}

func reduceByLeastCommon(lines [][]int) []int {
	lineLen := len(lines[0])
	for i := 0; i < lineLen; i++ {
		leftLines := [][]int{}
		currLinesLen := len(lines)
		sums := calculateSums(lines)
		lcb := 0
		if sums[i] >= currLinesLen/2 {
			lcb = 1
		}

		for _, bits := range lines {
			if bits[i] != lcb {

				leftLines = append(leftLines, bits)
			}
		}
		if len(leftLines) == 1 {
			return leftLines[0]
		}
		lines = leftLines
	}
	return nil
}

func main() {
	lines, err := readLines()
	if err != nil {
		panic(err)
	}
	lineCount := len(lines)
	sums := calculateSums(lines)
	fmt.Println(sums)
	mcb := determineMostCommonBit(sums, lineCount)
	inversed := inverseBitArray(mcb)
	mcbInt, err := bitIntArrayAsInt(mcb)
	if err != nil {
		panic(err)
	}
	inversedInt, err := bitIntArrayAsInt(inversed)
	if err != nil {
		panic(err)
	}
	sol1 := mcbInt * inversedInt
	fmt.Println(sol1)

	oxy := reduceByMostCommon(lines)
	oxyInt, err := bitIntArrayAsInt(oxy)
	if err != nil {
		panic(err)
	}
	fmt.Println(oxy)

	co2 := reduceByLeastCommon(lines)
	co2Int, err := bitIntArrayAsInt(co2)
	if err != nil {
		panic(err)
	}
	sol2 := co2Int * oxyInt
	fmt.Println(sol2)

}
