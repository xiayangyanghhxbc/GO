package main

import "log"
//找到>=num的最左位置
func main(){
	arr := []int {2,2,2,3,6,6,7,8,9}
	search := 6
	log.Println("数组为：",arr)
	log.Println(">=",search,"的最左位置的下标为：",moreEqualMostLeft(arr,search))
}

func moreEqualMostLeft(arr []int ,val int)int{
	if arr == nil || len(arr) < 1{
		return -1
	}
	l := 0
	r := len(arr) - 1
	m := 0
	ans := -1
	for l <= r{
		m = l + ((r - l) >> 1)
		if arr[m] >= val{
			ans = m
			r = m-1
		}else{
			l = m+1
		}
	}
	return ans
}