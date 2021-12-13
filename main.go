package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func fold(points *[]Point, dir string, idx int) {
	minX, minY := 0, 0

	// Do the folding
	for i := 0; i < len(*points); i++ {
		point := &(*points)[i]

		switch dir {
		case "y":
			if point.Y == idx {
				// Remove this point since this point is located at the folding index.
				*points = append((*points)[:i], (*points)[i+1:]...)
				continue
			}

			if point.Y > idx {
				diff := point.Y - idx
				point.Y = idx - diff

				if minY > point.Y {
					minY = point.Y
				}
			}

		case "x":
			if point.X == idx {
				// Remove this point since this point is located at the folding index.
				*points = append((*points)[:i], (*points)[i+1:]...)
				continue
			}

			if point.X > idx {
				diff := point.X - idx
				point.X = idx - diff

				if minX > point.X {
					minX = point.X
				}
			}
		}
	}

	// Ensure that the minimum x and y is 0
	diffX := 0
	if minX < 0 {
		diffX = 0 - minX
	}

	diffY := 0
	if minY < 0 {
		diffY = 0 - minY
	}

	if diffX != 0 || diffY != 0 {
		for _, point := range *points {
			point.X += diffX
			point.Y += diffY
		}
	}
}

func countVisiblePoints(points []Point) int {
	pointsMap := make(map[string]int)
	count := 0

	for _, point := range points {
		k := fmt.Sprintf("%d,%d", point.X, point.Y)
		_, ok := pointsMap[k]

		if !ok {
			pointsMap[k] = 1
			count++
		}
	}

	return count
}

func main() {
	rc, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %s\n", err)
		return
	}

	c := string(rc)
	points := make([]Point, 0)
	lines := strings.Split(c, "\n")
	firstFold := true
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, ",") {
			xy := strings.Split(line, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])

			points = append(points, Point{X: x, Y: y})
			continue
		}

		if strings.HasPrefix(line, "fold along") {
			// Perform the folding
			foldData := strings.Split(line, "=")
			dir := foldData[0][len(foldData[0])-1:]
			idx, _ := strconv.Atoi(foldData[1])

			fold(&points, dir, idx)

			if firstFold {
				firstFold = false
				fmt.Println("PART 1 ANSWER")
				fmt.Println(countVisiblePoints(points))
			}
		}
	}

	fmt.Println("PART 2 ANSWER")

	minX, minY, maxX, maxY := 0, 0, 0, 0
	pointsMap := make(map[string]int)
	for _, point := range points {
		if point.X < minX {
			minX = point.X
		}

		if point.X > maxX {
			maxX = point.X
		}

		if point.Y < minY {
			minY = point.Y
		}

		if point.Y > maxY {
			maxY = point.Y
		}

		k := fmt.Sprintf("%d,%d", point.X, point.Y)
		pointsMap[k] = 1
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			k := fmt.Sprintf("%d,%d", x, y)
			_, ok := pointsMap[k]

			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
