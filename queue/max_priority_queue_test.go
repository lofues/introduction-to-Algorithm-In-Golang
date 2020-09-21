package queue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	Convey("", t, func() {
		q := PriorityQueue{}
		q.Init()
		for _, num := range []int{3, 4, 5, 1, 2} {
			q.Insert(num)
		}
		So(q.Size, ShouldEqual, 5)
		So(q.Max(), ShouldEqual, 5)
		maxNum, err := q.Pop()
		So(err, ShouldBeNil)
		So(maxNum, ShouldEqual, 5)
		So(q.Length(), ShouldEqual, 4)
		maxNum, err = q.Pop()
		So(err, ShouldBeNil)
		So(maxNum, ShouldEqual, 4)
		q.Insert(7)
		So(q.Max(), ShouldEqual, 7)
		So(q.Length(), ShouldEqual, 4)
		maxNum, err = q.Pop()
		So(maxNum, ShouldEqual, 7)
		So(q.Length(), ShouldEqual, 3)
	})
}
