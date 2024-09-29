package pretty_test

import (
	"testing"

	"github.com/David-Rushton/pretty"
)

func Test(t *testing.T) {

	// what if we provide too many?

	t.Log(pretty.Sprintf("some text"))
	t.Log(pretty.Sprintf("some text", pretty.Red))
	t.Log(pretty.Sprintf("some text", pretty.BackgroundWhite))
	t.Log(pretty.Sprintf("some text", pretty.BackgroundWhite, pretty.Red))
	t.Log(pretty.Sprintf("some text", pretty.Green, pretty.BackgroundWhite, pretty.SlowBlink))
	t.Log(pretty.Sprintf("some text", pretty.Green, pretty.BackgroundWhite, pretty.RapidBlink))
	t.Log(pretty.Sprintf("some text", pretty.Red, pretty.BackgroundWhite, pretty.Bold))
	t.Log(pretty.Sprintf("some text", pretty.Red, pretty.BackgroundWhite, pretty.Italic))
	t.Log(pretty.Sprintf("some text", pretty.Red, pretty.BackgroundWhite, pretty.Bold, pretty.Italic))
	t.Log(pretty.Sprintf("some text", pretty.Red, pretty.BackgroundWhite, pretty.Strikethrough))

	t.Log(pretty.Sprintf("some text", pretty.WithRgb(0, 172, 0)))
	t.Log(pretty.Sprintf("some text", pretty.WithBackgroundRgb(0, 172, 0)))

	t.Log(pretty.Sprintf("some link", pretty.WithLink("https://example.com")))
	t.Log(pretty.Sprintf("some link", pretty.Colour(pretty.Blue), pretty.WithLink("https://example.com")))
	t.Log(pretty.Sprintf("some link", pretty.Colour(pretty.Blue), pretty.WithLink("https://news.bbc.com")))
}
