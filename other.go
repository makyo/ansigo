package ansigo

import (
	"regexp"
	"strings"
)

var (
	wsRE    = regexp.MustCompile(`\s`)
	dashRE  = regexp.MustCompile(`-+`)
	underRE = regexp.MustCompile(`_+`)
)

type other func(string) string

type Other interface {
	Apply(string) string
}

func (o other) Apply(s string) string {
	return o(s)
}

type others map[string]other

func (o others) Find(what string) (Other, error) {
	if attr, ok := Others[strings.ToLower(what)]; ok {
		return Other(attr), nil
	}
	return nil, CodeNotFound
}

var (
	AllCaps other = func(s string) string {
		return strings.ToUpper(s)
	}
	TitleCase other = func(s string) string {
		var builder strings.Builder
		capNext := true
		for _, r := range s {
			chr := string(r)
			if wsRE.MatchString(chr) {
				capNext = true
			} else if capNext {
				capNext = false
				chr = strings.ToUpper(chr)
			}
			builder.WriteString(chr)
		}
		return builder.String()
	}
	CamelCase other = func(s string) string {
		var builder strings.Builder
		capNext := false
		started := false
		for _, r := range s {
			chr := string(r)
			if !wsRE.MatchString(chr) {
				started = true
				if capNext {
					capNext = false
					chr = strings.ToUpper(chr)
				}
			} else {
				if started {
					capNext = true
				}
				continue
			}
			builder.WriteString(chr)
		}
		return builder.String()
	}
	UpperCamelCase other = func(s string) string {
		var builder strings.Builder
		capNext := true
		for _, r := range s {
			chr := string(r)
			if wsRE.MatchString(chr) {
				capNext = true
				continue
			} else if capNext {
				capNext = false
				chr = strings.ToUpper(chr)
			}
			builder.WriteString(chr)
		}
		return builder.String()
	}
	SnakeCase other = func(s string) string {
		return underRE.ReplaceAllString(wsRE.ReplaceAllString(s, "_"), "_")
	}
	KebabCase other = func(s string) string {
		return dashRE.ReplaceAllString(wsRE.ReplaceAllString(s, "-"), "-")
	}
)

var Others others = map[string]other{
	"allcaps":        AllCaps,
	"titlecase":      TitleCase,
	"camelcase":      CamelCase,
	"uppercamelcase": UpperCamelCase,
	"snakecase":      SnakeCase,
	"kebabcase":      KebabCase,
}
