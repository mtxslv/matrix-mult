package main

import "fmt"

func matrixAdd(a, b, c [][]int, n int) [][]int {
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            c[i][j] = a[i][j] + b[i][j]
        }
    }
    return c
}

func matrixSubtract(a, b, c [][]int, n int) [][]int {
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            c[i][j] = a[i][j] - b[i][j]
        }
    }
    return c
}



func main() {
    // Define matrix_a and matrix_b as 4x4 matrices
    matrix_a := [][]int{
        {1, 2, 0, 0},
        {3, 4, 0, 0},
        {0, 0, 1, 2},
        {0, 0, 3, 4},
    }

    matrix_b := [][]int{
        {5, 6, 0, 0},
        {7, 8, 0, 0},
        {0, 0, 5, 6},
        {0, 0, 7, 8},
    }

    // Initialize matrix_c to store the result
    matrix_c := [][]int{
        {0, 0, 0, 0},
        {0, 0, 0, 0},
        {0, 0, 0, 0},
        {0, 0, 0, 0},
    }

    // Perform matrix multiplication
    matrix_c = matrixAdd(matrix_a, matrix_b, matrix_c, 4)

    // Print the result matrix_c
    fmt.Println("Matrix C (result of A + B):")
    for _, row := range matrix_c {
        fmt.Println(row)
    }

    // Perform matrix multiplication
    matrix_c = matrixSubtract(matrix_a, matrix_b, matrix_c, 4)

    // Print the result matrix_c
    fmt.Println("Matrix C (result of A - B):")
    for _, row := range matrix_c {
        fmt.Println(row)
    }	
}
