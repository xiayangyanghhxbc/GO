package main

//二分搜索

import "log"

func main() {
	//请保证arr有序
	arr := []int{1,2,2,5,6,7,8,9}
	search := 7
	log.Println("数组为：",arr)
	log.Println("搜索的数为：",search)
	log.Println("存在与否：",exist(arr,search))
}

func exist(arr []int,num int)bool{
	if arr == nil || len(arr) == 0{
		return false
	}
	l := 0//左边界
	r := len(arr) - 1//右边界
	m := 0//中点变量先准备好
	for l <= r{
		//m=(l+r)/2
		m = l + (r-l) >> 1//右移一位
		if arr[m] == num {
			return true
		}else if arr[m] > num {
			r = m - 1//有序
		}else{
			l = m + 1
		}
	}
	return false
}