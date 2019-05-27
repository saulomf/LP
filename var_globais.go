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
var blank *ebiten.Image //Imagem 3 quadrado branco

var tela [][]Celula // canal para grid de print
var teste int = 0 // variavel para teste de controle do grid
var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

var catalog = [2]Rule{rule_conway(),rule_conway()}