package main 

import "log"

//原始数组arr[lr]范围上的累加和
func getSum(sum []int,l int,r int) int{
	if l == 0 {
		return sum[r]
	}else{
		//l ！= 0
		return sum[r] - sum[l-1]
	}
}

func main() {
	/*
		假设有一个数组arr,用户总是频繁的查询arr中
		范围的累加和，你该如何组织数据，能让这种查询变得每次都变得特别快捷
	*/
	arr := []int{5, 2, 8, 9, 6, -5}
	log.Print(arr)
	sum := preSumArray(arr)

	l := 1
	r := 3
	log.Println("数组范围[",l,"...",r,"]上累加和：" ,getSum(sum,l,r))

	l = 0
	r = 2
	log.Println("数组范围[",l,"...",r,"]上累加和：" ,getSum(sum,l,r))
	
}

func preSumArray(arr []int) []int{
	n := len(arr)
	sum := make([]int,n)
	sum[0] = arr[0]
	for i := 1 ; i < n ; i++{
		sum[i] = sum[i-1] + arr[i]
	}
	return sum
}

