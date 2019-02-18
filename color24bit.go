package ansigo

import (
	"fmt"
	"strconv"
)

type color24bit rgb

// FGStart returns the ANSI code to start writing runes in this color
// as the foreground.
func (c color24bit) FGStart() string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

// BGStart returns the ANSI code to start writing runes in this color
// as the background.
func (c color24bit) BGStart() string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c.R, c.G, c.B)
}

// FG turns the text for the provided string the specified color by surrounding
// it with the start and end codes.
func (c color24bit) FG(s string) string {
	return fmt.Sprintf("%s%s%s", c.FGStart(), s, FGEnd)
}

// BG turns the background for the provided string the specified color by
// surrounding it with the start and end codes.
func (c color24bit) BG(s string) string {
	return fmt.Sprintf("%s%s%s", c.BGStart(), s, BGEnd)
}

// BGWithReset turns the background for the provided string the specified color
// by surrounding it with the start and reset codes.
func (c color24bit) BGWithReset(s string) string {
	return fmt.Sprintf("%s%s%s", c.BGStart(), s, Reset)
}

// FGWithReset turns the text for the provided string the specified color by
// surrounding it with the start and reset codes.
func (c color24bit) FGWithReset(s string) string {
	return fmt.Sprintf("%s%s%s", c.FGStart(), s, Reset)
}

// Apply applies the color to the string with the given modifier ("bg" for
// background, "fg" for foreground (or default)).
func (c color24bit) Apply(s, mod string) string {
	if mod == "bg" {
		return c.BG(s)
	}
	return c.FG(s)
}

// ApplyWithReset applies the color to the string with the given modifier ("bg"
// for background, "fg" for foreground (or default)). It terminates with a
// Reset instead of a color-off.
func (c color24bit) ApplyWithReset(s, mod string) string {
	if mod == "bg" {
		return c.BGWithReset(s)
	}
	return c.FGWithReset(s)
}

type colors24bit struct{}

func (c colors24bit) Find(what string) (Color, error) {
	if matches := hexRegexp.FindAllStringSubmatch(what, -1); len(matches) == 1 && len(matches[0]) == 4 {
		r, _ := strconv.ParseUint(matches[0][1], 16, 8)
		g, _ := strconv.ParseUint(matches[0][2], 16, 8)
		b, _ := strconv.ParseUint(matches[0][3], 16, 8)
		return color24bit{R: uint8(r), G: uint8(g), B: uint8(b)}, nil
	} else if matches := rgbRegexp.FindAllStringSubmatch(what, -1); len(matches) == 1 && len(matches[0]) == 4 {
		var r, g, b int
		r, _ = strconv.Atoi(matches[0][1])
		g, _ = strconv.Atoi(matches[0][2])
		b, _ = strconv.Atoi(matches[0][3])
		if r > 255 || g > 255 || b > 255 {
			return nil, InvalidColorSpec
		}
		return color24bit{R: uint8(r), G: uint8(g), B: uint8(b)}, nil
	} else if matches := hslRegexp.FindAllStringSubmatch(what, -1); len(matches) == 1 && len(matches[0]) == 4 {
		var h, s, l float64
		h, _ = strconv.ParseFloat(matches[0][1], 32)
		s, _ = strconv.ParseFloat(matches[0][2], 32)
		l, _ = strconv.ParseFloat(matches[0][3], 32)
		if h > 360 || s > 100 || l > 100 {
			return nil, InvalidColorSpec
		}
		r, g, b := decodeHSL(h, s, l)
		return color24bit{R: r, G: g, B: b}, nil
	} else {
		return nil, CodeNotFound
	}
}

var Colors24bit colors24bit
