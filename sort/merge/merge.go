package merge

import (
	"fmt"
	"myProject/goAlgorithm/util"
)

func Merge()  {
	arr := util.GetArr(8, 20)
	fmt.Println("[UNSORTED]:\t", arr)

	sort := mergeSort(arr)

	fmt.Println("[SORTED]:",sort)


}

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	fmt.Println("--", arr)
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])


	fmt.Println("left", left)

	fmt.Println("right", right)
	return merge(left, right)

}

func merge(left, right []int) []int  {

	fmt.Println("=====", left, right)
	// 遍历比较后合并两个数组
	res := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		// 数组提前排序完毕
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		// 比较更小的值追加到 res[] 中
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}