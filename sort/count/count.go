package count

import (
	"fmt"
	"myProject/goAlgorithm/util"
)

func CountSort()  {
	arr := util.GetArr(5, 20)
	fmt.Println("[UNSORTED]:\t", arr)

	max := getMax(arr)
	sortedArr := make([]int, len(arr))
	countsArr := make([]int, max+1) // max+1 是为了防止 countsArr[] 计数时溢出

	// 元素计数
	for _, v := range arr {
		countsArr[v]++
	}
	fmt.Println("[countsArr]:\t",countsArr)

	// 统计独特数字个数并累加
	for i := 1; i <= max; i++ {
		countsArr[i] += countsArr[i-1]
	}

	fmt.Println("[DEBUG countsArr]:\t", countsArr)

	// 让 arr 中每个元素找到其位置
	for _, v := range arr {
		fmt.Println(countsArr[v])
		sortedArr[countsArr[v]-1] = v
		//fmt.Print(countsArr[v]-1, " ")
		// 保证稳定性
		countsArr[v]--
	}

	sortedArr = append(sortedArr, )


	fmt.Println(sortedArr)


}

func getMax(arr []int) int  {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}