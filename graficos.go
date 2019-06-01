package main

import (
    	"os"
    	"os/exec"
    	"runtime"
	
        _ "image/png"
        "log"

        "github.com/hajimehoshi/ebiten"
        "github.com/hajimehoshi/ebiten/ebitenutil"
)

func init() {
    	clear = make(map[string]func()) //Initialize it
    	clear["linux"] = func() {
    	    cmd := exec.Command("clear") //Linux example, its tested
    	    cmd.Stdout = os.Stdout
    	    cmd.Run()
    	}
    	clear["windows"] = func() {
    	    cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
    	    cmd.Stdout = os.Stdout
    	    cmd.Run()
    	}
    	//Parte Grafica
    	var err error
	blu, _, err = ebitenutil.NewImageFromFile("blue.png", ebiten.FilterDefault)
	red, _, err = ebitenutil.NewImageFromFile("pink.png", ebiten.FilterDefault)
	fundo, _, err = ebitenutil.NewImageFromFile("mapa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}