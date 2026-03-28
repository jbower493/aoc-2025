package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var path string = "../input.txt"
	data, err := os.ReadFile(path)
	check(err)

	var text string = string(data)

	zeros := 0
	dial := 50

	var lines []string = strings.Split(strings.TrimSpace(text), "\n")

	for _, line := range lines {
		direction := string(line[0])
		turnsStr := string(line[1:])
		turns, err := strconv.Atoi(turnsStr)
		check(err)

		for i := 0; i < turns; i++ {
			if direction == "L" {
				dial--
			} else {
				dial++
			}

			if dial == -1 {
				dial = 99
			} else if dial == 100 {
				dial = 0
			}
		}

		if dial == 0 {
			zeros++
		}
	}

	fmt.Println(zeros)
}
