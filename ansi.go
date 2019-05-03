package ansigo

import (
	"errors"
	"strings"
)

// Formatter represents something which can apply formatting to a string.
type Formatter interface {
	Apply(string) string
	ApplyWithReset(string) string
}

// Collection represents something which can be used to find a formatter.
type Collection interface {
	Find(string) (Formatter, error)
}

// CodeNotFound is returned when an ANSI code is requested which does not exist.
var CodeNotFound error = errors.New("ANSI code not found")

func applyOne(spec, s string, withReset bool) (string, error) {
	if other, err := Others.Find(spec); err == nil {
		return other.Apply(s), nil
	}
	if attr, err := Attributes.Find(spec); err == nil {
		if withReset {
			return attr.ApplyWithReset(s), nil
		}
		return attr.Apply(s), nil
	}
	cols := strings.Split(spec, ":")
	var col, mod string
	if len(cols) == 2 {
		col, mod = cols[0], cols[1]
	} else {
		col = cols[0]
		mod = "fg"
	}
	if c, err := Colors8.Find(col); err == nil {
		if withReset {
			return c.ApplyWithReset(s, mod), nil
		}
		return c.Apply(s, mod), nil
	}
	if c, err := Colors256.Find(col); err == nil {
		if withReset {
			return c.ApplyWithReset(s, mod), nil
		}
		return c.Apply(s, mod), nil
	}
	if c, err := Colors24bit.Find(col); err == nil {
		if withReset {
			return c.ApplyWithReset(s, mod), nil
		}
		return c.Apply(s, mod), nil
	}
	return s, CodeNotFound
}

// ApplyOne applies one code to a string. If it fails, it returns an error.
func ApplyOne(spec, s string) (string, error) {
	return applyOne(spec, s, false)
}

// MaybeApplyOne attempts to apply a code; if it fails, it just returns the
// string.
func MaybeApplyOne(spec, s string) string {
	a, _ := ApplyOne(spec, s)
	return a
}

// Apply attempts to apply all of the codes requested to the string, separated
// by +. If any of them fail, it stops and returns an error.
func Apply(specs, s string) (string, error) {
	var err error
	for _, spec := range strings.Split(specs, "+") {
		s, err = ApplyOne(spec, s)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}

// MaybeApply attempts to apply all of the codes requested to the string,
// separated by +. If any of them fail, it ignores the failure and continues on.
func MaybeApply(specs, s string) string {
	for _, spec := range strings.Split(specs, "+") {
		s = MaybeApplyOne(spec, s)
	}
	return s
}

// ApplyOneWithReset applies one code to a string. If it fails, it returns an
// error.
func ApplyOneWithReset(spec, s string) (string, error) {
	return applyOne(spec, s, true)
}

// MaybeApplyOneWithReset attempts to apply a code; if it fails, it just
// returns the string.
func MaybeApplyOneWithReset(spec, s string) string {
	a, _ := ApplyOneWithReset(spec, s)
	return a
}

// ApplyWithReset attempts to apply all of the codes requested to the string,
// separated by +. If any of them fail, it stops and returns an error.
func ApplyWithReset(specs, s string) (string, error) {
	var err error
	for _, spec := range strings.Split(specs, "+") {
		s, err = ApplyOneWithReset(spec, s)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}

// MaybeApplyWithReset attempts to apply all of the codes requested to the string,
// separated by +. If any of them fail, it ignores the failure and continues on.
func MaybeApplyWithReset(specs, s string) string {
	for _, spec := range strings.Split(specs, "+") {
		s = MaybeApplyOneWithReset(spec, s)
	}
	return s
}
