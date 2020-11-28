package main

import (
	"fmt"
)

func createMatrix(row, col int) [][]int {
	matrix := make([][]int, row)
	for i := range matrix {
		matrix[i] = make([]int, col)
	}
	return matrix
}

func StandartMult(m1, m2 [][]int) [][]int {
	if (len(m1) == 0 || len(m2) == 0) || (len(m1[0]) != len(m2)) {
		return nil
	}
	res := createMatrix(len(m1), len(m2[0]))
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}

func VinogradMult(m1, m2 [][]int) [][]int {

	if (len(m1) == 0 || len(m2) == 0) || (len(m1[0]) != len(m2)) {
		return nil
	}
	res := createMatrix(len(m1), len(m2[0]))
	rowF := make([]int, len(m1))
	colF := make([]int, len(m2[0]))

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1[0])/2; j++ {
			rowF[i] += m1[i][j*2] * m1[i][j*2+1]
		}
	}

	for i := 0; i < len(m2[0]); i++ {
		for j := 0; j < len(m2)/2; j++ {
			colF[i] += m2[j*2][i] * m2[j*2+1][i]
		}
	}

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			res[i][j] = -rowF[i] - colF[j]
			for k := 0; k < len(m1[0])/2; k++ {
				res[i][j] += (m1[i][2*k+1] + m2[2*k][j]) * (m1[i][2*k] + m2[2*k+1][j])
			}
		}
	}

	if len(m1[0])%2 == 1 {
		for i := 0; i < len(m1); i++ {
			for j := 0; j < len(m2[0]); j++ {
				res[i][j] += m1[i][len(m1[0])-1] * m2[len(m1[0])-1][j]
			}
		}
	}
	return res
}

func VinOptimMult(matrix1 [][]int, matrix2 [][]int) [][]int {
	var n1 int = len(matrix1)
	var n2 int = len(matrix2)

	if n1 == 0 || n2 == 0 {
		return nil
	}

	var m1 int = len(matrix1[0])
	var m2 int = len(matrix2[0])

	if m1 != n2 {
		return nil
	}

	mulH := make([]int, n1)
	mulV := make([]int, m2)
	result := createMatrix(n1, m2)

	var m1Mod2 int = m1 % 2
	var n2Mod2 int = n2 % 2

	for i := 0; i < n1; i++ {
		for j := 0; j < m1-m1Mod2; j += 2 {
			mulH[i] += matrix1[i][j] * matrix1[i][j+1]
		}
	}

	for i := 0; i < m2; i++ {
		for j := 0; j < n2-n2Mod2; j += 2 {
			mulV[i] += matrix2[j][i] * matrix2[j+1][i]
		}
	}

	var buff int
	for i := 0; i < n1; i++ {
		for j := 0; j < m2; j++ {
			buff = -mulH[i] - mulV[j]
			for k := 0; k < m1-m1Mod2; k += 2 {
				buff += (matrix1[i][k+1] + matrix2[k][j]) * (matrix1[i][k] + matrix2[k+1][j])
			}
			result[i][j] = buff
		}
	}

	if m1Mod2 == 1 {
		var m1Min1 int = m1 - 1
		for i := 0; i < n1; i++ {
			for j := 0; j < m2; j++ {
				result[i][j] += matrix1[i][m1Min1] * matrix2[m1Min1][j]
			}
		}
	}

	return result
}

func main() {
	matrix1 := [][]int{{1, 2, 3}, {1,2,3}, {1,2,3}}
	matrix2 := [][]int{{1, 2, 3}, {1, 2, 3},{1, 2, 3}}
	matrixStd := StandartMult(matrix1, matrix2)
	matrixVin := VinogradMult(matrix1, matrix2)
	matrixVinOpt := VinOptimMult(matrix1, matrix2)
	fmt.Println(matrixStd)
	fmt.Println(matrixVin)
	fmt.Println(matrixVinOpt)

}
