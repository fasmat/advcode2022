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
		game := scan.Text()

		switch game[2] {
		case 'X':
			played := int(game[0])%3 + 1
			score += 0 + played
		case 'Y':
			played := int(game[0]+1)%3 + 1
			score += 3 + played
		case 'Z':
			played := int(game[0]+2)%3 + 1
			score += 6 + played
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(score)
}
