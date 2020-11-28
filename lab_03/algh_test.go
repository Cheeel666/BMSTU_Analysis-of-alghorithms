package main

import "testing"

func BenchmarkBubbleSort(b *testing.B) {
	intSlice := generateSlice(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(intSlice)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	intSlice := generateSlice(100000)
	for i := 0; i < b.N; i++ {
		InsertionSort(intSlice)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	intSlice := generateSlice(100000)
	for i := 0; i < b.N; i++ {
		QuickSort(intSlice)
	}
}
