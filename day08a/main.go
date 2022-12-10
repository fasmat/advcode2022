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

	visible := make([][]bool, len(trees))

	for y := range trees {
		visible[y] = make([]bool, len(trees[y]))

		for x := range trees[y] {
			if y == 0 || y == len(trees)-1 {
				visible[y][x] = true
				continue
			}

			if x == 0 || x == len(trees[y])-1 {
				visible[y][x] = true
				continue
			}

			visible[y][x] = checkVisibility(x, y, trees)
		}
	}

	count := 0
	for i := range trees {
		for j := range trees[i] {
			if visible[i][j] {
				count++
			}
		}
	}

	for _, row := range trees {
		fmt.Println(row)
	}

	fmt.Println()

	for _, row := range visible {
		fmt.Println(row)
	}

	fmt.Println()

	fmt.Println(count)
}

func checkVisibility(x, y int, tree [][]int) bool {
	// check from left
	visible := true
	for x2 := 0; x2 < x; x2++ {
		if tree[y][x] <= tree[y][x2] {
			visible = false
			break
		}
	}

	if visible {
		fmt.Println(x, y, "visible from left")
		return true
	}

	// check from right
	visible = true
	for x2 := len(tree[y]) - 1; x2 > x; x2-- {
		if tree[y][x] <= tree[y][x2] {
			visible = false
			break
		}
	}

	if visible {
		fmt.Println(x, y, "visible from right")
		return true
	}

	// check from top
	visible = true
	for y2 := 0; y2 < y; y2++ {
		if tree[y][x] <= tree[y2][x] {
			visible = false
			break
		}
	}

	if visible {
		fmt.Println(x, y, "visible from top")
		return true
	}

	// check from bottom
	for y2 := len(tree) - 1; y2 > y; y2-- {
		if tree[y][x] <= tree[y2][x] {
			return false
		}
	}

	fmt.Println(x, y, "visible from bottom")
	return true
}
