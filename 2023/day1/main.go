package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums []int
	for scanner.Scan() {
		text := scanner.Text()
		// fmt.Printf(text)
		nums = append(nums, processString(text))
	}

	total := 0
	for _, v := range nums {
		total += v
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processString(text string) int {
	left, right := 0, len(text)-1

	for left < len(text) && (text[left] < '0' || text[left] > '9') {
		left++
	}

	for right >= 0 && (text[right] < '0' || text[right] > '9') {
		right--
	}

	if left >= len(text) || right < 0 {
		return 0
	}
	leftDigit := int(text[left] - '0')
	rightDigit := int(text[right] - '0')
	combined := leftDigit*10 + rightDigit
	return combined

}
