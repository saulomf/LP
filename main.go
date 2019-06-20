package main

import (
    	"fmt"
    	"time"
    	"math/rand"
    	"os"
    	"os/exec"
    	"runtime"

        _ "image/png"
        "log"
        //"image/color"

        "github.com/hajimehoshi/ebiten"
        "github.com/hajimehoshi/ebiten/ebitenutil"
)


var clear map[string]func() //create a map for storing clear funcs
var fundo *ebiten.Image
var img *ebiten.Image //Imagem 1 quadrado preto
var img2 *ebiten.Image //Imagem 2 quadrado branco
var grid = make([][]Celula,200,200) // grid global
var teste int = 0 // variavel para teste de controle do grid
type Rule [512]bool


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
	img, _, err = ebitenutil.NewImageFromFile("black.png", ebiten.FilterDefault)
	img2, _, err = ebitenutil.NewImageFromFile("blank.png", ebiten.FilterDefault)
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

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	//screen.Fill(color.RGBA{0, 0xff, 0, 0xff})
    screen.DrawImage(fundo, nil)
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(1, 1)
	//screen.DrawImage(img, op))

	//x := 0.0
	for j := 0; j < 146; j++ {
		op.GeoM.Translate(grid[0][j].posX, grid[0][j].posY)
        op.GeoM.Scale(0.5, 0.5)
        if(grid[0][j].viva == true){
		          screen.DrawImage(img, op)
        }
        op.GeoM.Reset()
		//x-=11.0
	}
	//op.GeoM.Translate(x, 0)
	//x = 0.0

	for i := 0; i < 110; i++ {
		//op.GeoM.Translate(0, 11)
		//screen.DrawImage(img, op)
		for j := 0; j < 146; j++ {
    		op.GeoM.Translate(grid[i][j].posX, grid[i][j].posY)
            op.GeoM.Scale(0.5, 0.5)

            if(grid[i][j].viva == true){
                time_agora := time.Now()
                su := grid[i][j].time_vida
                diff := time_agora.Sub(su)

                diff1 := int(diff/1e9)
            //fmt.Println("agora: ", time_agora)
            //fmt.Println("celula: ", su)
                if(diff1 > 0){
                    fmt.Println("diff de tempo: ", diff1)
                }
                if(diff1 < 5){
                    screen.DrawImage(img, op)
                } else{
                    grid[i][j].viva = false
                }
            }

            op.GeoM.Reset()
			//x-=11.0
		}
		//op.GeoM.Translate(x, 0)
		//x = 0.0
	}
    if(teste < 10){
        teste++
    }
    if(teste == 10){
        AtualizaGrid(110,146)
        // Modificações para testar que o grid se altera
        teste = 0
    }
	return nil
}

//Conjunto de Genes para sobreviver a regiões que cada celula pode ter
type Genes struct{
    agua bool
    calor bool
    frio bool
    altitude bool
    comida bool
}

//Estrutura de cada celula do grid
type Celula struct{
    viva bool
    especie int
    mutacoes Genes
    posX float64
    posY float64
    time_vida time.Time
}


//Inicia cada celula do grid
func IniciaGrid(N int,M int, grid [][]Celula){
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    mutacao := -1
    eixoX, eixoY := 1.0, 1.0

    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {

            grid[i][j].posX=eixoX
            grid[i][j].posY=eixoY
            eixoX+=6.0
            //Seta a vida da celula com 5% de chance dela nascer
            if(r1.Intn(20) == 1){
                grid[i][j].viva = true
                grid[i][j].time_vida = time.Now()
                //Caso a celula esteja viva sorteia com 8% de chance cada gene de mutacao
                for l := 0; l < 6; l++ {
                    mutacao = r1.Intn(25)
                    switch l {
                        case 0:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.agua = true
                            }
                            break
                        case 1:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.calor = true
                            }
                            break
                        case 2:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.frio = true
                            }
                            break
                        case 3:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.altitude = true
                            }
                            break
                        case 4:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.comida = true
                            }
                    }
                }
            }else {
                grid[i][j].viva = false
            }

        }//end j

        eixoX=1.0
        eixoY+=6.0
    }//end i


}

//Percorre a matriz do grid e mostra na tela
func MostraGrid(N int, M int, grid [][]Celula){
    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {
		if(grid[i][j].viva){
			fmt.Printf("#")
		} else {
			fmt.Printf(" ")
		}
        }
        fmt.Printf("\n")
    }

}

type Old_State struct{
	i,j int
	vida bool
}

func Jogo(rule Rule,N int, M int, i int, j int, hold chan int, result chan Old_State) {
    	//Verifica todos os vizinhos da celula
    	domain := bitfy(i,j,N,M)

	//Abaixo estão sendo usada as regras do jogo da vida de conway, a ideia e tornar essas 		regras flexiveis a outros valores
    	var ret Old_State
    	ret.i= i
    	ret.j= j
    	ret.vida = apply_rule(rule,domain)

    	hold <- 1
    	result <- ret
}



func AtualizaGrid(N int,M int){
    	//cont := 0
    	//continua := true

	conway := rule_conway()
	//conway := rule_random()

    	//for continua != false {

        	//CallClear()
        	//MostraGrid(N,M,grid)

        	//time.Sleep(100 * time.Millisecond)
    hold := make(chan int,0)
	result := make(chan Old_State)

    for i := 0; i < N; i++ {
    	for j := 0; j < M; j++ {
        	go Jogo(conway,N,M,i,j,hold,result)
    	}
    }

	for i := 0; i < N*M; i++{
		<-hold
	}

	var new Old_State
	for i := 0; i < N*M; i++{
		new = <-result
		grid[new.i][new.j].viva = new.vida
        if(new.vida == true){
            grid[new.i][new.j].time_vida = time.Now()
        }
	}


}

func Teste (N int,M int) {
    eixoX, eixoY := 1.0, 1.0

	for i:=0;i<N;i++ {
		for j:=0;j<M;j++{
			grid[i][j].viva=false
            grid[i][j].posX=eixoX
            grid[i][j].posY=eixoY
            eixoX+=6.0
		}
        eixoX=1.0
        eixoY+=6.0
	}
    //Glider
	grid[0][1].viva = true
    grid[0][1].time_vida = time.Now()
	grid[1][2].viva = true
    grid[1][2].time_vida = time.Now()
	grid[2][0].viva = true
    grid[2][0].time_vida = time.Now()
	grid[2][1].viva = true
    grid[2][1].time_vida = time.Now()
	grid[2][2].viva = true
    grid[2][2].time_vida = time.Now()

    //10 cell row
    grid[20][12].viva = true
	grid[20][13].viva = true
	grid[20][14].viva = true
	grid[20][15].viva = true
	grid[20][16].viva = true
    grid[20][17].viva = true
	grid[20][18].viva = true
	grid[20][19].viva = true
	grid[20][20].viva = true
	grid[20][21].viva = true

}



func main(){

	for i := range grid {
    		grid[i] = make([]Celula,200,200)
	}

    //Teste(110,146)
    IniciaGrid(110,146, grid)

    if err := ebiten.Run(update, 439, 331, 2, "Novo jogo da vida"); err != nil {
		log.Fatal(err)
	}
    //IniciaGrid(100,100,grid)
    //MostraGrid(100,100,grid)

}
