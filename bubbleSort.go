//go 1.6.2

package main

import (
	. "fmt" //Silent import
	"math/rand"
	"time"
)

func main() {
	a := generateArray()
	sortTime(a, defaultBubbleSort)
}
func generateArray() []int {
	var array []int
	for i := 0; i < 999999999; i++ {
		array = append(array, rand.Intn(99999999))
	}

	return array
}

//Calculate sorting time
func sortTime(array []int, f func(a []int) []int) {

	t0 := time.Now()
	b := f(array)
	t1 := time.Now()
	Println(b)
	var t = int(t1.Sub(t0).Seconds())

	Printf("Elapsed time: %v secs", t)

}

//Bubble sort with optimization
func optimizedBubbleSort(array []int) []int {
	for i := 1; i <= len(array); i++ {
		sorted := true
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				sorted = false
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
		if sorted == true {
			break
		}
	}
	return array
}

//Bubble sort without optimization
func defaultBubbleSort(array []int) []int {
	for i := 1; i <= len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
