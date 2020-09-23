package linkedlist

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	Convey("", t, func() {
		l := LinkedList{}
		So(l.Empty(), ShouldBeTrue)
		So(l.Length(), ShouldBeZeroValue)
		l.Insert(1)
		l.Insert(2)
		l.Insert(3)
		So(l.Length(), ShouldEqual, 3)
		So(l.Empty(), ShouldBeFalse)
		So(l.Head.Val, ShouldEqual, 1)
		l.Delete(4)
		So(l.Length(), ShouldEqual, 3)
		l.Delete(3)
		So(l.Length(), ShouldEqual, 2)
		l.Delete(2)
		So(l.Empty(), ShouldBeFalse)
		So(l.Length(), ShouldEqual, 1)
		l.Delete(1)
		So(l.Length(), ShouldEqual, 0)
		So(l.Empty(), ShouldBeTrue)
	})
}
