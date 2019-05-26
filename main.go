package main

import (
    	"fmt"
    	"time"
    	"math/rand"
    	"os"
    	"os/exec"
    	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

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
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
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

func Jogo(rule Rule,N int, M int,grid [][]Celula, i int, j int, hold chan int, result chan Old_State) {
    	//Verifica todos os vizinhos da celula
    	domain := bitfy(i,j,N,M,grid)

	//Abaixo estão sendo usada as regras do jogo da vida de conway, a ideia e tornar essas 		regras flexiveis a outros valores
    	var ret Old_State
    	ret.i= i
    	ret.j= j
    	ret.vida = apply_rule(rule,domain)
    	
    	hold <- 1
    	result <- ret
}

func AtualizaGrid(N int,M int,grid [][]Celula){
    	cont := 0
    	continua := true
	//conway := rule_conway()
	conway := rule_random()

    	for continua != false {
		
        	CallClear()
        	MostraGrid(N,M,grid)
        	time.Sleep(100 * time.Millisecond)

		hold := make(chan int,0)
		result := make(chan Old_State)
		
        	for i := 0; i < N; i++ {
        		for j := 0; j < M; j++ {
        	        	go Jogo(conway,N,M,grid,i,j,hold,result)
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

func Teste (N int,M int,Celulas [][]Celula) {
	for i:=0;i<N;i++ {
		for j:=0;j<M;j++{
			Celulas[i][j].viva=false
		}
	}
	Celulas[0][1].viva = true
	Celulas[1][2].viva = true
	Celulas[2][0].viva = true
	Celulas[2][1].viva = true
	Celulas[2][2].viva = true
}

func main(){
    	grid := make([][]Celula,100,100) 
	for i := range grid {
    		grid[i] = make([]Celula,100,100)
	}

    	Teste(100,100,grid)
    	//IniciaGrid(100,100,grid)
    	MostraGrid(100,100,grid)
    	AtualizaGrid(100,100,grid)
}
