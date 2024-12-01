package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filename string) []string {
	result := make([]string, 0)
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal("could not open file", filename)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
