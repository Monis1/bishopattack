package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bwX, bwY, err := readPosition("Enter black bishop position: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bbX, bbY, err := readPosition("Enter white bishop position: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if areBishopsAttacking(bbX, bbY, bwX, bwY) {
		fmt.Println("They can attack each other.")
	} else {
		fmt.Println("They cannot attack each other.")
	}
}

// readPosition prompts the user to enter a chess position (e.g., "C1") and returns its coordinates (0-indexed).
func readPosition(prompt string) (int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	if len(input) != 2 {
		return 0, 0, fmt.Errorf("input must be 2 characters (e.g. 'C1')")
	}

	file := input[0]
	rank := input[1]

	if file < 'a' || file > 'h' || rank < '1' || rank > '8' {
		return 0, 0, fmt.Errorf("invalid square: %s", input)
	}

	x := int(file - 'a')
	y := int(rank - '1')
	return x, y, nil
}

// isWithinBounds checks if the given x and y coordinates are within the 8x8 chess board.
func isWithinBounds(x int, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

// areBishopsAttacking checks if a bishop at (startX, startY) can attack a bishop at (endX, endY).
func areBishopsAttacking(b1X int, b1Y int, b2X int, b2Y int) bool {
	// Check the top-right diagonal.
	if b2X < b1X && b2Y > b1Y {
		x, y := b1X-1, b1Y+1
		for isWithinBounds(x, y) {
			if x == b2X && y == b2Y {
				return true
			}
			x--
			y++
		}
	}

	// Check the top-left diagonal.
	if b2X < b1X && b2Y < b1Y {
		x, y := b1X-1, b1Y-1
		for isWithinBounds(x, y) {
			if x == b2X && y == b2Y {
				return true
			}
			x--
			y--
		}
	}

	// Check the bottom-right diagonal.
	if b2X > b1X && b2Y > b1Y {
		x, y := b1X+1, b1Y+1
		for isWithinBounds(x, y) {
			if x == b2X && y == b2Y {
				return true
			}
			x++
			y++
		}
	}

	// Check the bottom-left diagonal.
	if b2X > b1X && b2Y < b1Y {
		x, y := b1X+1, b1Y-1
		for isWithinBounds(x, y) {
			if x == b2X && y == b2Y {
				return true
			}
			x++
			y--
		}
	}

	return false
}
