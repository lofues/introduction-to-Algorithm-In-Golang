package stack

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStack(t *testing.T) {
	Convey("", t, func() {
		s := Stack{}
		s.Init(1024)
		So(s.Length(), ShouldBeZeroValue)
		s.Insert(1)
		So(s.Length(), ShouldEqual, 1)
		s.Insert(2)
		s.Insert(3)
		So(s.Length(), ShouldEqual, 3)
		top, err := s.Top()
		So(err, ShouldBeNil)
		So(top, ShouldEqual, 3)
		top, err = s.Pop()
		So(err, ShouldBeNil)
		So(top, ShouldEqual, 3)
		So(s.Length(), ShouldEqual, 2)
		top, err = s.Pop()
		So(err, ShouldBeNil)
		So(s.Length(), ShouldEqual, 1)
		So(top, ShouldEqual, 2)
	})
}
