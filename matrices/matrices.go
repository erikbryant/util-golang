package matrices

import (
	"fmt"
	"log"
)

type Matrix [][]float64

// New returns a new, empty matrix of the given dimensions
func New(rows, cols int) Matrix {
	if rows <= 0 || cols <= 0 {
		return nil
	}

	A := make([][]float64, rows)

	for row := 0; row < rows; row++ {
		A[row] = make([]float64, cols)
	}

	return A
}

func (A Matrix) Rows() int {
	return len(A)
}

func (A Matrix) Cols() int {
	return len(A[0])
}

func (A Matrix) Copy() Matrix {
	rowsA := len(A)
	colsA := len(A[0])

	B := New(rowsA, colsA)

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			B[i][j] = A[i][j]
		}
	}

	return B
}

// Print prints the matrix to the screen, f() formats row/col numbers to strings
func (A Matrix) Print(title string, f func(int) string) {
	rowsA := len(A)
	colsA := len(A[0])

	fmt.Printf("\n    %s\n", title)

	fmt.Printf("    ")
	for col := 0; col < colsA; col++ {
		fmt.Printf("%6s ", f(col))
	}
	fmt.Println()

	for row := 0; row < rowsA; row++ {
		fmt.Printf("%3s  ", f(row))
		for col := 0; col < colsA; col++ {
			if A[row][col] == 0 {
				fmt.Printf("%6s ", " ·  ")
			} else {
				fmt.Printf("%6.2f ", A[row][col]*100)
			}
		}
		fmt.Println()
	}
}

// SetRow sets all cells in that row to value
func (A Matrix) SetRow(row int, value float64) {
	colsA := len(A[0])
	for col := 0; col < colsA; col++ {
		A[row][col] = value
	}
}

// Set sets all cells to value
func (A Matrix) Set(value float64) {
	rowsA := len(A)
	colsA := len(A[0])

	for row := 0; row < rowsA; row++ {
		for col := 0; col < colsA; col++ {
			A[row][col] = value
		}
	}
}

// Mul multiplies AxB, putting the result in C
func (A Matrix) Mul(B, C Matrix) Matrix {
	rowsA := len(A)
	colsA := len(A[0])
	rowsB := len(B)
	colsB := len(B[0])
	rowsC := len(C)
	colsC := len(C[0])

	if colsA != rowsB {
		log.Fatalf("invalid input dimensions: %dx%d and %dx%d", rowsA, colsA, rowsB, colsB)
	}

	if rowsC != rowsA || colsC != colsB {
		log.Fatalf("invalid output dimensions: %dx%d should be %dx%d", rowsC, colsC, rowsA, colsB)
	}

	C.Set(0.0)

	for row := 0; row < rowsA; row++ {
		for col := 0; col < colsB; col++ {
			for k := 0; k < colsA; k++ {
				C[row][col] += A[row][k] * B[k][col]
			}
		}
	}

	return C
}
