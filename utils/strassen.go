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

// splitMatrix splits a matrix into four submatrices (quadrants)
func splitMatrix(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {
    size := len(matrix)
    half := size / 2

    // Create slices for the four quadrants
    A11 := make([][]int, half)
    A12 := make([][]int, half)
    A21 := make([][]int, half)
    A22 := make([][]int, half)

    for i := 0; i < half; i++ {
        A11[i] = make([]int, half)
        A12[i] = make([]int, half)
        A21[i] = make([]int, half)
        A22[i] = make([]int, half)
        
        for j := 0; j < half; j++ {
            A11[i][j] = matrix[i][j]
            A12[i][j] = matrix[i][j+half]
            A21[i][j] = matrix[i+half][j]
            A22[i][j] = matrix[i+half][j+half]
        }
    }

    return A11, A12, A21, A22
}

func main() {

}


    // TESTING MATRIX SPLITTING
    // matrix := [][]int{
        // {1, 2, 3, 4},
        // {5, 6, 7, 8},
        // {9, 10, 11, 12},
        // {13, 14, 15, 16},
    // }
// 
    // A11, A12, A21, A22 := splitMatrix(matrix)
// 
    // fmt.Println("A11:", A11)
    // fmt.Println("A12:", A12)
    // fmt.Println("A21:", A21)
    // fmt.Println("A22:", A22)