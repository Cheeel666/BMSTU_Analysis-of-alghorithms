package main
import(
  "fmt"
  "math/rand"
)


func generateSlice(size int) []int {

    slice := make([]int, size, size)
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}

func BubbleSort(numbers []int) []int{
  for i:= len(numbers); i > 0; i--{
    for j := 1; j < i; j++{
      if numbers[j-1] > numbers[j]{
        t := numbers[j]
        numbers[j] = numbers[j-1]
        numbers[j-1] = t
      }
    }
  }
  return numbers
}

func InsertionSort(arr []int) []int {
    len := len(arr)
    for i := 1; i < len; i++ {
        for j := 0; j < i; j++ {
            if arr[j] > arr[i] {
                arr[j], arr[i] = arr[i], arr[j]
            }
        }
    }
    return arr
}

func QuickSort(a []int) []int {
    if len(a) < 2 {
        return a
    }

    left, right := 0, len(a)-1

    pivot := rand.Int() % len(a)

    a[pivot], a[right] = a[right], a[pivot]

    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }

    a[left], a[right] = a[right], a[left]

    QuickSort(a[:left])
    QuickSort(a[left+1:])

    return a
}
func main(){
  slice := generateSlice(10)
  //slice := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
  fmt.Println("Standart slice: ", slice)
  fmt.Println("Sorted with bubble sort: ", BubbleSort(slice))
  fmt.Println("Sorted with insertion sort: ", InsertionSort(slice))
  fmt.Println("Sorted with bubble sort: ", QuickSort(slice))
}
