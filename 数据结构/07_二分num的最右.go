package main

import "log"

func main() {
	arr := []int{1,2,3,3,4,5,6,9}
	search := 5
	log.Println("数组为：",arr)
	log.Println("<=",search,"的最右位置的下表标为：",lessEqualMostRight(arr,search))
}

func lessEqualMostRight(arr []int ,num int)int{
	if arr ==nil || len(arr) < 1{
		return -1
	}
	l := 0
	r := len(arr) - 1
	m := 0
	ans := -1
	for l < r{
		m = l + ( r - l ) >> 1
		if arr[m] <= num{
			ans = m
			l = m+1
		} else{
			r = m-1
		}
	}
	return ans
}