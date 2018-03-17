package main

import "fmt"

func BinarySearch(array []int ,target int) int {
	left :=0
	right := len(array) - 1
	//mid := (left + right)/2
	mid := left + (right - left)/2

	for left <= right {
		//如果查找的目标恰好是mid,那么查找完成
		if target == array[mid] {
			return mid
		}

		//target比中间的小，在左半部分查找
		if target < array[mid] {
			right = mid - 1
		}

		//target比中间的大,在右半部分查找
		if target > array[mid] {
			left = left + 1
		}

	}

	//如果在数组中找不到，返回-1
	return -1
}

func main(){
	var a = []int{1,2,3,4,5}
	rst := BinarySearch(a,3)
	fmt.Println(rst)
}



