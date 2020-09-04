package insert

import (
	"fmt"
	"myProject/goAlgorithm/util"
)

func Insert()  {
	arr := util.GetArr(5, 20)
	fmt.Println("[UNSORTED]:\t", arr)

	l := len(arr)

	for i := 1; i < l; i++ {
		//向前寻找合适位置插入
		for j := i; j > 0;j-- {
			if arr[j-1] > arr[j] {
				arr[j-1],arr[j] = arr[j], arr[j-1]
			}
			fmt.Println("[DEBUG]: ", arr)
		}

	}

	fmt.Println("[SORTED]:\t",arr)

}

