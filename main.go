package main

import (
    	"log"

     	"github.com/hajimehoshi/ebiten"
)



//Inicia cada celula do grid
func IniciaGrid(grid vetCelula){
	mutacao := -1
   	for i := 0; i < grid.I; i++ {
        	for j := 0; j < grid.J; j++ {
            	//Seta a vida da celula com 10% de chance dela nascer
            	if(r1.Intn(2) == 1){
                	grid.array[i][j].viva = true
				grid.array[i][j].especie = r1.Intn(2)+1 // especie
                	//Caso a celula esteja viva sorteia com 5% de chance cada gene de mutacao
                	for l := 0; l < 6; l++ {
                    	mutacao = r1.Intn(30)
                    	switch l {
                        	case 0:
                            		if((mutacao >= 0) && (mutacao <1)){
                              	 	 	grid.array[i][j].mutacoes.agua = true
                            		}
                            		break
                        	case 1:
                            		if((mutacao >= 0) && (mutacao <1)){
                                		grid.array[i][j].mutacoes.calor = true
                            		}
                            		break
                        	case 2:
                            		if((mutacao >= 0) && (mutacao <1)){
                                		grid.array[i][j].mutacoes.frio = true
                            		}
                            		break
                        	case 3:
                            		if((mutacao >= 0) && (mutacao <1)){
                                		grid.array[i][j].mutacoes.altitude = true
                            		}
                            		break
                        	case 4:
                            		if((mutacao >= 0) && (mutacao <1)){
                                		grid.array[i][j].mutacoes.comida = true
                            		}
                    		}
                	}
            	} else {
                	grid.array[i][j].viva = false
			grid.array[i][j].especie = 0
            	}
        	}//end j
   	}//end i
}

func main(){
	aloca_output()

	simulacao2 := aloca_simulacao(50,70)
	simulacao3 := aloca_simulacao(50,70)

	IniciaGrid(simulacao2.grid)
	IniciaGrid(simulacao3.grid)

	go simulate(rule_conway(),simulacao2)
	go simulate(rule_random(),simulacao3)
	
	selecoes := aloca_muxsimulacao(2)
	selecoes.array[0] = simulacao2
	selecoes.array[1] = simulacao3
	go chose(selecoes)

    	if err := ebiten.Run(update, outputx, outputy, 2, "Novo jogo da vida"); err != nil {
		log.Fatal(err)
	}

}
