// Package ansigo provides control over ANSI escape codes for terminal apps.
//
// There are a few entry points to the library, but most commonly, one will
// just use <code>Apply</code> or <code>MaybeApply</code>, which accept a
// spec string and a string to format. <code>Apply</code> returns an error
// if any bit of the spec string failed to match and bails as soon as a
// failure occurs, while <code>MaybeApply</code> simply returns the string
// with all successful formatting applied (after all, it's just text, it'll
// still be text at the end).
//
// The spec string is simply a list of formatting steps to take separated
// by +. For instance, one could have bold green text with
// <code>bold+green</code>. Additionally, colors can be specified as
// background rather than foreground colors by appending <code>:bg</code>
// so we could have the previous on a blue background with
// <code>bold+green+blue:bg</code>.
//
// ansigo supports all SGR codes from the ANSI specification. This includes
// attributes such as bold and underline, as well as three different color
// spaces: 3/4-bit[0] (<code>Colors8</code>), 8-bit (<code>Colors256</code>),
// and 24-bit true-color (<code>Colors24bit</code>).
//
// For the list of which attributes and colors are available, as well as to see
// which your terminal supports,
//
//	go run _examples/capability_check.go
//
// <code>Attributes</code> and all three color spaces above, as implementers
// of <code>Collection</code>, implement a <code>Find</code> method which
// returns a <code>Formatter</code>, if one is found, and an error if one is
// not. For <code>Attributes</code> and <code>Colors8</code>, the search
// term is just a name ("bold", "green", etc), but <code>Colors256</code>
// and <code>Colors24bit</code>, you have more leeway. For the former,
// you can search by color name (though note that there are some duplicate
// names in there, which will lead to you getting the first match back),
// color ID[1], and the color's CSS-style hex code (e.g: "#ff0000"), rgb
// function (e.g: "rgb(255, 0, 0)"), and hsl function (e.g: "hsl(0,100%,50%)").
// However, all of those are strictly specified. If you search
// <code>Colors256</code> for, say, "#123512", you won't find it, despite that
// being a valid hex code. For that, use <code>Colors24bit</code>, which will
// succeed for any valid CSS hex/rgb/hsl function that uses whole numbers.
//
// That's a lot of choices, though, so it's usually better to just use
// <code>(Maybe)Apply</code> :)
//
// For a list of attributes and colors, see https://ansigo.projects.makyo.io .
//
// [0]: Despite the name, <code>Colors8</code> contains 16 colors, the 8
// original colors, and their "bright" variants: "green" + "brightgreen",
// etc.
//
// [1]: Which the author was personally quite fascinated with:
//
//     0-  7:  standard colors (as in ESC [ 30–37 m)
//     8- 15:  high intensity colors (as in ESC [ 90–97 m)
//    16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//   232-255:  grayscale from black to white in 24 steps
package ansigo
