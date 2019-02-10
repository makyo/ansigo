package main

import (
	"fmt"

	a "github.com/makyo/ansigo"
)

func main() {
	fmt.Printf(
		"%s\t(reset)\n%s\t(reset)\n%s\t(reset)\n%s\t(reset)\n",
		a.MaybeApplyOne("bold", "Attributes"),
		a.MaybeApply("brightyellow", "Colors"),
		a.MaybeApply("DeepSkyBlue4:bg", "Backgrounds"),
		a.MaybeApply("#123456+rgb(255, 127, 0):bg+underline", "Combinations"))
}
