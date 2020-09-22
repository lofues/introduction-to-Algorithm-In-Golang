package queue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQueue(t *testing.T) {
	Convey("", t, func() {
		q := Queue{}
		q.Init(3)
		So(q.Length(), ShouldBeZeroValue)
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		So(q.Length(), ShouldEqual, 3)
		top, err := q.Top()
		So(err, ShouldBeNil)
		So(top, ShouldEqual, 1)
		top, err = q.Dequeue()
		So(top, ShouldEqual, 1)
		So(err, ShouldBeNil)
		So(q.Length(), ShouldEqual, 2)
		top, err = q.Dequeue()
		So(top, ShouldEqual, 2)
		So(err, ShouldBeNil)
		So(q.Length(), ShouldEqual, 1)
	})
}
