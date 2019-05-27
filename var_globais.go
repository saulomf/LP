package main

import (
    	"time"
    	"math/rand"
      _ "image/png"
      "github.com/hajimehoshi/ebiten"
)


var clear map[string]func() //create a map for storing clear funcs
var img *ebiten.Image //Imagem 1 quadrado preto
var img2 *ebiten.Image //Imagem 2 quadrado branco
var tela [][]Celula // canal para grid de print
var teste int = 0 // variavel para teste de controle do grid
var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)
