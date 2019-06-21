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
            	if(r.Intn(2) == 1){
                	grid.array[i][j].viva = true
				grid.array[i][j].especie = r.Intn(2)+1 // especie
                	//Caso a celula esteja viva sorteia com 5% de chance cada gene de mutacao
                	for l := 0; l < 6; l++ {
                    	mutacao = r.Intn(30)
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

	ebiten.SetMaxTPS(30)
	// Removida a funcionalidade de executar mais de uma simulação
	//aloca_output()

	simulacao := aloca_simulacao(cellsy,cellsx)

	IniciaGrid(simulacao.backup)

	tela.output = &simulacao.output
	go simulate(regra_celula,simulacao)

	// Eliminei a funcionalidade de executar mais de uma simulação	
	//selecoes := aloca_muxsimulacao(2)
	//selecoes.array[0] = simulacao2
	//selecoes.array[1] = simulacao3
	//go chose(selecoes)


    	if err := ebiten.Run(update, outputx, outputy, 2, "Novo jogo da vida"); err != nil {
		log.Fatal(err)
	}
}
