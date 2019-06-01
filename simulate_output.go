package main

func AtualizaOutput(simulacao Simulacao) {
	for i := 0 ; i < simulacao.output.I; i++ {
		for j := 0 ; j < simulacao.output.J; j++ {
    			if(simulacao.grid.array[i][j].viva){
		     		if (simulacao.grid.array[i][j].especie == 1){
					simulacao.output.array[i][j].imagem = blu
				} else {    			
					simulacao.output.array[i][j].imagem = red
				}
			} else {
				simulacao.output.array[i][j].imagem = nil	
			}
		}
	}
}
