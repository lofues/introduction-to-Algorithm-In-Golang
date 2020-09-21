package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReversePair(t *testing.T) {
	Convey("长度小与1时数据不变", t, func() {
		nums := []int{}
		count := ReversePair(nums)
		So(count, ShouldBeZeroValue)
		nums = []int{1}
		count = ReversePair(nums)
		So(count, ShouldBeZeroValue)
	})

	Convey("逆序数保持正确", t, func() {
		nums := []int{5, 4, 3, 2, 1}
		count := ReversePair(nums)
		So(count, ShouldEqual, 10)
	})
}
