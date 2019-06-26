package main

import (
	
	"image/color"
	_ "image/png"
	//"fmt"

	"github.com/hajimehoshi/ebiten"
)

/*func update(screen *ebiten.Image) error {

	Atualiza_agua <- struct{}{}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Trava acesso para os grids - no inicio para otimizar print
	Water_lock.Lock()
	Ground_lock.Lock()
	
	// Fill with solid colors
	screen.Fill(color.NRGBA{0xff, 0xff, 0xff,0xff})
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	for i:=0; i< cell_num_y; i++ {
		for j:=0; j< cell_num_x; j++ {
			altitude := relevo[i][j]
	
			// Impedir alteracoes no meio da leitura
			coluna_de_agua := agua[i][j]	
			
			// Printar quadrado verde
			op.ColorM.Scale(0, 0, 0, 0)
			rT,gT,bT,alphaT := Scaled_RGBA(&c_grass,altitude,1)
			op.ColorM.Translate(rT, gT, bT, alphaT)
			screen.DrawImage(square, op)
			
			if(coluna_de_agua > 0.00) {
				op.ColorM.Scale(0, 0, 0, 0)
				rA,gA,bA,alphaA := Scaled_RGBA(&c_water,altitude_max,0.95)
				op.ColorM.Translate(rA, gA, bA, alphaA)
				screen.DrawImage(square, op)
			} 

			op.GeoM.Translate(cell_size_x,0)
		}
		op.GeoM.Translate(-screen_x,0)
		op.GeoM.Translate(0,cell_size_y)
	}
	
	// Libera acesso para os grids
	Water_lock.Unlock()
	Ground_lock.Unlock()

	return nil
}

func Scaled_RGBA(cor *color.RGBA,altitude,densidade float64) (float64,float64,float64,float64) {
	// Set color
	r := float64(cor.R)/0xff
	r = 0.2+(r-0.2)*(altitude/altitude_max)

	g := float64(cor.G)/0xff
	g = 0.2+(g-0.2)*(altitude/altitude_max)

	b := float64(cor.B)/0xff
	b = 0.2+(b-0.2)*(altitude/altitude_max)

	alpha := float64(cor.A)/0xff
	alpha = 0.2 + (alpha-0.2)*densidade

	return r,g,b,alpha
}*/

func Queue() {
	
}

func Recieve() {

}

func update(screen *ebiten.Image) error {

	Queue()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	Recieve(screen)

	return nil
}

func Recieve (screen *ebiten.Image) {
	


}