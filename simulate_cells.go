package main

func Jogo(regra func(int,int,vetCelula) Celula, i int, j int, result chan Old_State,grid vetCelula) {

	//Abaixo estão sendo usada as regras do jogo da vida de conway, a ideia e tornar essas 	regras flexiveis a outros valores
    	var ret Old_State
    	ret.i= i
    	ret.j= j

	ret.celula =  regra(i,j,grid)

    	result <- ret
}

func regra_celula(i int,j int,grid vetCelula) Celula {
	//Verifica todos os vizinhos da celula
    	domain_blu := bitfy(i,j,grid,1)
	domain_red := bitfy(i,j,grid,2)

	var ret Celula

	if(apply_rule(rule_conway(),domain_blu) == apply_rule(rule_conway(),domain_red)){
		ret.viva = false
		ret.especie = 0
	} else {
		ret.viva = true
		if(apply_rule(rule_conway(),domain_blu)){
			ret.especie = 1
		} else { // red
			ret.especie = 2
		}	
	}

	return ret
}

func AtualizaGrid(regra func(int,int,vetCelula) Celula,grid vetCelula,backup vetCelula){
	result := make(chan Old_State,grid.I*grid.J)

    	for i := 0; i < backup.I; i++ {
    		for j := 0; j < backup.J; j++ {
     	   		go Jogo(regra,i,j,result,backup)
    		}
    	}

	var new Old_State
	for i := 0; i < grid.I*grid.J; i++{
		new = <-result
		grid.array[new.i][new.j].viva = new.celula.viva
		grid.array[new.i][new.j].especie=new.celula.especie
	}
}
