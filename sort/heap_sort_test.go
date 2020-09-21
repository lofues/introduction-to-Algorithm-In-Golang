package sort

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	Convey("建大顶堆", t, func() {
		nums := []int{1, 2, 3, 4, 5}
		res := MakeBigHeap(nums)
		fmt.Println(res)
	})

	Convey("建小顶堆", t, func() {
		nums := []int{5, 4, 3, 2, 1}
		res := MakeSmallHeap(nums)
		fmt.Println(res)
	})

	Convey("大顶堆排序", t, func() {
		nums := []int{5, 4, 2, 3, 1, 1, 0}
		HeapSort(nums)
		So(sort.IsSorted(sort.IntSlice(nums)), ShouldBeTrue)
	})

}
