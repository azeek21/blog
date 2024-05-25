package drawer

import (
	"bufio"
	"bytes"
	"io"
	"os"

	drawer_parser "github.com/azeek21/blog/apps/drawer/parser"
)

func openSourceImage(input string) (*bufio.Reader, error) {
	fid, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(fid), nil
}

func (d *Drawer) ParseFile(filName string) error {
	src, err := openSourceImage(filName)

	if err != nil {
		return err
	}

	for {
		line, _, err := src.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if bytes.HasPrefix(line, drawer_parser.KEYWORDS.COLORS_START) {
			if err := d.Colors.ReadTillEndAndParse(src); err != nil {
				return err
			}
		}

		if bytes.HasPrefix(line, drawer_parser.KEYWORDS.IMAGE_START) {
			err := d.Src.ReadTillEndAndParse(src)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
