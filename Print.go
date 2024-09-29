// Pretty prints to the terminal using [ANSI escape codes].
//
// [ANSI escape codes]: https://en.wikipedia.org/wiki/ANSI_escape_code
package pretty

import (
	"fmt"
	"hash/fnv"
)

const (
	ansiEscapeCode = "\033"
	ansiReset      = "\033[0m"
)

// Applies the supplied formats and prints the result to standard out.
func Printf(a any, f ...any) {
	fmt.Print(Sprintf(a, f...))
}

// Applies the supplied format and prints the result, with trailing newline, to standard out.
func Println(a any, f ...any) {
	fmt.Println(Sprintf(a, f...))
}

// Applies the supplied formats and returns the result.
func Sprintf(a any, f ...any) string {
	var inBandSignal string
	var url Link

	for _, format := range f {
		switch t := format.(type) {
		case prettyFormat, prettyColour, prettyBackgroundColour:
			inBandSignal = fmt.Sprintf(
				"%v%v[%vm",
				inBandSignal,
				ansiEscapeCode,
				t)
		case Rgb:
			inBandSignal = fmt.Sprintf(
				"%v%v[38;2;%v;%v;%vm",
				inBandSignal,
				ansiEscapeCode,
				t.Red,
				t.Green,
				t.Blue)
		case BackgroundRgb:
			inBandSignal = fmt.Sprintf(
				"%v%v[48;2;%v;%v;%vm",
				inBandSignal,
				ansiEscapeCode,
				t.Red,
				t.Green,
				t.Blue)
		case Link:
			url = t
		}
	}

	result := fmt.Sprintf("%v%v%v", inBandSignal, a, ansiReset)

	if len(url) > 0 {
		result = fmt.Sprintf(
			"%v]8;id=%v;%v%v\\%v%v]8;;%v\\",
			ansiEscapeCode,
			hash(string(url)),
			url,
			ansiEscapeCode,
			result,
			ansiEscapeCode,
			ansiEscapeCode)
	}

	return result
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
