package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

type SVector struct {
	x, y int
}

func parseLine(line string) SVector {
	parts := strings.Split(line, " ")
	amount, err := strconv.Atoi(parts[1])

	if err != nil {
		log.Fatal("this should never ever happen")
	}

	if parts[0] == "forward" {
		return SVector{
			x: amount,
			y: 0,
		}
	} else if parts[0] == "up" {
		return SVector{
			x: 0,
			y: -amount,
		}
	}

	return SVector{
		x: 0,
		y: amount,
	}
}

func main() {
	lines := strings.Split(getInput(), "\n")

	pos := SVector{0, 0}
	for _, line := range lines {
		sv := parseLine(line)
		pos.x += sv.x
		pos.y += sv.y
	}

	fmt.Printf("point: %+v", pos)
}

