package main

import (
    	"os"
	"os/exec"
    	"runtime"

     _ "image/png"
     "log"
     "image/color"

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
	blu, _, err = ebitenutil.NewImageFromFile("blu.png", ebiten.FilterDefault)
	red, _, err = ebitenutil.NewImageFromFile("red.png", ebiten.FilterDefault)
	blank, _, err = ebitenutil.NewImageFromFile("blank.png", ebiten.FilterDefault)
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

func update(screen *ebiten.Image) error{
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.Fill(color.RGBA{0, 0xff, 0, 0xff})
	
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(1, 1)
	//screen.DrawImage(img, op))
	
	var grid [][]Celula
	grid = tela

	//x := 0.0
	for j := 0; j < 38; j++ {
		op.GeoM.Translate(grid[0][j].posX, grid[0][j].posY)
 	     	if(grid[0][j].viva){
		     	if (grid[0][j].especie == 1){
				screen.DrawImage(blu, op)
			} else {    			
				screen.DrawImage(red, op)
			}
     		} else {
        		screen.DrawImage(blank, op)
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
          	if(grid[i][j].viva){
			if (grid[i][j].especie == 1){
				screen.DrawImage(blu, op)
			} else {    			
				screen.DrawImage(red, op)
			}
		} else {
           	screen.DrawImage(blank, op)
          	}
          	op.GeoM.Reset()
			//x-=11.0
		}
		//op.GeoM.Translate(x, 0)
		//x = 0.0
	}
	return nil
}


//Inicia cada celula do grid
func IniciaGrid(I int,J int, grid [][]Celula){
    mutacao := -1
	posX := 1.0
	posY := 1.0
    	for i := 0; i < I; i++ {
		posX = 1.0
        	for j := 0; j < J; j++ {
			grid[i][j].posX=posX
			grid[i][j].posY=posY
            	//Seta a vida da celula com 10% de chance dela nascer
            	if(r1.Intn(2) == 1){
                	grid[i][j].viva = true
				grid[i][j].especie = r1.Intn(2)+1 // especie
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
				grid[i][j].especie = 0
            	}
			posX+=11.0
        	}//end j
		posY+=11.0
    }//end i


}

func Jogo(rule Rule,I int, J int, i int, j int, hold chan int, result chan Old_State,grid [][]Celula) {
    	//Verifica todos os vizinhos da celula
    	domain_blu := bitfy(i,j,I,J,grid,1)
	domain_red := bitfy(i,j,I,J,grid,2)

	//Abaixo estão sendo usada as regras do jogo da vida de conway, a ideia e tornar essas 		regras flexiveis a outros valores
    	var ret Old_State
    	ret.i= i
    	ret.j= j
	if(apply_rule(rule,domain_blu) == apply_rule(rule,domain_red)){
		ret.vida = false
		ret.especie = 0
	} else {
		ret.vida = true
		if(apply_rule(rule,domain_blu)){
			ret.especie = 1
		} else { // red
			ret.especie = 2
		}	
	}

    	hold <- 1
    	result <- ret
}



func AtualizaGrid(rule Rule,grid [][]Celula,I int,J int){
	hold := make(chan int,0)
	result := make(chan Old_State)

    	for i := 0; i < I; i++ {
    		for j := 0; j < J; j++ {
     	   		go Jogo(rule,I,J,i,j,hold,result,grid)
    		}
    	}

	for i := 0; i < I*J; i++{
		<-hold
	}

	var new Old_State
	for i := 0; i < I*J; i++{
		new = <-result
		grid[new.i][new.j].viva = new.vida
		grid[new.i][new.j].especie=new.especie
	}
}

func glider(grid [][]Celula,I int,J int) {
    eixoX, eixoY := 1.0, 1.0

	for i:=0;i<I;i++ {
		for j:=0;j<J;j++{
			grid[i][j].viva=false
			grid[i][j].especie=0
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

	grid[0][1].especie = 1
	grid[1][2].especie = 1
	grid[2][0].especie = 1
	grid[2][1].especie = 1
	grid[2][2].especie = 1

}

func simulate(rule Rule,I int,J int,grid [][]Celula){		
	teste := 0
	for true {
    		if(teste >= 30){
    	    		AtualizaGrid(rule,grid,I,J)
    	    		// Modificações para testar que o grid se altera
    	    		teste = 0
    		}
		teste++
	}
}

func chose(grid1 [][]Celula,grid2 [][]Celula){
	for i:=0; true; i = (i+1)%20000{
		if(i<10000){
			tela = grid1
		} else {
			tela = grid2
		}	
	}
}

func aloca_simulacao(I int,J int)[][]Celula {
	simulacao := make([][]Celula,I,I)
	for i := range simulacao {
    		simulacao[i] = make([]Celula,J,J)
	}

	return simulacao
}

func main(){
	simulacao1 := aloca_simulacao(28,38)
	simulacao2 := aloca_simulacao(28,38)
	simulacao3 := aloca_simulacao(28,38)

   	glider(simulacao1,28,38)
	IniciaGrid(28,38,simulacao2)
	IniciaGrid(28,38,simulacao3)

	go simulate(rule_conway(),28,38,simulacao1)
	go simulate(rule_conway(),28,38,simulacao2)
	go simulate(rule_random(),28,38,simulacao3)
	tela = simulacao2
	//go chose(simulacao1,simulacao2)

    	if err := ebiten.Run(update, 420, 310, 2, "Novo jogo da vida"); err != nil {
		log.Fatal(err)
	}

}
