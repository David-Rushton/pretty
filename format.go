package pretty

type prettyFormat int

const (
	Bold          prettyFormat = 1
	Dim           prettyFormat = 2
	Italic        prettyFormat = 3
	Underline     prettyFormat = 4
	SlowBlink     prettyFormat = 5
	RapidBlink    prettyFormat = 6
	Strikethrough prettyFormat = 9
)

func WithBold() prettyFormat {
	return Bold
}

func WithDim() prettyFormat {
	return Dim
}

func WithItalic() prettyFormat {
	return Italic
}

func WithUnderline() prettyFormat {
	return Underline
}

func WithSlowBlink() prettyFormat {
	return SlowBlink
}

func WithRapidBlink() prettyFormat {
	return RapidBlink
}

func WithStrikethrough() prettyFormat {
	return Strikethrough
}
