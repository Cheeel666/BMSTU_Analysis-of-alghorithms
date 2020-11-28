package main

import "testing"

func BenchmarkStandartMult(b *testing.B) {
	matrix1 := [][]int{{0, 1, 2}, {3, 4, 5}}
	matrix2 := [][]int{{0, 1}, {3, 4}, {6, 7}}
	for n := 0; n < b.N; n++ {
		StandartMult(matrix1, matrix2)
	}
}

func BenchmarkVinogradMult(b *testing.B) {
	matrix1 := [][]int{{0, 1, 2}, {3, 4, 5}}
	matrix2 := [][]int{{0, 1}, {3, 4}, {6, 7}}
	for n := 0; n < b.N; n++ {
		VinogradMult(matrix1, matrix2)
	}
}

func BenchmarkVinOptimMult(b *testing.B) {
	matrix1 := [][]int{{0, 1, 2}, {3, 4, 5}}
	matrix2 := [][]int{{0, 1}, {3, 4}, {6, 7}}
	for n := 0; n < b.N; n++ {
		VinOptimMult(matrix1, matrix2)
	}
}
