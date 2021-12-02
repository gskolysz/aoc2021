package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	value     int
}

type position struct {
	horizontal int
	depth      int
	aim        int
}

func (currentPosition *position) adjustPosition(comm command) *position {
	if comm.direction == "forward" {
		currentPosition.horizontal = currentPosition.horizontal + comm.value
	} else if comm.direction == "up" {
		currentPosition.depth = currentPosition.depth - comm.value
	} else if comm.direction == "down" {
		currentPosition.depth = currentPosition.depth + comm.value
	}
	return currentPosition
}

func (currentPosition *position) adjustPositionWithAim(comm command) *position {
	if comm.direction == "forward" {
		currentPosition.horizontal = currentPosition.horizontal + comm.value
		currentPosition.depth = currentPosition.depth + (currentPosition.aim * comm.value)
	} else if comm.direction == "up" {
		currentPosition.aim = currentPosition.aim - comm.value
	} else if comm.direction == "down" {
		currentPosition.aim = currentPosition.aim + comm.value
	}
	return currentPosition
}

func readLines() ([]string, error) {
	file, err := os.Open("input.txt")
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

func linesToCommands(lines []string) ([]command, error) {
	commands := make([]command, len(lines))
	for i, b := range lines {
		var comm command

		splitted := strings.Split(b, " ")
		comm.direction = splitted[0]
		intVal, err := strconv.Atoi(splitted[1])
		if err != nil {
			return nil, err
		}
		comm.value = intVal
		commands[i] = comm
	}
	return commands, nil
}

func calculateMovement(commands []command, position *position) *position {
	for _, c := range commands {
		position.adjustPosition(c)
	}
	return position
}

func calculateMovementWithAim(commands []command, position *position) *position {
	for _, c := range commands {
		position.adjustPositionWithAim(c)
	}
	return position
}

func solve1(commands []command) {
	var currentPosition *position = &position{0, 0, 0}
	calculateMovement(commands, currentPosition)
	fmt.Println(*currentPosition)
	product := currentPosition.depth * currentPosition.horizontal
	fmt.Printf("Part1: %d\n", product)
}

func solve2(commands []command) {
	var currentPosition *position = &position{0, 0, 0}
	calculateMovementWithAim(commands, currentPosition)
	fmt.Println(*currentPosition)
	product := currentPosition.depth * currentPosition.horizontal
	fmt.Printf("Part2: %d\n", product)
}

func main() {
	lines, err := readLines()
	if err != nil {
		panic(err)
	}
	commands, err := linesToCommands(lines)
	if err != nil {
		panic(err)
	}
	solve1(commands)
	solve2(commands)
}
