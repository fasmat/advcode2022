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

	var setup []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			break
		}
		setup = append(setup, line)
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}

	exp := regexp.MustCompile(`([0-9]+)`)
	stacksRegEx := exp.FindAllStringSubmatch(setup[len(setup)-1], -1)
	stacks := make([]Stack[string], len(stacksRegEx))

	exp = regexp.MustCompile(`\[([A-Z])\]| ?(   )`)
	for i := len(setup) - 2; i >= 0; i-- {
		level := exp.FindAllStringSubmatch(setup[i], -1)

		for j, l := range level {
			if l[1] != "" {
				stacks[j] = stacks[j].push(l[1])
			}
		}
	}

	for _, s := range stacks {
		fmt.Println(len(s), s)
	}

	exp = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for scan.Scan() {
		cmd := exp.FindStringSubmatch(scan.Text())

		times, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatal(err.Error())
		}

		from, err := strconv.Atoi(cmd[2])
		if err != nil {
			log.Fatal(err.Error())
		}

		to, err := strconv.Atoi(cmd[3])
		if err != nil {
			log.Fatal(err.Error())
		}

		var tempStack Stack[string]
		for i := 0; i < times; i++ {
			var item string
			stacks[from-1], item = stacks[from-1].pop()
			tempStack = tempStack.push(item)
		}

		for len(tempStack) > 0 {
			var item string
			tempStack, item = tempStack.pop()
			stacks[to-1] = stacks[to-1].push(item)
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("----")

	for _, s := range stacks {
		fmt.Println(len(s), s)
	}
}

type Stack[T any] []T

func (s Stack[T]) push(item T) Stack[T] {
	return append(s, item)
}

func (s Stack[T]) pop() (Stack[T], T) {
	if len(s) == 0 {
		var zero T
		return s, zero
	}

	item := s[len(s)-1]
	s = s[:len(s)-1]
	return s, item
}
