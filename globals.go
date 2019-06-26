package water_flow

import (
	
	"image/color"
	_ "image/png"
	"sync"

	"github.com/hajimehoshi/ebiten"
)


var (
	square, _ = ebiten.NewImage(cell_size_x,  cell_size_y, ebiten.FilterDefault)
	c_grass = color.RGBA{127, 200, 127, 255}
	c_rock  = color.RGBA{127, 127, 127, 255}
	c_water = color.RGBA{63 , 127, 200, 255}

	relevo [cell_num_y][cell_num_x]float64	// Altitude
	agua [cell_num_y][cell_num_x]float64	// Tamanho da coluna
)

var output_lock sync.Mutex

const (
	altitude_max = 20.0

	cell_size_x = 4
	cell_size_y = 4

	cell_num_x = 150
	cell_num_y = 100

	screen_x = cell_size_x*cell_num_x
	screen_y = cell_size_y*cell_num_y
)


var output = make(chan Output)

const (

)