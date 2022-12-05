package main

import (
	"aoc22_5/pkg/instruction"
	"aoc22_5/pkg/stack"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("AOC22 Day 5")
	partOne()
	fmt.Println("\n")
	partTwo()
}

func partOne() {
	input := readInput("input.txt")
	fmt.Printf("Read %d lines \n", len(input))
	println("Reading cargo picture")
	stacks := getStacks(input)
	instructions := getInstructions(input)
	apply(instructions, stacks)
	for _, stack := range stacks {
		v, _ := stack.Pop()
		fmt.Print(v)
	}
}

func partTwo() {
	input := readInput("input.txt")
	fmt.Printf("Read %d lines \n", len(input))
	println("Reading cargo picture")
	stacks := getStacks(input)
	instructions := getInstructions(input)
	apply9001(instructions, stacks)
	for _, stack := range stacks {
		v, _ := stack.Pop()
		fmt.Print(v)
	}
}
func apply(instructions []instruction.Instruction, stacks []stack.Stack) {
	for _, instruction := range instructions {
		if instruction.Action == "move" {
			for i := 0; i < instruction.Amount; i++ {
				e, _ := stacks[instruction.Origin-1].Pop()
				stacks[instruction.Destination-1].Push(e)
			}
		}
	}
}
func apply9001(instructions []instruction.Instruction, stacks []stack.Stack) {
	for _, instruction := range instructions {
		elements := []string{}
		if instruction.Action == "move" {
			for i := 0; i < instruction.Amount; i++ {
				e, _ := stacks[instruction.Origin-1].Pop()
				elements = append(elements, e)
			}
			for i := len(elements) - 1; i >= 0; i-- {
				stacks[instruction.Destination-1].Push(elements[i])
			}
		}
	}
}

func getInstructions(input []string) []instruction.Instruction {
	lines := skipToInstructions(input)
	instructions := []instruction.Instruction{}
	for _, line := range lines {
		instructions = append(instructions, toInstruction(line))
	}
	return instructions
}
func toInstruction(line string) instruction.Instruction {
	var amount int
	var origin int
	var destination int
	fmt.Sscanf(line, "move %v from %v to %v", &amount, &origin, &destination)
	return instruction.Instruction{"move", amount, origin, destination}
}

func skipToInstructions(input []string) []string {
	for i, line := range input {
		if line == "" {
			return input[i+1:]
		}
	}
	return []string{}
}

func getStacks(input []string) []stack.Stack {

	stackDelimiter := " "
	const stackWidth int = 3
	const expectedWidth int = 9*(stackWidth+1) - 1

	if len(input) < 9*(stackWidth+1)-1 {
		fmt.Printf("The lenght of the stack describing lines is too short. Expected is %d but was %d \n", expectedWidth, len(input))
		fmt.Print(stackDelimiter)
		os.Exit(1)
	}

	stacks := []stack.Stack{
		stack.Stack{},
		stack.Stack{},
		stack.Stack{},
		stack.Stack{},
		stack.Stack{},
		stack.Stack{},
		stack.Stack{},
		stack.Stack{},
		stack.Stack{},
	}

	cargoItems := [][]string{}
	for i, line := range input {
		if line == "" {
			break
		}
		cargoItems = append(cargoItems, extractCargoItems(line))
		fmt.Printf("%d: %s \n", i, line)
	}
	for _, items := range cargoItems {
		for stackIndex, item := range items {
			if item != " " {
				stacks[stackIndex].PushBack(item)
			}
		}

	}
	return stacks
}
func extractCargoItems(line string) []string {
	cargoItems := []string{}

	for i, char := range line {
		if (i == 1) ||
			(i == 5) ||
			(i == 9) ||
			(i == 13) ||
			(i == 17) ||
			(i == 21) ||
			(i == 25) ||
			(i == 29) ||
			(i == 33) {
			if _, err := strconv.Atoi(string(char)); err != nil {
				cargoItems = append(cargoItems, string(char))
			}
		}
	}
	return cargoItems
}

func readInput(inputFile string) []string {
	fileHandle, err := os.Open(inputFile)
	if err != nil {
		println("Someting went wrong opening the file")
		println(err.Error())
		os.Exit(1)
	}
	reader := bufio.NewReader(fileHandle)
	line, hasNext := readNextLine(reader)

	lines := []string{}
	for hasNext {
		lines = append(lines, line)
		line, hasNext = readNextLine(reader)
	}
	return lines
}

func readNextLine(reader *bufio.Reader) (string, bool) {
	line, err := reader.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			return "", false
		}
		println("Somethig went wrong reading the file")
		println(err.Error())
		os.Exit(1)
	}
	return strings.Trim(line, "\n"), true
}
