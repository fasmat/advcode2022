package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	currentDir := ""
	listDir := false
	exp := regexp.MustCompile(`(\d+|dir) (.*)`)

	dirs := make(map[string]int)
	files := make(map[string]int)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		if strings.HasPrefix(line, "$") {
			listDir = false
		}

		switch {
		case strings.HasPrefix(line, "$ cd"):
			currentDir = filepath.Join(currentDir, strings.TrimPrefix(line, "$ cd "))
			fmt.Println("change dir to:", currentDir)
		case line == "$ ls":
			listDir = true
			continue
		case listDir:
			match := exp.FindStringSubmatch(line)
			if match[1] == "dir" {
				dirs[filepath.Join(currentDir, match[2])] = 0
				continue
			}
			size, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err.Error())
			}
			files[filepath.Join(currentDir, match[2])] = size

			for dirpath := currentDir; dirpath != "/"; dirpath = filepath.Join(dirpath, "..") {
				dirs[dirpath] += size
			}
			dirs["/"] += size
		default:
			fmt.Println("----- unknown")
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err.Error())
	}

	totalSize := 0
	for _, size := range dirs {
		if size < 100000 {
			totalSize += size
		}
	}

	fmt.Println(totalSize)
}
