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

	games := make([]string, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		games = append(games, line)
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}

	points := make(map[string]int)

	points["A X"] = 3 + 1
	points["A Y"] = 6 + 2
	points["A Z"] = 0 + 3
	points["B X"] = 0 + 1
	points["B Y"] = 3 + 2
	points["B Z"] = 6 + 3
	points["C X"] = 6 + 1
	points["C Y"] = 0 + 2
	points["C Z"] = 3 + 3

	score := 0
	for _, game := range games {
		score += points[game]
	}

	fmt.Println(score)
}
