package dayone

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows     int
	Columns  int
	Elements [][]int
}

// Method to get the number of rows
func (m *Matrix) GetRows() int {
	return m.Rows
}

// Method to get the number of columns
func (m *Matrix) GetColumns() int {
	return m.Columns
}

// Method to set the elements of the matrix at a given position (i, j)
func (m *Matrix) SetElement(i, j, value int) {
	if i >= 0 && i < m.Rows && j >= 0 && j < m.Columns {
		m.Elements[i][j] = value
	}
}

// Method to add two matrices
func (m *Matrix) AddMatrix(other Matrix) Matrix {
	result := Matrix{
		Rows:     m.Rows,
		Columns:  m.Columns,
		Elements: make([][]int, m.Rows),
	}
	for i := range result.Elements {
		result.Elements[i] = make([]int, m.Columns)
		for j := range result.Elements[i] {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result
}

// Method to print matrix structure as JSON
func (m *Matrix) ToJSON() string {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func Qa() { //main function
	// Example usage
	matrix1 := Matrix{
		Rows:    2,
		Columns: 2,
		Elements: [][]int{
			{1, 2},
			{3, 4},
		},
	}

	matrix2 := Matrix{
		Rows:    2,
		Columns: 2,
		Elements: [][]int{
			{5, 6},
			{7, 8},
		},
	}

	matrix1.SetElement(0, 1, 10)
	fmt.Println("Matrix 1:", matrix1.ToJSON())

	sumMatrix := matrix1.AddMatrix(matrix2)
	fmt.Println("Sum Matrix:", sumMatrix.ToJSON())

}
