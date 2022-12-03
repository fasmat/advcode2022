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
		set1 := make(map[rune]struct{})
		for _, item := range line {
			set1[item] = struct{}{}
		}

		scan.Scan()
		line = scan.Text()
		set2 := make(map[rune]struct{})
		for _, item := range line {
			set2[item] = struct{}{}
		}

		scan.Scan()
		line = scan.Text()
		set3 := make(map[rune]struct{})
		for _, item := range line {
			set3[item] = struct{}{}
		}

		for item := range set1 {
			_, ok2 := set2[item]
			_, ok3 := set3[item]
			if !ok2 || !ok3 {
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
