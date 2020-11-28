package main

import "testing"

func BenchmarkRecursiveLevenstain(b *testing.B) {
	strA := []rune("lololololo")
	strB := []rune("ololololol")
	for n := 0; n < b.N; n++ {
		RecursiveLevenstain(strA, strB)
	}
}

func BenchmarkRecursiveLevenstainMatrix(b *testing.B) {
	var strlenA, strlenB int
	strA := []rune("lololololo")
	strB := []rune("ololololol")
	strlenA = len(strA) + 1
	strlenB = len(strB) + 1
	matrix := make([][]int, strlenA)
	for i := range matrix {
		matrix[i] = make([]int, strlenB)
	}
	for n := 0; n < b.N; n++ {
		RecursiveLevenstainMatrix(strA, strB, matrix)
		clearMatrix(matrix, strlenA, strlenB)
	}
}

func BenchmarkMatrixLevenstain(b *testing.B) {
	var strlenA, strlenB int
	strA := []rune("lololololo")
	strB := []rune("ololololol")
	strlenA = len(strA) + 1
	strlenB = len(strB) + 1
	matrix := make([][]int, strlenA)
	for i := range matrix {
		matrix[i] = make([]int, strlenB)
	}
	for n := 0; n < b.N; n++ {
		MatrixLevenstain(strA, strB, matrix)
		clearMatrix(matrix, strlenA, strlenB)
	}
}

func BenchmarkDamerauLevenstainMatrix(b *testing.B) {
	var strlenA, strlenB int
	strA := []rune("lololololo")
	strB := []rune("ololololol")
	strlenA = len(strA) + 1
	strlenB = len(strB) + 1
	matrix := make([][]int, strlenA)
	for i := range matrix {
		matrix[i] = make([]int, strlenB)
	}
	for n := 0; n < b.N; n++ {
		DamerauLevenstainMatrix(strA, strB, matrix)
		clearMatrix(matrix, strlenA, strlenB)
	}
}
