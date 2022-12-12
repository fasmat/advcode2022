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

	cycleDiff := make([]int, 1, 2000)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		if line == "noop" {
			cycleDiff = append(cycleDiff, 0)
			continue
		}

		value, err := strconv.Atoi(line[5:])
		if err != nil {
			log.Fatal(err.Error())
		}

		cycleDiff = append(cycleDiff, 0, value)
	}

	sum := 1
	result := 0
	for i, value := range cycleDiff {
		if i%40 == 20 {
			fmt.Println(i, sum, i*sum)
			result += i * sum
		}
		sum += value
	}

	fmt.Println(result)

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}
}
