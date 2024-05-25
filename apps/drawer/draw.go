package drawer

import (
	"image"
	"image/png"
	"os"

	drawer_parser "github.com/azeek21/blog/apps/drawer/parser"
)

type Drawer struct {
	Colors *drawer_parser.ColorMap
	Src    *drawer_parser.ImageSource
	Dst    *image.RGBA
}

func NewDrawer() Drawer {
	return Drawer{
		Colors: drawer_parser.NewColorMap(),
		Src:    drawer_parser.NewImageSource(),
		Dst:    nil,
	}
}

func (d *Drawer) DrawPngFromFile(in, out string) error {

	if err := d.ParseFile(in); err != nil {
		return err
	}
	d.Dst = image.NewRGBA(image.Rect(0, 0, d.Src.W, d.Src.H))
	for y, row := range d.Src.Bytes {
		for x, pixel := range row {
			d.Dst.Set(x, y, (*d.Colors)[string(pixel)])
		}
	}
	outFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer outFile.Close()
	png.Encode(outFile, d.Dst)
	return nil
}
