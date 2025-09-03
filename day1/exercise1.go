package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct{
	Rows, Cols int
	Ele [][]int
}

func NewMatrix(rows, cols int) *Matrix{
	elements := make([][]int, rows)
	for i:= range elements{
		elements[i] = make([]int, cols)
	}
	return &Matrix{
		Rows : rows,
		Cols : cols,
		Ele : elements,
	}
}

func (m *Matrix) getRows() int{
	return m.Rows
}

func (m *Matrix) getCols() int{
	return m.Cols
}

func (m *Matrix) setElements(i, j, value int) error{
    if i<0 || i>=m.Rows || j<0 || j>=m.Cols{
		return fmt.Errorf("Index out of bound: (%d, %d)", i,j)
	}
	m.Ele[i][j] = value
	return nil
}

func (m *Matrix) Add(other *Matrix) (*Matrix, error){
    if m.Rows != other.Rows || m.Cols != other.Cols{
		fmt.Errorf("Matrix cannot be added due to different dimensions")
	}
	result := NewMatrix(m.Rows, m.Cols)
	for i:=0; i<m.Rows; i++{
		for j:=0; j<m.Cols; j++{
			result.Ele[i][j] = m.Ele[i][j] + other.Ele[i][j]
		}
	}
	return result, nil
}

// returning a json representation of the matrix

func(m *Matrix) ToJSON() (string, error){
    jsonBytes, err:= json.Marshal(m)
	if err != nil{
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return string(jsonBytes), nil
}

func(m *Matrix) PrintMatrix(){
	for i:= range m.Ele{
		fmt.Println(m.Ele[i])
	}
}

func main(){
   m1 := NewMatrix(2,2)
   m1.setElements(0, 0, 1)
   m1.setElements(0, 1, 2)
   m1.setElements(1, 0, 3)
   m1.setElements(1, 1, 4)

   m2 := NewMatrix(2,2)
   m2.setElements(0, 0, 5)
   m2.setElements(0, 1, 6)
   m2.setElements(1, 0, 7)
   m2.setElements(1, 1, 8)

   fmt.Println("Matrix 1:")
   m1.PrintMatrix()
   fmt.Printf("Rows: %d, Columns: %d\n\n", m1.getRows(), m1.getCols())

   fmt.Println("Matrix 2:")
   m2.PrintMatrix()
   fmt.Printf("Rows: %d, Columns: %d\n\n", m2.getRows(), m2.getCols())

   sum, err:= m1.Add(m2)
   if err!=nil{
	   fmt.Println("Error in adding matrices:", err)
	   return
   }
   fmt.Println("Matrix1 + Matrix2")
   sum.PrintMatrix()


   //Get json representation

   jsonResponse, err := m1.ToJSON()
   if err != nil{
	  fmt.Printf("Error in converting to json:", err)
	  return
   }
   fmt.Printf("Matrix 1 in JSON format:")
   fmt.Printf(jsonResponse)
}


