package ansigo_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	ansi "github.com/makyo/ansigo"
)

func TestANSI(t *testing.T) {
	Convey("When using ansigo", t, func() {

		Convey("When applying one ANSI code", func() {

			Convey("It can succeed", func() {
				s, err := ansi.ApplyOne("bold", "rose")
				So(err, ShouldBeNil)
				So(s, ShouldEqual, "\x1b[1mrose\x1b[22m")
				s, err = ansi.ApplyOne("black", "rose")
				So(err, ShouldBeNil)
				So(s, ShouldEqual, "\x1b[30mrose\x1b[39m")
				s, err = ansi.ApplyOne("DeepSkyBlue4", "rose")
				So(err, ShouldBeNil)
				So(s, ShouldEqual, "\x1b[38;5;23mrose\x1b[39m")
				s, err = ansi.ApplyOne("rgb(255, 128, 1)", "rose")
				So(err, ShouldBeNil)
				So(s, ShouldEqual, "\x1b[38;2;255;128;1mrose\x1b[39m")
			})

			Convey("But it can fail", func() {
				s, err := ansi.ApplyOne("bald", "rose")
				So(err, ShouldEqual, ansi.CodeNotFound)
				So(s, ShouldEqual, "rose")
			})

			Convey("Or one can just blithely go ahead anyway!", func() {
				So(ansi.MaybeApplyOne("bald", "rose"), ShouldEqual, "rose")
			})

			Convey("It accepts mods for colors", func() {
				So(ansi.MaybeApplyOne("black:bg", "rose"), ShouldEqual, "\x1b[40mrose\x1b[49m")

				Convey("But defaults to foreground", func() {
					So(ansi.MaybeApplyOne("black", "rose"), ShouldEqual, "\x1b[30mrose\x1b[39m")
				})
			})
		})

		Convey("When applying multiple ANSI codes", func() {

			Convey("It can succeed", func() {
				s, err := ansi.Apply("bold+black", "rose")
				So(err, ShouldBeNil)
				So(s, ShouldEqual, "\x1b[30m\x1b[1mrose\x1b[22m\x1b[39m")
			})

			Convey("But it can fail", func() {
				s, err := ansi.Apply("bald", "rose")
				So(err, ShouldEqual, ansi.CodeNotFound)
				So(s, ShouldEqual, "rose")
			})

			Convey("Or one can just blithely go ahead anyway!", func() {
				So(ansi.MaybeApply("bald", "rose"), ShouldEqual, "rose")
			})
		})
	})
}
