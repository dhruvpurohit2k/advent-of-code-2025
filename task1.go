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
	var count int = 0
	var dialPosition int = 50
	for _, line := range lines {
		var direction int
		if line[0] == 'L' {
			direction = -1
		} else {
			direction = 1
		}
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		// for _ = range rotation {
		// 	if direction == -1 {
		// 		dialPosition = (dialPosition - 1 + 100) % 100
		// 	} else {
		// 		dialPosition = (dialPosition + 1) % 100
		// 	}
		// 	if dialPosition == 0 {
		// 		count++
		// 	}
		// }
		count += rotation / 100
		rotation %= 100
		newDialPosition := dialPosition + (rotation * direction)
		if newDialPosition < 0 {
			newDialPosition = 100 + (newDialPosition % 100)
		} else {
			newDialPosition %= 100
		}
		if direction == -1 {
			if (dialPosition-newDialPosition) < 0 && dialPosition != 0 {
				count++
			} else if newDialPosition == 0 {
				count++
			}
		}
		if direction == 1 {
			if (dialPosition-newDialPosition) > 0 && dialPosition != 0 {
				count++
			} else if newDialPosition == 0 {
				count++
			}
		}
		dialPosition = newDialPosition
	}
	fmt.Println(count)
}
