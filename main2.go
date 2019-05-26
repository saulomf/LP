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
        "image/color"

        "github.com/hajimehoshi/ebiten"
        "github.com/hajimehoshi/ebiten/ebitenutil"
)







func rule_random () Rule {
    	s1 := rand.NewSource(time.Now().UnixNano())
    	r1 := rand.New(s1)
	var rule Rule
	// 000 000 000 até 111 111 111
	for domain := 0; domain < 512; domain++ {
		if(r1.Intn(2)==1){
			rule[domain] = true
		} else {
			 rule[domain] = false
		}
	}

	return rule
}
func bitfy(pos_i int,pos_j int,N int,M int) int {
	bit := 1 // Bit atual
	ret := 0 // RETORNO
	// de 0.00 00.0 001 até 1.00 00.0 000
	// (i,j) = célula sendo analisada
	for i:= pos_i-1; i <= pos_i+1; i++ {
		for j:= pos_j-1; j <= pos_j+1; j++ {
			if(grid[(i+N)%N][(j+M)%M].viva){
				ret |= bit // obter bit caso tenha célula viva
			}
			bit <<= 1
		}
	}
	return ret
}

func apply_rule(rule Rule,domain int) bool {
	return rule[domain]
}
func conway_game (vivos int, vida bool) bool {
	if(vivos == 3) {
            	vida = true //Celula nasce se morta ou permanece se viva
        } else if(vivos < 2) {
		vida = false // Morre de solidao se viva
	} else if(vivos > 3) {
		vida = false // Morre de superpopulacao se viva
   	}
	return vida
}

func rule_conway () Rule {
	var rule Rule
	// 000 000 000 até 111 111 111
	for domain := 0; domain < 512; domain++ {
		vivos := 0
		for i := domain; i > 0 ; i>>=1 {
			if(i&0x1 == 1){
				vivos++
			}
		}

		// xxx x#x xxx -> # = celula em si
		if(domain&0x010 == 0) {
			rule[domain] = conway_game(vivos,false)
		} else {
			rule[domain] = conway_game(vivos-1,true)
		}
	}

	return rule
}















var clear map[string]func() //create a map for storing clear funcs
var img *ebiten.Image //Imagem 1 quadrado preto
var img2 *ebiten.Image //Imagem 2 quadrado branco
var grid = make([][]Celula,100,100) // grid global
var teste int = 0 // variavel para teste de controle do grid

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
	img, _, err = ebitenutil.NewImageFromFile("square.png", ebiten.FilterDefault)
	img2, _, err = ebitenutil.NewImageFromFile("blank.png", ebiten.FilterDefault)
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
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})

	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(1, 1)
	//screen.DrawImage(img, op))

	//x := 0.0
	for j := 0; j < 38; j++ {
		op.GeoM.Translate(grid[0][j].posX, grid[0][j].posY)
        if(grid[0][j].viva == true){
		          screen.DrawImage(img2, op)
        } else {
            screen.DrawImage(img, op)
        }
        op.GeoM.Reset()
		//x-=11.0
	}
	//op.GeoM.Translate(x, 0)
	//x = 0.0

	for i := 0; i < 28; i++ {
		//op.GeoM.Translate(0, 11)
		//screen.DrawImage(img, op)
		for j := 0; j < 38; j++ {
    		op.GeoM.Translate(grid[i][j].posX, grid[i][j].posY)
            if(grid[i][j].viva == true){
    		          screen.DrawImage(img2, op)
            } else {
                screen.DrawImage(img, op)
            }
            op.GeoM.Reset()
			//x-=11.0
		}
		//op.GeoM.Translate(x, 0)
		//x = 0.0
	}
    if(teste < 300){
        teste++
    }
    if(teste == 300){
        teste = 0
    }
    if(teste == 180){
        //AtualizaGrid(28,38)
        // Modificações para testar que o grid se altera
        grid[0][10].viva = true
        grid[1][20].viva = true
        grid[2][0].viva = true
        grid[2][10].viva = true
        grid[2][20].viva = true
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
}


//Inicia cada celula do grid
func IniciaGrid(N int,M int, grid [][]Celula){
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    mutacao := -1
    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {
            //Seta a vida da celula com 10% de chance dela nascer
            if(r1.Intn(2) == 1){

                grid[i][j].viva = true
                //Caso a celula esteja viva sorteia com 5% de chance cada gene de mutacao
                for l := 0; l < 6; l++ {
                    mutacao = r1.Intn(30)
                    switch l {
                        case 0:
                            if((mutacao >= 0) && (mutacao <1)){
                                grid[i][j].mutacoes.agua = true
                            }
                            break
                        case 1:
                            if((mutacao >= 0) && (mutacao <1)){
                                grid[i][j].mutacoes.calor = true
                            }
                            break
                        case 2:
                            if((mutacao >= 0) && (mutacao <1)){
                                grid[i][j].mutacoes.frio = true
                            }
                            break
                        case 3:
                            if((mutacao >= 0) && (mutacao <1)){
                                grid[i][j].mutacoes.altitude = true
                            }
                            break
                        case 4:
                            if((mutacao >= 0) && (mutacao <1)){
                                grid[i][j].mutacoes.comida = true
                            }
                    }
                }
            } else {
                grid[i][j].viva = false
            }

        }//end j
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

type Rule [512]bool

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
    	cont := 0
    	continua := true

	//conway := rule_conway()
	conway := rule_random()

    	for continua != false {

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
		}

        	//fmt.Printf("\n")

        	if(cont == 100){//Garantir que não rode infinitamente
        		continua = false
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
            eixoX+=11.0
		}
        eixoX=1.0
        eixoY+=11.0
	}
	grid[0][1].viva = true
	grid[1][2].viva = true
	grid[2][0].viva = true
	grid[2][1].viva = true
	grid[2][2].viva = true

}



func main(){

	for i := range grid {
    		grid[i] = make([]Celula,100,100)
	}

    Teste(28,38)

    if err := ebiten.Run(update, 420, 310, 2, "Novo jogo da vida"); err != nil {
		log.Fatal(err)
	}
    //IniciaGrid(100,100,grid)
    //MostraGrid(100,100,grid)

}
