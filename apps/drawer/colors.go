package drawer

import (
	"errors"
	"image/color"
	"strconv"
	"strings"
)

var BAD_RGBA_ERROR = errors.New("Error parsing RGBA color from input. e.g: x=255,255,0,10")

type ColorMap map[string]*color.RGBA

func NewColorMap() *ColorMap {
	return &ColorMap{}
}

func (cm *ColorMap) ParseAndAssign(str string) error {
	keyVal := strings.Split(str, "=")
	var err error

	if len(keyVal) != 2 {
		return BAD_RGBA_ERROR
	}

	(*cm)[keyVal[0]], err = ParseRgba(&keyVal[1])
	if err != nil {
		return err
	}

	return nil
}

func ParseRgba(rgba *string) (*color.RGBA, error) {
	res := &color.RGBA{A: 255}
	RGBA := strings.Split((*rgba), ",")

	if len(RGBA) != 3 && len(RGBA) != 4 {
		return nil, BAD_RGBA_ERROR
	}

	r, err := strconv.Atoi(RGBA[0])
	if err != nil || r > 255 || r < 0 {
		return res, BAD_RGBA_ERROR
	}
	res.R = uint8(r)

	g, err := strconv.Atoi(RGBA[1])
	if err != nil || g > 255 || g < 0 {
		return res, BAD_RGBA_ERROR
	}
	res.G = uint8(g)

	b, err := strconv.Atoi(RGBA[2])
	if err != nil || b > 255 || b < 0 {
		return res, BAD_RGBA_ERROR
	}
	res.B = uint8(b)

	if len(RGBA) == 4 {
		a, err := strconv.Atoi(RGBA[3])
		if err != nil || a > 255 || a < 0 {
			return res, BAD_RGBA_ERROR
		}
		res.A = uint8(a)
	}

	return res, nil
}
