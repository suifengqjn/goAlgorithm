package bubble

import (
	"fmt"
	"myProject/goAlgorithm/util"
)

// 普通的冒泡排序
func CommonBubble()  {

	arr := util.GetArr(5, 50)
	fmt.Println(arr)

	l := len(arr)
	for i:= 0;i < l-1;i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

   fmt.Println(arr)
}


// 冒泡优化1
func FlagBubble()  {
	arr := util.GetArr(10, 50)
	fmt.Println(arr)
	l := len(arr)

	var flag bool
	for i:= 0;i < l-1;i++ {
		flag = false
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if flag == false{
			break
		}
	}

	fmt.Println(arr)
}

// 冒泡优化2
func EndBubble()  {
	arr := util.GetArr(20, 50)
	fmt.Println(arr)

	l := len(arr)
	end := 0
	for i:= 0;i < l-1;i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				end = j
			}
			if j == end {
				continue
			}
		}
	}

	fmt.Println(arr)
}

func SharkBubble()  {

	arr := util.GetArr(10, 50)
	l := len(arr)
	left := 0
	right := l - 1

	// left  以左已有序
	// right 以右已有序
	// 两个区间游标相遇则集合有序
	for left < right {

		//左半部分
		for i:=left;i<right;i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--


		//右半部分
		for j:=right;j > left;j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
		left ++
	}


	fmt.Println(arr)


}