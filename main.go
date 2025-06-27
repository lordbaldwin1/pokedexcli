package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Poxedex > ")
		scanner.Scan()
		text := scanner.Text()

		cleanedLine := cleanInput(text)
		fmt.Println("Your command was:", cleanedLine[0])
	}
}

// input:    " hello world ",
// expected: []string{"hello", "world"},
func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	strSlice := strings.Split(strings.ToLower(text), " ")
	var res []string

	for _, str := range strSlice {
		trimmed := strings.TrimSpace(str)

		if trimmed != "" {
			res = append(res, trimmed)
		}
	}

	return res
}
