package main

func aloca_simulacao(I int,J int) Simulacao {
	var simulacao Simulacao
	simulacao.grid.array = make([][]Celula,I,I)
	for i := range simulacao.grid.array {
    		simulacao.grid.array[i] = make([]Celula,J,J)
	}
	simulacao.backup.array = make([][]Celula,I,I)
	for i := range simulacao.backup.array {
    		simulacao.backup.array[i] = make([]Celula,J,J)
	}
	simulacao.output.array = make([][]Output,I,I)
	for i := range simulacao.output.array {
    		simulacao.output.array[i] = make([]Output,J,J)
	}
	simulacao.grid.I,simulacao.grid.J = I,J
	simulacao.backup.I,simulacao.backup.J = I,J
	simulacao.output.I,simulacao.output.J = I,J

	return simulacao
}


func simulate(regra func(int,int,vetCelula) Celula,simulacao Simulacao){		
	for true {
		for testes:=0 ; testes<0 ; testes ++{
		} 
    	    	AtualizaGrid(regra,simulacao.grid,simulacao.backup)
		AtualizaOutput(simulacao)

		aux := simulacao.grid
		simulacao.grid = simulacao.backup
		simulacao.backup = aux

		//SwapFrame(simulacao)
	}
}

func SwapFrame(simulacao Simulacao) {
	// Lembrando que usamos slices, é só trocar os ponteiros
	aux := simulacao.grid
	simulacao.grid = simulacao.backup
	simulacao.backup = aux
}