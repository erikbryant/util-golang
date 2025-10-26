package matrices

import (
	"fmt"
	"log"
	"strconv"

	"github.com/erikbryant/util-golang/common"
)

type Matrix[T common.Numbers] [][]T

// New returns a new, empty matrix of the given dimensions
func New[T common.Numbers](rows, cols int) Matrix[T] {
	if rows <= 0 || cols <= 0 {
		return nil
	}

	A := make(Matrix[T], rows)

	for row := 0; row < rows; row++ {
		A[row] = make([]T, cols)
	}

	return A
}

func (A Matrix[T]) Rows() int {
	return len(A)
}

func (A Matrix[T]) Cols() int {
	return len(A[0])
}

func (A Matrix[T]) Copy() Matrix[T] {
	rowsA := len(A)
	colsA := len(A[0])

	B := New[T](rowsA, colsA)

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			B[i][j] = A[i][j]
		}
	}

	return B
}

func (A Matrix[T]) Get(row, col int) T {
	return A[row][col]
}

// Print prints the matrix to the screen, f() [optional] formats row/col numbers to strings
func (A Matrix[T]) Print(title string, f func(int) string) {
	rowsA := len(A)
	colsA := len(A[0])

	if f == nil {
		f = func(i int) string {
			return strconv.Itoa(i)
		}
	}

	if title != "" {
		fmt.Printf("\n    %s\n", title)
	}

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
				fmt.Printf("%6v ", A[row][col]*100)
			}
		}
		fmt.Println()
	}
}

// SetRow sets all cells in that row to value
func (A Matrix[T]) SetRow(row int, value T) {
	colsA := len(A[0])
	for col := 0; col < colsA; col++ {
		A[row][col] = value
	}
}

// Set sets all cells to value
func (A Matrix[T]) Set(value T) {
	rowsA := len(A)
	colsA := len(A[0])

	for row := 0; row < rowsA; row++ {
		for col := 0; col < colsA; col++ {
			A[row][col] = value
		}
	}
}

// Mul multiplies AxB, putting the result in C
func (A Matrix[T]) Mul(B, C Matrix[T]) Matrix[T] {
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
