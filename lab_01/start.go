package main

import (
	"fmt"
	"math"
)

func min(a, b int)int {
	if a > b{
		return b
	}
	return a
}
func max(a, b int) int{
	if a < b{
		return b
	}
	return a
}

func RecursiveLevenstain(a, b []rune) int {
	var cost int
	if len(a) == 0 || len(b) == 0 {
		return int(math.Abs(float64(len(a) - len(b))))
	}
	cost = 1
	if a[len(a)-1] == b[len(b)-1] {
		cost = 0
	}

	return min(min(RecursiveLevenstain(a, b[:len(b)-1])+1,
		RecursiveLevenstain(a[:len(a)-1], b)+1), RecursiveLevenstain(a[:len(a)-1], b[:len(b)-1])+cost)
}

func RecursiveLevenstainMatrix(a, b []rune, matrix [][]int) int {
	var cost int
	if len(a) == 0 || len(b) == 0 {
		matrix[len(a)][len(b)] = int(math.Abs(float64(len(a) - len(b))))
	}
	if matrix[len(a)][len(b)] != -1 {
		return matrix[len(a)][len(b)]
	}
	cost = 1
	if a[len(a)-1] == b[len(b)-1] {
		cost = 0
	}

	matrix[len(a)][len(b)] = min(min(RecursiveLevenstainMatrix(a, b[:len(b)-1], matrix)+1,
		RecursiveLevenstainMatrix(a[:len(a)-1], b, matrix)+1),
		RecursiveLevenstainMatrix(a[:len(a)-1], b[:len(b)-1], matrix)+cost)
	return matrix[len(a)][len(b)]
}

func MatrixLevenstain(a, b []rune, matrix [][]int) int {
	var cost int
	for i := 0; i < len(a)+1; i++ {
		for j := 0; j < len(b)+1; j++ {
			if i == 0 {
				matrix[i][j] = j
			} else if j == 0 && i > 0 {
				matrix[i][j] = i
			} else {
				cost = 1
				if a[i-1] == b[j-1] {
					cost = 0
				}
				matrix[i][j] = min(min(matrix[i][j-1]+1,
					matrix[i-1][j]+1),
					matrix[i-1][j-1]+cost)
			}
		}
	}
	return matrix[len(a)][len(b)]
}

func DamerauLevenstainMatrix(a, b []rune, matrix [][]int) int {
	var cost int
	for i := 0; i < len(a)+1; i++ {
		for j := 0; j < len(b)+1; j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = max(i, j)
			} else {
				cost = 1
				if a[i-1] == b[j-1] {
					cost = 0
				}
				matrix[i][j] = min(min(matrix[i][j-1]+1,
					matrix[i-1][j]+1),
					matrix[i-1][j-1]+cost)
				if j > 1 && i > 1 &&
					a[i-1] == b[j-2] &&
					a[i-2] == b[j-1] {
					matrix[i][j] = min(min(matrix[i][j], matrix[i-2][j-2]+1),matrix[i][j])
				}
			}
		}
	}
	return matrix[len(a)][len(b)]
}

func clearMatrix(matrix [][]int, strlenA int, strlenB int) [][]int {
	for i := 0; i < strlenA; i++ {
		for j := 0; j < strlenB; j++ {
			matrix[i][j] = -1
		}
	}
	return matrix
}

func print_matrix(matrix [][]int){
	fmt.Println("")
	fmt.Println("Матрица:")
	for i := range matrix{
		for j := range matrix[i]{
			fmt.Print(matrix[i][j]," ")
		}
		fmt.Println(" ")
	}
}



func main() {
	var choise int
	var strA string
	var strB string
	var distance int
	var strlenA, strlenB int
	choise = 1

	for choise > 0 {
		fmt.Println("1 - Нахождение расстояния Левенштейна рекурсивно")
		fmt.Println("2 - Нахождение расстояния Левенштейна рекурсивно с матрицей")
		fmt.Println("3 - Нахождение расстояния Левенштейна матрично")
		fmt.Println("4 - Нахождение расстояния Дамерау - Левенштейна матрично")
		fmt.Println("0 - Выход")
		fmt.Println("")
		fmt.Println("Введите действие:")
		fmt.Scanf("%d", &choise)

		if choise >= 5 {
			fmt.Println("Неверный ввод!")
			break
		}
		if choise == 0{
			break
		}
		fmt.Println("Введите первую строку:")
		fmt.Scanf("%s", &strA)
		fmt.Println("Введите вторую строку:")
		fmt.Scanf("%s", &strB)
		if strA == strB{
			choise = 6
		}
		strA := []rune(strA)
		strB := []rune(strB)

		strlenA = len(strA) + 1
		strlenB = len(strB) + 1
		matrix := make([][]int, strlenA)
		for i := range matrix {
			matrix[i] = make([]int, strlenB)
		}
		matrix = clearMatrix(matrix, strlenA, strlenB)
		switch choise {
		case 1:
			distance = RecursiveLevenstain(strA, strB)
		case 2:
			distance = RecursiveLevenstainMatrix(strA, strB, matrix)
		case 3:
			distance = MatrixLevenstain(strA, strB, matrix)
		case 4:
			distance = DamerauLevenstainMatrix(strA, strB, matrix)
		case 6:
			distance = 0
		}
		if choise > 1 && choise < 5 && distance != 0{
			print_matrix(matrix)
		}
		fmt.Println("Расстояние = ", distance)
	}
}
