package ansigo_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	ansi "github.com/makyo/ansigo"
)

func TestANSIAtIndex(t *testing.T) {

	Convey("When testing for what ANSI codes are active at an index", t, func() {

		Convey("It handles strings with no codes", func() {
			So(len(ansi.ANSIAtIndex("Rose Tyler", 5)), ShouldEqual, 0)
		})

		Convey("It ignores bad indices", func() {
			So(len(ansi.ANSIAtIndex("bad-wolf", -1)), ShouldEqual, 0)
			So(len(ansi.ANSIAtIndex("bad-wolf", 100)), ShouldEqual, 0)

			Convey("As the smallest a code could be is 4 bytes, <5 is a bad index", func() {
				So(len(ansi.ANSIAtIndex("bad-wolf", 2)), ShouldEqual, 0)
			})
		})

		Convey("It handles attributes", func() {

			Convey("One attribute", func() {
				So(ansi.ANSIAtIndex(ansi.Bold.Apply("Rose Tyler"), 10),
					ShouldResemble,
					[]string{"\x1b[1m"})
				So(ansi.ANSIAtIndex("Rose Tyler"+ansi.Bold.Apply(" and the Doctor"), 5),
					ShouldResemble,
					[]string{})
				So(ansi.ANSIAtIndex("Rose Tyler"+ansi.Bold.Apply(" and the Doctor"), 15),
					ShouldResemble,
					[]string{"\x1b[1m"})
				So(ansi.ANSIAtIndex(ansi.Bold.Apply("Rose Tyler")+" and the Doctor", 25),
					ShouldResemble,
					[]string{})
				So(ansi.ANSIAtIndex(ansi.Bold.Apply("Rose Tyler")+" and the Doctor", 5),
					ShouldResemble,
					[]string{"\x1b[1m"})
			})

			Convey("Multiple attributes", func() {
				So(ansi.ANSIAtIndex(ansi.Bold.Apply(ansi.Italic.Apply("Rose Tyler")), 10),
					ShouldResemble,
					[]string{"\x1b[1m", "\x1b[3m"})
				So(ansi.ANSIAtIndex(ansi.Bold.Apply(ansi.Italic.Apply("Rose Tyler")+" and the Doctor"), 10),
					ShouldResemble,
					[]string{"\x1b[1m", "\x1b[3m"})
				So(ansi.ANSIAtIndex(ansi.Bold.Apply(ansi.Italic.Apply("Rose Tyler")+" and")+" the Doctor", 10),
					ShouldResemble,
					[]string{"\x1b[1m", "\x1b[3m"})
				So(ansi.ANSIAtIndex(ansi.Bold.Apply(ansi.Italic.Apply("Rose Tyler")+" and the Doctor"), 25),
					ShouldResemble,
					[]string{"\x1b[1m"})
				So(ansi.ANSIAtIndex(ansi.Bold.Apply(ansi.Italic.Apply("Rose Tyler")+" and")+" the Doctor", 25),
					ShouldResemble,
					[]string{"\x1b[1m"})
				So(ansi.ANSIAtIndex(ansi.Bold.Apply(ansi.Italic.Apply("Rose Tyler")+" and")+" the Doctor", 35),
					ShouldResemble,
					[]string{})
			})
		})

		Convey("It handles colors", func() {

			Convey("It handles foregrounds", func() {
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red", "Rose Tyler"), 10),
					ShouldResemble,
					[]string{"\x1b[31m"})
				So(ansi.ANSIAtIndex("Rose Tyler"+ansi.MaybeApply("red", " and the Doctor"), 5),
					ShouldResemble,
					[]string{})
				So(ansi.ANSIAtIndex("Rose Tyler"+ansi.MaybeApply("red", " and the Doctor"), 15),
					ShouldResemble,
					[]string{"\x1b[31m"})
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red", "Rose Tyler")+" and the Doctor", 25),
					ShouldResemble,
					[]string{})
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red", "Rose Tyler")+" and the Doctor", 5),
					ShouldResemble,
					[]string{"\x1b[31m"})
			})

			Convey("It handles backgrounds", func() {
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red:bg", "Rose Tyler"), 10),
					ShouldResemble,
					[]string{"\x1b[41m"})
				So(ansi.ANSIAtIndex("Rose Tyler"+ansi.MaybeApply("red:bg", " and the Doctor"), 5),
					ShouldResemble,
					[]string{})
				So(ansi.ANSIAtIndex("Rose Tyler"+ansi.MaybeApply("red:bg", " and the Doctor"), 15),
					ShouldResemble,
					[]string{"\x1b[41m"})
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red:bg", "Rose Tyler")+" and the Doctor", 25),
					ShouldResemble,
					[]string{})
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red:bg", "Rose Tyler")+" and the Doctor", 5),
					ShouldResemble,
					[]string{"\x1b[41m"})
			})

			Convey("It handles both", func() {
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("blue:bg", "Rose Tyler")), 15),
					ShouldResemble,
					[]string{"\x1b[31m", "\x1b[44m"})
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("blue:bg", "Rose Tyler")+" and the Doctor"), 10),
					ShouldResemble,
					[]string{"\x1b[31m", "\x1b[44m"})
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("blue:bg", "Rose Tyler")+" and")+" the Doctor", 10),
					ShouldResemble,
					[]string{"\x1b[31m", "\x1b[44m"})
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("blue:bg", "Rose Tyler")+" and the Doctor"), 25),
					ShouldResemble,
					[]string{"\x1b[31m"})
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("blue:bg", "Rose Tyler")+" and")+" the Doctor", 25),
					ShouldResemble,
					[]string{"\x1b[31m"})
				So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("blue:bg", "Rose Tyler")+" and")+" the Doctor", 35),
					ShouldResemble,
					[]string{})
			})
		})

		Convey("It handles attributes and colors", func() {
			So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("overlined", "Rose Tyler")), 15),
				ShouldResemble,
				[]string{"\x1b[31m", "\x1b[53m"})
			So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("overlined", "Rose Tyler")+" and the Doctor"), 10),
				ShouldResemble,
				[]string{"\x1b[31m", "\x1b[53m"})
			So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("overlined", "Rose Tyler")+" and")+" the Doctor", 10),
				ShouldResemble,
				[]string{"\x1b[31m", "\x1b[53m"})
			So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("overlined", "Rose Tyler")+" and the Doctor"), 25),
				ShouldResemble,
				[]string{"\x1b[31m"})
			So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("overlined", "Rose Tyler")+" and")+" the Doctor", 25),
				ShouldResemble,
				[]string{"\x1b[31m"})
			So(ansi.ANSIAtIndex(ansi.MaybeApply("red", ansi.MaybeApply("overlined", "Rose Tyler")+" and")+" the Doctor", 35),
				ShouldResemble,
				[]string{})
		})
	})
}
