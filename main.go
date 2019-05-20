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
    viva int
    mutacoes Genes
}

//Inicia cada celula do grid
func IniciaGrid(grid *[30][30]Celula){
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    mutacao := -1
    for i := 0; i < 30; i++ {
        for j := 0; j < 30; j++ {
            //Seta a vida da celula com 10% de chance dela nascer
            if(r1.Intn(2) == 1){

                grid[i][j].viva = 1
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
                grid[i][j].viva = 0
            }

        }//end j
    }//end i


}

//Percorre a matriz do grid e mostra na tela
func MostraGrid(grid [30][30]Celula){
    for i := 0; i < 30; i++ {
        for j := 0; j < 30; j++ {
            fmt.Printf("%b ", grid[i][j].viva)
        }
        fmt.Printf("\n")
    }

}


func Jogo(grid *[30][30]Celula, i int, j int)  {
    vivos := 0 //quantidade de vizinhos vivos
    //Verifica todos os vizinhos da celula
    for posI := i-1; posI < i+1; posI++ {
        for posJ := j-1; posJ < j+1; posJ++ {
            //fmt.Printf("\n%d %d", posI, posJ)

            if(grid[posI][posJ].viva == 1){
                vivos++
            }
        }
    }

    //Abaixo estão sendo usada as regras do jogo da vida de conway, a ideia e tornar essas regras flexiveis a outros valores

    vivos-- //retira a vida da propria celula da contagem
    if(grid[i][j].viva == 1){
        if(vivos < 2){//morre de solidao
            grid[i][j].viva = 0
        }
        if(vivos > 3){//morre de superpopulacao
            grid[i][j].viva = 0
        }
        //caso tenha 2 ou 3 vizinhos continua viva
    } else {
        if(vivos == 3){
            grid[i][j].viva = 1//Celula antes morta nasce
        }
    }
}

func AtualizaGrid(grid [30][30]Celula){
    cont := 0
    continua := true

    for continua != false {
        for i := 1; i < 29; i++ {
            for j := 1; j < 29; j++ {
                Jogo(&grid, i, j)
            }
        }
        time.Sleep(500 * time.Millisecond)
        CallClear()
        MostraGrid(grid)
        //fmt.Printf("\n")

        if(cont == 100){//Garantir que não rode infinitamente
            continua = false
        }
    }
}

func main(){
    var grid [30][30]Celula

    IniciaGrid(&grid)
    MostraGrid(grid)
    AtualizaGrid(grid)


}
