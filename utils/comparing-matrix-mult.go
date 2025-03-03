package main

import (
    "fmt"
    "log"
    "math/rand/v2"
    "math"
    "time"
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

// matrixMult performs matrix multiplication of two n x n matrices
func matrixMult(a, b, c [][]int, n int) [][]int {
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            sum := 0
            for k := 0; k < n; k++ {
                sum += a[i][k] * b[k][j]
            }
            c[i][j] = sum
        }
    }
    return c
}

func matrixAdd(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
    for i := 0; i < n; i++ {
		c[i] = make([]int, n)
        for j := 0; j < n; j++ {
            c[i][j] = a[i][j] + b[i][j]
        }
    }
    return c
}

func matrixSubtract(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
    for i := 0; i < n; i++ {
		c[i] = make([]int, n)
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

// combineMatrix combines four submatrices into one larger matrix
func combineMatrix(A11, A12, A21, A22 [][]int) [][]int {
    half := len(A11)  // Size of the submatrices (A11, A12, A21, A22)
    size := half * 2  // Size of the resulting larger matrix

    // Initialize the result matrix with all zeroes
    result := make([][]int, size)
    for i := 0; i < size; i++ {
        result[i] = make([]int, size)
    }

    // Fill the result matrix with values from the four submatrices
    for i := 0; i < half; i++ {
        for j := 0; j < half; j++ {
            result[i][j] = A11[i][j]            // Top-left
            result[i][j+half] = A12[i][j]       // Top-right
            result[i+half][j] = A21[i][j]       // Bottom-left
            result[i+half][j+half] = A22[i][j]  // Bottom-right
        }
    }

    return result
}

// Strassen performs Strassen matrix multiplication on two matrices A and B.
func strassen(A, B [][]int) [][]int {
    size := len(A)

    // Base case for 1x1 matrix
    if size == 1 {
        return [][]int{{A[0][0] * B[0][0]}}
    }

    // Split matrices into quadrants
    A11, A12, A21, A22 := splitMatrix(A)
    B11, B12, B21, B22 := splitMatrix(B)

    // Compute the 7 Strassen products
    M1 := strassen(matrixAdd(A11, A22), matrixAdd(B11, B22))       // M1 = (A11 + A22) * (B11 + B22)
    M2 := strassen(matrixAdd(A21, A22), B11)                       // M2 = (A21 + A22) * B11
    M3 := strassen(A11, matrixSubtract(B12, B22))                  // M3 = A11 * (B12 - B22)
    M4 := strassen(A22, matrixSubtract(B21, B11))                  // M4 = A22 * (B21 - B11)
    M5 := strassen(matrixAdd(A11, A12), B22)                       // M5 = (A11 + A12) * B22
    M6 := strassen(matrixSubtract(A21, A11), matrixAdd(B11, B12))  // M6 = (A21 - A11) * (B11 + B12)
    M7 := strassen(matrixSubtract(A12, A22), matrixAdd(B21, B22))  // M7 = (A12 - A22) * (B21 + B22)

    // Combine the results to get the final quadrants of the result matrix
    C11 := matrixAdd(matrixSubtract(matrixAdd(M1, M4), M5), M7)
    C12 := matrixAdd(M3, M5)
    C21 := matrixAdd(M2, M4)
    C22 := matrixAdd(matrixSubtract(matrixAdd(M1, M3), M2), M6)

    // Combine the four quadrants into a single matrix
    return combineMatrix(C11, C12, C21, C22)
}

// Main function to test Strassen matrix multiplication
func main() {
    r := rand.New(rand.NewPCG(1, 2))
    // Example matrices (4x4)
    for i := 1; i < 10; i++ {
		power := math.Pow(4, float64(i))
        n:=int(power)
        A := generateMatrix(n, -1, r) // trigger random number
        B := generateMatrix(n, -1, r) // trigger random number
        C_regular := generateMatrix(n, 0, r)
//         A := [][]int{
//             {1, 2, 0, 0, 1, 2, 0, 0},
//             {3, 4, 0, 0, 3, 4, 0, 0},
//             {0, 0, 1, 2, 0, 0, 1, 2},
//             {0, 0, 3, 4, 0, 0, 3, 4},
//             {1, 2, 0, 0, 1, 2, 0, 0},
//             {3, 4, 0, 0, 3, 4, 0, 0},
//             {0, 0, 1, 2, 0, 0, 1, 2},
//             {0, 0, 3, 4, 0, 0, 3, 4},		
//         }
// //  
//         B := [][]int{
//             {5, 6, 0, 0, 5, 6, 0, 0},
//             {7, 8, 0, 0, 7, 8, 0, 0},
//             {0, 0, 5, 6, 0, 0, 5, 6},
//             {0, 0, 7, 8, 0, 0, 7, 8},
//             {5, 6, 0, 0, 5, 6, 0, 0},
//             {7, 8, 0, 0, 7, 8, 0, 0},
//             {0, 0, 5, 6, 0, 0, 5, 6},
//             {0, 0, 7, 8, 0, 0, 7, 8},
//         }

//         C_regular := [][]int{
//             {0, 0, 0, 0, 0, 0, 0, 0},
//             {0, 0, 0, 0, 0, 0, 0, 0},
//             {0, 0, 0, 0, 0, 0, 0, 0},
//             {0, 0, 0, 0, 0, 0, 0, 0},
//             {0, 0, 0, 0, 0, 0, 0, 0},
//             {0, 0, 0, 0, 0, 0, 0, 0},
//             {0, 0, 0, 0, 0, 0, 0, 0},
//             {0, 0, 0, 0, 0, 0, 0, 0},
//         }

        // Perform Strassen matrix multiplication
        start_strassen := time.Now()
        C_strassen := strassen(A, B)
        t_strassen := time.Now()
        elapsed_strassen := t_strassen.Sub(start_strassen)

        // Perform regular matrix multiplication
        start_regular := time.Now()
        C_ans := matrixMult(A,B,C_regular,n)
        t_regular := time.Now()
        elapsed_regular := t_regular.Sub(start_regular)

	    // Checking equal matrices
	    if !compareResults(C_ans, C_strassen) {
	    	defer func() {
	    		if r := recover(); r != nil {
	    			log.Fatalf("Error: The matrices are not equal: %v", r)
        	    	log.Println("Strassen and regular matrices are different!")

	    		}
	    	}()
	    	log.Println("STRASSEN AND REGULAR DIFFERENT!")
	    }else{
            // Print the result matrix
            // fmt.Println("Result of Strassen matrix multiplication:")
            // for _, row := range C_strassen {
            //     fmt.Println(row)
            // }
            // fmt.Println("Elapsed Time: ", elapsed_strassen)

            // fmt.Println("Result of Regular matrix multiplication:")
            // for _, row := range C_ans {
            //     fmt.Println(row)
            // }
            // fmt.Println("Elapsed Time: ", elapsed_regular)    
            fmt.Println("N: ", n, " | Strassen Time: ", elapsed_strassen, " | Regular Time: ", elapsed_regular)
            // fmt.Println(n)
        }
    }    

}
	// TESTING MATRIX SUMMATION
	// A11 := [][]int{
		// {1, 2},
		// {3, 4},
	// }
	// A12 := [][]int{
		// {5, 6},
		// {7, 8},
	// }
	// C := matrixAdd(A11, A12)
    // for _, row := range C {
        // fmt.Println(row)
    // } 

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

	// TESTING MATRIX AGGREGATION
	// Example submatrices
	// A11 := [][]int{
		// {1, 2},
		// {3, 4},
	// }
	// A12 := [][]int{
		// {5, 6},
		// {7, 8},
	// }
	// A21 := [][]int{
		// {9, 10},
		// {11, 12},
	// }
	// A22 := [][]int{
		// {13, 14},
		// {15, 16},
	// }
// 
	// Combine the submatrices into one large matrix
	// result := combineMatrix(A11, A12, A21, A22)
// 
	// Print the result
	// fmt.Println("Combined matrix:")
	// for _, row := range result {
		// fmt.Println(row)
	// }

	// strassen testing
    // A := [][]int{
    //     {1, 2, 0, 0, 1, 2, 0, 0},
    //     {3, 4, 0, 0, 3, 4, 0, 0},
    //     {0, 0, 1, 2, 0, 0, 1, 2},
    //     {0, 0, 3, 4, 0, 0, 3, 4},
    //     {1, 2, 0, 0, 1, 2, 0, 0},
    //     {3, 4, 0, 0, 3, 4, 0, 0},
    //     {0, 0, 1, 2, 0, 0, 1, 2},
    //     {0, 0, 3, 4, 0, 0, 3, 4},		
    // }

    // B := [][]int{
    //     {5, 6, 0, 0, 5, 6, 0, 0},
    //     {7, 8, 0, 0, 7, 8, 0, 0},
    //     {0, 0, 5, 6, 0, 0, 5, 6},
    //     {0, 0, 7, 8, 0, 0, 7, 8},
    //     {5, 6, 0, 0, 5, 6, 0, 0},
    //     {7, 8, 0, 0, 7, 8, 0, 0},
    //     {0, 0, 5, 6, 0, 0, 5, 6},
    //     {0, 0, 7, 8, 0, 0, 7, 8},
    // }
	
	// Result of Strassen matrix multiplication:
	// [38 44 0 0 38 44 0 0]
	// [86 100 0 0 86 100 0 0]
	// [0 0 38 44 0 0 38 44]
	// [0 0 86 100 0 0 86 100]
	// [38 44 0 0 38 44 0 0]
	// [86 100 0 0 86 100 0 0]
	// [0 0 38 44 0 0 38 44]
	// [0 0 86 100 0 0 86 100]	