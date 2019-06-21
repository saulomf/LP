package main

import (
	"github.com/hajimehoshi/ebiten"
        "image/color"
)

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Eliminada a funcionalidade de escolher mais de uma tela
	//grid := <-tela.canal // Recebe ponteiro para a tela	

	print_background(screen,fundo)
	
	print_cells(screen,tela.output)

	return nil
}

func print_background(screen *ebiten.Image , background *ebiten.Image) {
    	// Coloca terreno
	screen.Fill(color.RGBA{0, 0x0f, 0, 0xff})
	screen.DrawImage(background, nil)
}

func print_cells(screen *ebiten.Image , grid *vetOutput){
	// Obtem configurações
	op := &ebiten.DrawImageOptions{}

	// offset reset
	x_reset := float64(-(jumpx*grid.J)) 

	op.GeoM.Reset()
	op.GeoM.Translate(bordax,borday)
	for i := 0; i < grid.I; i++ {
		for j := 0; j < grid.J; j++{
			if(grid.array[i][j].imagem != nil) {
          			screen.DrawImage(grid.array[i][j].imagem, op)
			}
    			op.GeoM.Translate(jumpx,0.0)
		}
		op.GeoM.Translate(x_reset,0.0)
		op.GeoM.Translate(0.0,jumpy)
	}
}