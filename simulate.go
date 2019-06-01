package main

func aloca_simulacao(I int,J int) Simulacao {
	var simulacao Simulacao
	simulacao.grid.array = make([][]Celula,I,I)
	for i := range simulacao.grid.array {
    		simulacao.grid.array[i] = make([]Celula,J,J)
	}
	simulacao.output.array = make([][]Output,I,I)
	for i := range simulacao.output.array {
    		simulacao.output.array[i] = make([]Output,J,J)
	}
	simulacao.grid.I,simulacao.grid.J = I,J
	simulacao.output.I,simulacao.output.J = I,J

	return simulacao
}


func simulate(rule Rule,simulacao Simulacao){		
	teste := 0
	for true {
    		if(teste >= 30){
    	    		AtualizaGrid(rule,simulacao.grid)
			AtualizaOutput(simulacao)
    	    		// Modificações para testar que o grid se altera
    	    		teste = 0
    		}
		teste++
	}
}