package main

import "fmt"

package water_flow

import (
	
	"image/color"
	_ "image/png"
	"log"
	"math"
	//"fmt"

	"github.com/hajimehoshi/ebiten"
)



func main() {	
	go Run_Simulation(6,1,1)

	if err := ebiten.Run(update, screen_x, screen_y, 2, "Colors"); err != nil {
		log.Fatal(err)
	}
}