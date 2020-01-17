//nolint: gochecknoglobals
package types_test

import (
	"strconv"
	"testing"

	"github.com/KevinAnthony/go-tree/types"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	a  = types.NewInt(1)
	b  = types.NewInt(2)
	b2 = types.NewInt(2)
)

func TestIntData_Equals(t *testing.T) {
	Convey("Equals", t, func() {
		Convey("should return false when A != B", func() {
			So(a.Equals(b), ShouldBeFalse)
		})
		Convey("should return true when B == B", func() {
			So(b.Equals(b), ShouldBeTrue)
			So(b.Equals(b2), ShouldBeTrue)
		})
	})
}

func TestIntData_GreaterThan(t *testing.T) {
	Convey("GreaterThen", t, func() {
		Convey("should return false when A < B", func() {
			So(a.GreaterThan(b), ShouldBeFalse)
		})
		Convey("should return true when B > A", func() {
			So(b.GreaterThan(a), ShouldBeTrue)
		})
		Convey("should return false when B <= B", func() {
			So(b.GreaterThan(b), ShouldBeFalse)
			So(b.GreaterThan(b2), ShouldBeFalse)
		})
	})
}

func TestIntData_GreaterThanOrEqual(t *testing.T) {
	Convey("GreaterThenOrEqual", t, func() {
		Convey("should return false when A < B", func() {
			So(a.GreaterThanOrEqual(b), ShouldBeFalse)
		})
		Convey("should return true when B > A", func() {
			So(b.GreaterThanOrEqual(a), ShouldBeTrue)
		})
		Convey("should return false when B <= B", func() {
			So(b.GreaterThanOrEqual(b), ShouldBeTrue)
			So(b.GreaterThanOrEqual(b2), ShouldBeTrue)
		})
	})
}

func TestIntData_LessThan(t *testing.T) {
	Convey("LessThan", t, func() {
		Convey("should return false when A < B", func() {
			So(a.LessThan(b), ShouldBeTrue)
		})
		Convey("should return true when B > A", func() {
			So(b.LessThan(a), ShouldBeFalse)
		})
		Convey("should return false when B <= B", func() {
			So(b.LessThan(b), ShouldBeFalse)
			So(b.LessThan(b2), ShouldBeFalse)
		})
	})
}

func TestIntData_LessThanOrEqual(t *testing.T) {
	Convey("LessThanOrEqual", t, func() {
		Convey("should return false when A < B", func() {
			So(a.LessThanOrEqual(b), ShouldBeTrue)
		})
		Convey("should return true when B > A", func() {
			So(b.LessThanOrEqual(a), ShouldBeFalse)
		})
		Convey("should return false when B <= B", func() {
			So(b.LessThanOrEqual(b), ShouldBeTrue)
			So(b.LessThanOrEqual(b2), ShouldBeTrue)
		})
	})
}

func TestNewInt(t *testing.T) {
	Convey("NewInt", t, func() {
		Convey("should return value with data set", func() {
			i := types.NewInt(77)
			So(i.Get(), ShouldResemble, 77)
		})
	})
}

func TestIntData_String(t *testing.T) {
	Convey("String", t, func() {
		Convey("should convert it to string", func() {
			i := types.NewInt(8675309)
			s := i.String()
			i2, err := strconv.Atoi(s)
			So(err, ShouldBeNil)
			So(i.Get(), ShouldResemble, i2)
		})
	})
}

func TestIntData_Get(t *testing.T) {
	Convey("Get", t, func() {
		Convey("should return value as an int", func() {
			i := types.NewInt(77)
			So(i.Get(), ShouldResemble, 77)
		})
	})
}
