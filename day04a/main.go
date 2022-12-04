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

		if start1 <= start2 && end1 >= end2 {
			overlap++
			continue
		}

		if start2 <= start1 && end2 >= end1 {
			overlap++
			continue
		}
	}

	fmt.Println(overlap)
}
