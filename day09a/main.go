package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	posXHead := 0
	posYHead := 0

	posXTail := 0
	posYTail := 0

	tailPositions := make(map[int]map[int]bool)

	exp := regexp.MustCompile(`(\w) (\d+)`)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()

		match := exp.FindStringSubmatch(line)
		times, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err.Error())
		}

		for i := 0; i < times; i++ {
			switch match[1] {
			case "R":
				posXHead++
				if abs(posXHead-posXTail) > 1 {
					posXTail = posXHead - 1
					posYTail = posYHead
				}
			case "L":
				posXHead--
				if abs(posXHead-posXTail) > 1 {
					posXTail = posXHead + 1
					posYTail = posYHead
				}
			case "U":
				posYHead++
				if abs(posYHead-posYTail) > 1 {
					posXTail = posXHead
					posYTail = posYHead - 1
				}
			case "D":
				posYHead--
				if abs(posYHead-posYTail) > 1 {
					posXTail = posXHead
					posYTail = posYHead + 1
				}
			}

			if _, ok := tailPositions[posXTail]; !ok {
				tailPositions[posXTail] = make(map[int]bool)
			}

			tailPositions[posXTail][posYTail] = true
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}

	positions := 0
	for _, cols := range tailPositions {
		for range cols {
			positions++
		}
	}

	fmt.Println(positions)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
