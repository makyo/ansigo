package ansigo

import (
	"fmt"
	"strings"
)

// attribute describes an attribute that a string of runes might have.
// (e.g: bold, undelrine, etc)
type attribute struct {
	// ANSI codes usually have one code to turn a functionality on, and
	// another to turn it off. These are labeled start and end here as they
	// often show up in pairs like this rather than being left on. The
	// universal end, of course, is 0.
	start, end uint8
}

// Attribute describes a string-modifying code that's not a color.
type Attribute interface {
	Start() string
	End() string
	Apply(string) string
}

// Reset is the ANSI code to reset all attributes and colors of a string.
const Reset string = "\x1b[0m"

// Start prints the start ANSI escape code for the attribute.
func (a attribute) Start() string {
	return fmt.Sprintf("\x1b[%dm", a.start)
}

// End prints the ANSI escape code to turn off the attribute's functionality.
func (a attribute) End() string {
	return fmt.Sprintf("\x1b[%dm", a.end)
}

// Apply turns on the attribute for the given string by surrounding it with
// the start and end codes.
func (a attribute) Apply(s string) string {
	return fmt.Sprintf("%s%s%s", a.Start(), s, a.End())
}

var (
	Bold                    attribute = attribute{start: 1, end: 22}
	Faint                             = attribute{start: 2, end: 22}
	Italic                            = attribute{start: 3, end: 23}
	Underline                         = attribute{start: 4, end: 24}
	Blink                             = attribute{start: 5, end: 25}
	Flash                             = attribute{start: 6, end: 25}
	Reverse                           = attribute{start: 7, end: 27}
	Conceal                           = attribute{start: 8, end: 28}
	CrossedOut                        = attribute{start: 9, end: 29}
	AltFont1                          = attribute{start: 11, end: 10}
	AltFont2                          = attribute{start: 12, end: 10}
	AltFont3                          = attribute{start: 13, end: 10}
	AltFont4                          = attribute{start: 14, end: 10}
	AltFont5                          = attribute{start: 15, end: 10}
	AltFont6                          = attribute{start: 16, end: 10}
	AltFont7                          = attribute{start: 17, end: 10}
	AltFont8                          = attribute{start: 18, end: 10}
	AltFont9                          = attribute{start: 19, end: 10}
	Fraktur                           = attribute{start: 20, end: 23}
	DoubleUnderline                   = attribute{start: 21, end: 24}
	Framed                            = attribute{start: 51, end: 54}
	Encircled                         = attribute{start: 52, end: 54}
	Overlined                         = attribute{start: 53, end: 55}
	IdeogramUnderline                 = attribute{start: 60, end: 65}
	IdeogramDoubleUnderline           = attribute{start: 61, end: 65}
	IdeogramOverline                  = attribute{start: 62, end: 65}
	IdeogramDoubleOverline            = attribute{start: 63, end: 65}
	IdeogramStressMarking             = attribute{start: 64, end: 65}
)

// attributes maps an attribute to its name.
type attributes map[string]attribute

// Find attempts to find an attribute by its name.
func (a attributes) Find(what string) (Attribute, error) {
	if attr, ok := a[strings.ToLower(what)]; ok {
		return Attribute(attr), nil
	}
	return nil, CodeNotFound
}

// Attributes is the list of available attributes.
var Attributes attributes = map[string]attribute{
	"bold":                    Bold,
	"faint":                   Faint,
	"italic":                  Italic,
	"underline":               Underline,
	"blink":                   Blink,
	"flash":                   Flash,
	"reverse":                 Reverse,
	"conceal":                 Conceal,
	"crossedout":              CrossedOut,
	"altfont1":                AltFont1,
	"altfont2":                AltFont2,
	"altfont3":                AltFont3,
	"altfont4":                AltFont4,
	"altfont5":                AltFont5,
	"altfont6":                AltFont6,
	"altfont7":                AltFont7,
	"altfont8":                AltFont8,
	"altfont9":                AltFont9,
	"fraktur":                 Fraktur,
	"doubleunderline":         DoubleUnderline,
	"framed":                  Framed,
	"encircled":               Encircled,
	"overlined":               Overlined,
	"ideogramunderline":       IdeogramUnderline,
	"ideogramdoubleunderline": IdeogramDoubleUnderline,
	"ideogramoverline":        IdeogramOverline,
	"ideogramdoubleoverline":  IdeogramDoubleOverline,
	"ideogramstressmarking":   IdeogramStressMarking,
}
