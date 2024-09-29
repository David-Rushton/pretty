package pretty

// Values align with ANSI escape codes and standard terminal colours.
// https://en.wikipedia.org/wiki/ANSI_escape_code#3-bit_and_4-bit
type prettyColour int
type prettyBackgroundColour int

const (
	Black prettyColour = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BrightGray prettyColour = iota + 90
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

const (
	BackgroundBlack prettyBackgroundColour = iota + 40
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
)

const (
	BackgroundBrightGray prettyBackgroundColour = iota + 100
	BackgroundBrightRed
	BackgroundBrightGreen
	BackgroundBrightYellow
	BackgroundBrightBlue
	BackgroundBrightMagenta
	BackgroundBrightCyan
	BackgroundBrightWhite
)

func Colour(colour prettyColour) prettyColour {
	return colour
}

func BackgroundColour(colour prettyBackgroundColour) prettyBackgroundColour {
	return colour
}
