package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTheNthNumber(t *testing.T) {
	Convey("测试第3小的数", t, func() {
		nums := []int{4, 3, 1, 2, 5}
		So(TheNthNumber(nums, 0, len(nums)-1, 1), ShouldEqual, 1)
	})
}
