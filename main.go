package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cells [30][30]bool

func main() {
	cells := firstGeneration()
	printCells(cells)

	for {
		time.Sleep(time.Second)
		cells = nextGeneration(cells)
		fmt.Println()
		printCells(cells)
	}
}

func nextGeneration(cells cells) cells {
	nextCells := cells

	for rowIndex := 0; rowIndex < len(cells); rowIndex++ {
		// repeated instructions go here:
		for colIndex := 0; colIndex < len(cells[rowIndex]); colIndex++ {
			neighbourCount := calculateNeighbourCount(cells, rowIndex, colIndex)
			alive := cells[rowIndex][colIndex]

			// && operator has a gigher priority than || operator -> needs parenthesis
			if alive && (neighbourCount == 2 || neighbourCount == 3) {
				// cell stays alive due to good amount of neighbours
				nextCells[rowIndex][colIndex] = true
			} else if !alive && neighbourCount == 3 {
				// cell becomes alive
				nextCells[rowIndex][colIndex] = true

			} else {
				// cell dies due to overpopulation or says dead
				nextCells[rowIndex][colIndex] = false
			}
		}
	}
	return nextCells
}

func firstGeneration() cells {
	rand.NewSource(time.Now().UnixNano())

	// by default boolean fields are false
	var cells cells
	rowStart := len(cells)/2 - 5
	rowEnd := len(cells)/2 + 4
	colStart := len(cells[0])/2 - 5
	colEnd := len(cells[0])/2 + 4

	fmt.Printf("rowStart: %d, rowEnd: %d, colStart: %d, colEnd: %d\n", rowStart, rowEnd, colStart, colEnd)

	for rowIndex := rowStart; rowIndex < rowEnd; rowIndex++ {
		for colIndex := colStart; colIndex < colEnd; colIndex++ {
			if rowIndex >= 0 && rowIndex < len(cells) && colIndex >= 0 && colIndex < len(cells[0]) {
				// 1/4 chance of cell being alive
				if rand.Intn(4) == 0 {
					cells[rowIndex][colIndex] = true
				}
			}
		}
	}
	return cells
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func calculateNeighbourCount(cells cells, curRow, curCol int) int {
	rowStart := max(curRow-1, 0)
	rowEnd := min(curRow+1, len(cells)-1)
	colStart := max(curCol-1, 0)
	colEnd := min(curCol+1, len(cells[0])-1)

	neighbourCount := 0
	for rowIndex := rowStart; rowIndex <= rowEnd; rowIndex++ {
		for colIndex := colStart; colIndex <= colEnd; colIndex++ {
			isRefcell := rowIndex == curRow && colIndex == curCol
			// increase neighbours count if this is not our reference cell and there is neighbour alive
			if !isRefcell && cells[rowIndex][colIndex] {
				neighbourCount++
			}
		}
	}
	return neighbourCount
}

func printCells(cells cells) {
	for rowIndex := 0; rowIndex < len(cells[rowIndex]); rowIndex++ {
		for colIndex := 0; colIndex < len(cells[rowIndex]); colIndex++ {
			if cells[rowIndex][colIndex] {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
