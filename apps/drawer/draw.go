package drawer

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	drawer_parser "github.com/azeek21/blog/apps/drawer/parser"
)

//"image"
//	"image/color"
//	"image/draw"
//	"image/png"
//	"os"

func openSourceImage(input string) (*bufio.Reader, error) {
	fid, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(fid), nil
}

func DrawPngFromFile(in, out string) error {
	src, err := openSourceImage(in)

	if err != nil {
		return err
	}

	colorsMap := NewColorMap()
	for {
		line, _, err := src.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if bytes.HasPrefix(line, drawer_parser.KEYWORDS.COLORS_START) {
			if err := colorsMap.ReadTillEndAndParse(src); err != nil {
				return err
			}
			continue
		}

		if bytes.HasPrefix(line, drawer_parser.KEYWORDS.IMAGE_START) {
			for {
				line, _, err := src.ReadLine()
				if err != nil {
					if err == io.EOF {
						break
					}
					return err
				}

				if bytes.HasPrefix(line, drawer_parser.KEYWORDS.END) {
					break
				}
				log.Printf("IMG: %v\n", string(line))
			}
			continue
		}
	}
	return nil
}
