package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	elf := make([]int, 1)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			elf = append(elf, 0)
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err.Error())
		}

		elf[len(elf)-1] += cal
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}

	slices.Sort(elf)
	fmt.Printf("Max calories: %d\n", elf[len(elf)-1])
}
