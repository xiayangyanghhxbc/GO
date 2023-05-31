package main

import "log"

func main() {
	arr := []int{10, 3, 0, -1, 3, 4, 8, -9}
	log.Println(arr)
	bubbleSort(arr)
	log.Println(arr)
}

func bubbleSort(arr []int){
	if arr == nil || len(arr) < 2{
		return 
	}
	n := len(arr)
	for end := n-1; end > 0; end--{
		for i := 0; i < end; i++{
			if arr[i] > arr[i+1]{
				swap(arr, i, i+1)
			}
		} 
	} 
}

func swap(arr []int, i int,j int){
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}


