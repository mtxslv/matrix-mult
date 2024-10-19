package main

import (
	// "fmt"
	"log"
)

// compareResults compares two matrices (Strassen and regular) and returns true if they are identical.
func compareResults(strassen, regular [][]int) bool {
	n := len(strassen)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if strassen[i][j] != regular[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	// Example 1: Equal matrices
	strassen1 := [][]int{
		{19, 22, 43},
		{30, 38, 80},
		{51, 62, 134},
	}

	regular1 := [][]int{
		{19, 22, 43},
		{30, 38, 80},
		{51, 62, 134},
	}

	// Example 2: Different matrices
	strassen2 := [][]int{
		{19, 22, 43},
		{30, 38, 80},
		{51, 62, 134},
	}

	regular2 := [][]int{
		{19, 22, 44}, // Different value here (43 vs 44)
		{30, 38, 80},
		{51, 62, 134},
	}

	// Checking equal matrices
	if !compareResults(strassen1, regular1) {
		defer func() {
			if r := recover(); r != nil {
				log.Fatalf("Error: The matrices are not equal: %v", r)
			}
		}()
		log.Println("Strassen and regular matrices are equal!")
	}

	// Checking different matrices
	if !compareResults(strassen2, regular2) {
		defer func() {
			if r := recover(); r != nil {
				log.Fatalf("Error: The matrices are not equal: %v", r)
			}
		}()
		log.Fatal("Error: The matrices are not equal!")
	}
}
