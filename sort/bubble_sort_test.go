package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	Convey("长度小与1时数据不变", t, func() {
		params := []int{}
		BubbleSort(params)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})

	Convey("乱序变为正序", t, func() {
		params := []int{3, 2, 5, 1, 4}
		BubbleSort(params)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})

	Convey("完全相同序列", t, func() {
		params := []int{1, 1, 1, 1}
		BubbleSort(params)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})
}
