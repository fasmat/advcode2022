package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	score := 0
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()

		set := make(map[rune]struct{})
		for _, item := range line[:len(line)/2] {
			set[item] = struct{}{}
		}

		for _, item := range line[len(line)/2:] {
			if _, ok := set[item]; !ok {
				continue
			}

			if item > 'Z' {
				score += int(item-'a') + 1
				break
			}

			score += int(item-'A') + 27
			break
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(score)
}
