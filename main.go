package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`Kill:.*:\s(.+)\skilled\s(.+)\sby\s(.+)`)

func main() {
	file, err := os.Open("./games.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		if match != nil {
			fmt.Printf("| %-32s | %-32s | %-32s |\n", match[1], match[2], match[3])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
