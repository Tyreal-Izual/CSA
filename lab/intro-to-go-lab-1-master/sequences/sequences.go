package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	//var sli []int
	//sli := make(map[int]int)
	//sli := [] int{}
	sli := make([]int, 0, len(slice))
	for i, e := range slice {
		sli[i] = f(e)
	}
	//return sli
	fmt.Println("sli", sli)
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	//var arr [3]int
	arr := [3]int{}
	for i, e := range array {
		arr[i] = f(e)
	}
	return arr
	//fmt.Println("arr", arr)

}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intsSlice)

	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)
}
