package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	Convey("长度小与1时数据不变", t, func() {
		params := []int{}
		QuickSort(params, 0, len(params)-1)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})

	Convey("乱序变为正序", t, func() {
		params := []int{5, 4, 3, 2, 1}
		QuickSort(params, 0, len(params)-1)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})

	Convey("完全相同序列", t, func() {
		params := []int{1, 1, 1, 1}
		QuickSort(params, 0, len(params)-1)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})
}

func TestQuickSortTailRecur(t *testing.T) {
	Convey("长度小与1时数据不变", t, func() {
		params := []int{}
		QuickSortTailRecur(params, 0, len(params)-1)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})

	Convey("乱序变为正序", t, func() {
		params := []int{5, 4, 3, 2, 1}
		QuickSortTailRecur(params, 0, len(params)-1)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})

	Convey("完全相同序列", t, func() {
		params := []int{1, 1, 1, 1}
		QuickSortTailRecur(params, 0, len(params)-1)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})
}

func TestQuickSortRandom(t *testing.T) {
	Convey("长度小与1时数据不变", t, func() {
		params := []int{}
		QuickSortRandom(params, 0, len(params)-1)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})

	Convey("乱序变为正序", t, func() {
		params := []int{5, 4, 3, 2, 1}
		QuickSortRandom(params, 0, len(params)-1)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})

	Convey("完全相同序列", t, func() {
		params := []int{1, 1, 1, 1}
		QuickSortRandom(params, 0, len(params)-1)
		So(sort.IsSorted(sort.IntSlice(params)), ShouldEqual, true)
	})
}
