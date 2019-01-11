//Package matrix gives you functions and types for working with matrices
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

//Matrix represents a 2D matrix
type Matrix struct {
	value [][]int
	rows  int
	cols  int
}

//New creates a new instance of matrix
func New(input string) (matrix *Matrix, err error) {
	matrix = new(Matrix)
	rows := strings.Split(input, "\n")
	matrix.value = make([][]int, len(rows))

	for row, line := range rows {
		matrix.rows++
		cols := strings.Split(strings.Trim(line, " "), " ")
		for col, item := range cols {
			if matrix.cols == 0 {
				matrix.cols = len(cols)
			}
			if matrix.cols != len(cols){
				return matrix, errors.New("non matching arrays")
			}
			val, err := strconv.Atoi(strings.Trim(item, " "))
			if err != nil {
				return matrix, err
			}
			if matrix.value[row] == nil {
				matrix.value[row] = make([]int, len(cols))
			}
			matrix.value[row][col] = val
		}
	}
	return
}

//Rows gets a copy of the rows of the matrix
func (matrix Matrix) Rows() (output [][]int) {
	output = make([][]int, len(matrix.value))
	for i := 0; i < matrix.rows; i++ {
		for j := 0; j < matrix.cols; j++ {
			if output[i] == nil {
				output[i] = make([]int, matrix.cols)
			}
			output[i][j] = matrix.value[i][j]
		}
	}
	return
}

//Cols gets a copy of the cols of the matrix
func (matrix Matrix) Cols() (output [][]int) {
	output = make([][]int, matrix.cols)
	for row := 0; row < matrix.rows; row++ {
		for col := 0; col < matrix.cols; col++ {
			if output[col] == nil {
				output[col] = make([]int, matrix.rows)
			}
			output[col][row] = matrix.value[row][col]
		}
	}
	return
}

//Set an item in the matrix
func (matrix Matrix) Set(row int, col int, val int) bool {
	if row >= matrix.rows || col >= matrix.cols || row < 0 || col < 0 {
		return false
	}
	matrix.value[row][col] = val
	return true
}