package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/exp/constraints"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	overlap := 0
	exp := regexp.MustCompile(`([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)`)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		matches := exp.FindStringSubmatch(line)
		fmt.Println(matches)

		start1, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		end1, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal(err.Error())
		}

		start2, err := strconv.Atoi(matches[3])
		if err != nil {
			log.Fatal(err.Error())
		}
		end2, err := strconv.Atoi(matches[4])
		if err != nil {
			log.Fatal(err.Error())
		}

		start := max(start1, start2)
		end := min(end1, end2)

		if start <= end {
			overlap++
		}
	}

	fmt.Println(overlap)
}

func max[T constraints.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func min[T constraints.Ordered](v1, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}
