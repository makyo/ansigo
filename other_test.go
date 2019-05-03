package ansigo_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	ansi "github.com/makyo/ansigo"
)

func TestOther(t *testing.T) {
	Convey("When using other modifications", t, func() {

		Convey("It should be able to...", func() {

			Convey("...create all caps", func() {
				So(ansi.AllCaps.Apply("rose"), ShouldEqual, "ROSE")
			})

			Convey("...apply title case", func() {
				So(ansi.TitleCase.Apply("rose tyler"), ShouldEqual, "Rose Tyler")
				So(ansi.TitleCase.Apply("the  doctor"), ShouldEqual, "The  Doctor")
			})

			Convey("...apply camel case", func() {
				So(ansi.CamelCase.Apply("rose tyler"), ShouldEqual, "roseTyler")
				So(ansi.CamelCase.Apply("the  doctor"), ShouldEqual, "theDoctor")
			})

			Convey("...apply upper camel case", func() {
				So(ansi.UpperCamelCase.Apply("rose tyler"), ShouldEqual, "RoseTyler")
				So(ansi.UpperCamelCase.Apply("the  doctor"), ShouldEqual, "TheDoctor")
			})

			Convey("...apply snake case", func() {
				So(ansi.SnakeCase.Apply("rose tyler"), ShouldEqual, "rose_tyler")
				So(ansi.SnakeCase.Apply("the  doctor"), ShouldEqual, "the_doctor")
			})

			Convey("...apply kebab case", func() {
				So(ansi.KebabCase.Apply("rose tyler"), ShouldEqual, "rose-tyler")
				So(ansi.KebabCase.Apply("the  doctor"), ShouldEqual, "the-doctor")
			})
		})
	})
}
