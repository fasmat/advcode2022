package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	trees := make([][]int, 0)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		row := make([]int, len(line))

		for x := range line {
			height, err := strconv.Atoi(string(line[x]))
			if err != nil {
				log.Fatal(err.Error())
			}

			row[x] = height
		}

		trees = append(trees, row)
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}

	maxScore := 0
	maxX := 0
	maxY := 0
	score := make([][]int, len(trees))
	for y := range trees {
		score[y] = make([]int, len(trees[y]))

		for x := range trees[y] {
			if x == 0 || x == len(trees[y])-1 {
				score[y][x] = 0
				continue
			}

			if y == 0 || y == len(trees)-1 {
				score[y][x] = 0
				continue
			}

			score[y][x] = calcScore(x, y, trees)

			if score[y][x] > maxScore {
				maxScore = score[y][x]
				maxX = x
				maxY = y
			}
		}
	}

	for _, row := range trees {
		fmt.Println(row)
	}

	fmt.Println()

	for _, row := range score {
		fmt.Println(row)
	}

	fmt.Println()

	fmt.Println("Max Score at", maxX, maxY, "of", maxScore)
}

func calcScore(x, y int, tree [][]int) int {
	// check to the left
	scoreLeft := 0
	for x2 := x - 1; x2 >= 0; x2-- {
		scoreLeft++
		if tree[y][x] <= tree[y][x2] {
			break
		}
	}

	// check to the right
	scoreRight := 0
	for x2 := x + 1; x2 < len(tree[y]); x2++ {
		scoreRight++
		if tree[y][x] <= tree[y][x2] {
			break
		}
	}

	// check to the top
	scoreTop := 0
	for y2 := y - 1; y2 >= 0; y2-- {
		scoreTop++
		if tree[y][x] <= tree[y2][x] {
			break
		}
	}

	// check to the bottom
	scoreBottom := 0
	for y2 := y + 1; y2 < len(tree); y2++ {
		scoreBottom++
		if tree[y][x] <= tree[y2][x] {
			break
		}
	}

	return scoreLeft * scoreRight * scoreTop * scoreBottom
}
