package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	println("AOC22 Day 5")
	partOne()
}

func partOne() {
	input := readInput("input.txt")
	fmt.Printf("Read %d lines \n", len(input))
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
