package renderer

import (
	"github.com/Depado/bfchroma"
	bf "github.com/russross/blackfriday/v2"
)

var (
	defaultOptions = []bfchroma.Option{
		bfchroma.EmbedCSS(),
	}
	defaultRenderer = bfchroma.NewRenderer(defaultOptions...)
)

// Render returns an HTML version of the given Markdown string.
func Render(md string) string {
	return string(bf.Run([]byte(md), bf.WithRenderer(defaultRenderer)))
}
