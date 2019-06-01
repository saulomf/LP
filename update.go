package main

import (
	"github.com/hajimehoshi/ebiten"
        "image/color"
)

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	grid := <-tela.canal // Recebe ponteiro para a tela

	// Obtem configurações
	op := &ebiten.DrawImageOptions{}

    	// Coloca terreno
	screen.Fill(color.RGBA{0, 0x0f, 0, 0xff})
	screen.DrawImage(fundo, nil)
	
	// Coloca output (tentarei otimizar dps)
	x_reset := float64(-6.0*grid.J) 
	
        op.GeoM.Reset()
	op.GeoM.Translate(1.0,1.0)
	for i := 0; i < grid.I; i++ {
		for j := 0; j < grid.J; j++{
			if(grid.array[i][j].imagem != nil) {
          			screen.DrawImage(grid.array[i][j].imagem, op)
			}
    			op.GeoM.Translate(6.0,0.0)
		}
		op.GeoM.Translate(x_reset,0.0)
		op.GeoM.Translate(0.0,6.0)
	}
	return nil
}

