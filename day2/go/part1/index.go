package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	strFile := strings.TrimSpace(string(file))
	idRanges := strings.Split(strFile, ",")

	for _, idRange := range idRanges {
		startAndEnd := strings.Split(idRange, "-")
		startStr := startAndEnd[0]
		endStr := startAndEnd[1]

		start, err := strconv.Atoi(startStr)
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(endStr)
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			var strId string = strconv.Itoa(i)
			length := len(strId)

			if length%2 != 0 {
				continue
			}

			firstHalf := strId[0 : length/2]
			secondHalf := strId[length/2:]

			if firstHalf == secondHalf {
				num, err := strconv.Atoi(strId)
				if err != nil {
					panic(err)
				}

				sum += num
			}
		}
	}

	fmt.Println(sum)
}
