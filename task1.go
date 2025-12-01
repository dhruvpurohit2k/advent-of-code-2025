package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filepath := "./task1"
	content, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewScanner(content)
	var lines []string
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}
	count := 0
	var dialPosition int64 = 50
	for _, line := range lines {
		var direction int64
		if line[0] == 'L' {
			direction = -1
		} else {
			direction = 1
		}
		rotation, err := strconv.ParseInt(line[1:], 10, 16)
		rotation %= 100
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("dial at : %d, rotating amnt : %d, ", dialPosition, rotation*direction)
		dialPosition += rotation * direction
		if dialPosition < 0 {
			dialPosition = 100 + (dialPosition % 100)
		} else {
			dialPosition %= 100
		}
		fmt.Printf("after rotation pos is : %d \n", dialPosition)
		if dialPosition == 0 {
			count++
		}
	}
	fmt.Println(count)
}
