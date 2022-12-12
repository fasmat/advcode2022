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

	posX := make([]int, 10)
	posY := make([]int, 10)

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
				posX[0]++
			case "L":
				posX[0]--
			case "U":
				posY[0]++
			case "D":
				posY[0]--
			}

			for i := 1; i < len(posX); i++ {
				switch {
				case posX[i-1]-posX[i] > 1 && posY[i-1]-posY[i] > 1: // RU
					posX[i] = posX[i-1] - 1
					posY[i] = posY[i-1] - 1
				case posX[i-1]-posX[i] > 1 && posY[i-1]-posY[i] < -1: // RD
					posX[i] = posX[i-1] - 1
					posY[i] = posY[i-1] + 1
				case posX[i-1]-posX[i] < -1 && posY[i-1]-posY[i] > 1: // LU
					posX[i] = posX[i-1] + 1
					posY[i] = posY[i-1] - 1
				case posX[i-1]-posX[i] < -1 && posY[i-1]-posY[i] < -1: // LD
					posX[i] = posX[i-1] + 1
					posY[i] = posY[i-1] + 1
				case posX[i-1]-posX[i] > 1: // R
					posX[i] = posX[i-1] - 1
					posY[i] = posY[i-1]
				case posX[i-1]-posX[i] < -1: // L
					posX[i] = posX[i-1] + 1
					posY[i] = posY[i-1]
				case posY[i-1]-posY[i] > 1: // U
					posX[i] = posX[i-1]
					posY[i] = posY[i-1] - 1
				case posY[i-1]-posY[i] < -1: // D
					posX[i] = posX[i-1]
					posY[i] = posY[i-1] + 1
				default:
				}
			}

			if _, ok := tailPositions[posX[9]]; !ok {
				tailPositions[posX[9]] = make(map[int]bool)
			}

			tailPositions[posX[9]][posY[9]] = true
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
