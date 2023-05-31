package main

import "log"

func main() {
	arr := []int{10, 3, 4, 2, 4, -4, 2, 2}
	log.Println(arr)
	insertionSort(arr)
	log.Println(arr)
}

func insertionSort(arr []int) {
	if arr == nil || len(arr) < 2{
		return 
	}
	n := len(arr)
	for i := 1;i < n; i++{
		for j := i-1; j >=0 && arr[j] > arr[j+1]; j--{
			swap(arr, j, j+1)
		} 
	} 
}

func swap(arr []int, i int, j int){
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
} 