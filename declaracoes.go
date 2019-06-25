package main

import (
        _ "image/png"
        "log"
        "github.com/hajimehoshi/ebiten"
        "github.com/hajimehoshi/ebiten/ebitenutil"
)


var fundo *ebiten.Image
var img *ebiten.Image //Imagem 1 quadrado preto
var img2 *ebiten.Image //Imagem 2 quadrado branco
var img3 *ebiten.Image
var grid = make([][]Celula,200,200) // grid global
var copia_grid = make([][]Celula,200,200)
var teste int = 0 // variavel para teste de controle do grid
//type Rule [512]bool

func init() {
    //Parte Grafica
    var err error
    img, _, err = ebitenutil.NewImageFromFile("verde.png", ebiten.FilterDefault)
    img2, _, err = ebitenutil.NewImageFromFile("pink.png", ebiten.FilterDefault)
    img3, _, err = ebitenutil.NewImageFromFile("laranja.png", ebiten.FilterDefault)
    fundo, _, err = ebitenutil.NewImageFromFile("mapa.png", ebiten.FilterDefault)
    if err != nil {
        log.Fatal(err)
    }

}
