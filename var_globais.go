package main

import (
    	"time"
    	"math/rand"
      _ "image/png"
      	"github.com/hajimehoshi/ebiten"
)


var clear map[string]func() //create a map for storing clear funcs

var blu *ebiten.Image //Imagem 1 quadrado azul
var red *ebiten.Image //Imagem 2 quadrado vermelho
var fundo *ebiten.Image //Imagem 3 de fundo

var tela struct {
	//canal chan vetOutput 
	//ready chan bool
	output *vetOutput
}

var s = rand.NewSource(time.Now().UnixNano())
var r = rand.New(s)

//var catalog = [2]Rule{rule_conway(),rule_conway()}



const (
	spacex = 3.0
	spacey = 3.0

	bordax = 1.0
	borday = 1.0

	cellsx = 100
	cellsy = 100
)

const (
	jumpx = spacex + bordax
	jumpy = spacey + borday 

	outputx = bordax + cellsx*jumpx
	outputy = borday + cellsy*jumpy
)