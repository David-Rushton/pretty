package pretty

type Rgb struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type BackgroundRgb Rgb

func WithRgb(r, g, b uint8) Rgb {
	return Rgb{r, g, b}
}

func WithBackgroundRgb(r, g, b uint8) BackgroundRgb {
	return BackgroundRgb{r, g, b}
}
