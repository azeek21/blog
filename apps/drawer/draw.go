package drawer

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func Draw(output string, w, h int) error {
	png_output := output // output image lives here

	img_source := image.NewRGBA(image.Rect(0, 0, w, h))
	colors := make(map[int]color.RGBA, 2)

	colors[0] = color.RGBA{0, 100, 0, 255}   // green
	colors[1] = color.RGBA{50, 205, 50, 255} // limegreen

	index_color := 0
	size_board := 8
	size_block := int(w / size_board)
	loc_x := 0

	for curr_x := 0; curr_x < size_board; curr_x++ {

		loc_y := 0
		for curr_y := 0; curr_y < size_board; curr_y++ {

			draw.Draw(img_source, image.Rect(loc_x, loc_y, loc_x+size_block, loc_y+size_block),
				&image.Uniform{colors[index_color]}, image.Point{}, draw.Src)

			loc_y += size_block
			index_color = 1 - index_color // toggle from 0 to 1 to 0 to 1 to ...
		}
		loc_x += size_block
		index_color = 1 - index_color // toggle from 0 to 1 to 0 to 1 to ...
	}
	myfile, err := os.Create(png_output)
	if err != nil {
		return err
	}
	defer myfile.Close()
	png.Encode(myfile, img_source) // ... save image
	return nil
}
