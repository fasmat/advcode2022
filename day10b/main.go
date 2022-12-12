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

	cycle := make([]int, 0)
	scan := bufio.NewScanner(file)
	for i := 1; scan.Scan(); i++ {
		line := scan.Text()
		if line == "noop" {
			cycle = append(cycle, 0)
			continue
		}

		value, err := strconv.Atoi(line[5:])
		if err != nil {
			log.Fatal(err.Error())
		}

		cycle = append(cycle, 0, value)
	}

	cur := 1
	for i := 0; i < len(cycle); i++ {
		x := i%40 + 1
		printPixel(x, cur)
		if x == 40 {
			fmt.Printf("\n")
		}

		cur += cycle[i]
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}
}

func printPixel(x, cur int) {
	if x == cur ||
		x == cur+1 ||
		x == cur+2 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}
