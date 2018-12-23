package png

import (
	"image"
	"io"
)

func Decode(r io.Reader) (image.Image, error)
func DecodeConfig(r io.Reader) (image.Config, error)
func init() {
	const pngHeader = "\x89PNG\r\n\x1a\n"
	image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
}
