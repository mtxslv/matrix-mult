package main

import (
    "fmt"
    "math/rand/v2"
)

// generateMatrix creates an n x n matrix with all elements set to the provided value or a random integer
func generateMatrix(n, value int, r *rand.Rand) [][]int {
    // Initialize a new random generator using PCG

    matrix := make([][]int, n) // Create an array of arrays (n x n)

    for i := 0; i < n; i++ {
        matrix[i] = make([]int, n) // Create each row
        for j := 0; j < n; j++ {
            if value < 0 { // If value is not set (assuming 0 as the unset marker)
                matrix[i][j] = int(r.Int32N(150)) // Use random integer between 0 and 149
            } else {
                matrix[i][j] = value // Set the provided value
            }
        }
    }

    return matrix
}

func displayMatrix(matrix [][]int) {
    for _, row := range matrix {
        fmt.Println(row) // Print each row of the matrix
    }
}

func testMatrixGeneration(){
    n := 4
    r := rand.New(rand.NewPCG(1, 2))
    matrix_a := generateMatrix(n, -1, r) // trigger random number
    matrix_b := generateMatrix(n, -1, r) // trigger random number
	matrix_c := generateMatrix(n, 0, r)

	fmt.Println("MATRIX A:")
	displayMatrix(matrix_a)
	fmt.Println("MATRIX B:")
	displayMatrix(matrix_b)
	fmt.Println("MATRIX C:")
	displayMatrix(matrix_c)
}

func main() {
	testMatrixGeneration()
}
