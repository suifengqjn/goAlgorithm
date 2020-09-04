package sort

import (
	"fmt"
	"myProject/goAlgorithm/util"
	"testing"
)

var arr []int

func init() {
	arr = util.GetArr(6, 20)
}

func TestInsert(t *testing.T) {
	fmt.Println(arr)
	fmt.Println(Insert(arr))
}
