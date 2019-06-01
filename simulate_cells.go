package main

func Jogo(rule Rule, i int, j int, hold chan int, result chan Old_State,grid vetCelula) {
    	//Verifica todos os vizinhos da celula
    	domain_blu := bitfy(i,j,grid,1)
	domain_red := bitfy(i,j,grid,2)

	//Abaixo estão sendo usada as regras do jogo da vida de conway, a ideia e tornar essas 	regras flexiveis a outros valores
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



func AtualizaGrid(rule Rule,grid vetCelula){
	hold := make(chan int,0)
	result := make(chan Old_State)

    	for i := 0; i < grid.I; i++ {
    		for j := 0; j < grid.J; j++ {
     	   		go Jogo(rule,i,j,hold,result,grid)
    		}
    	}

	for i := 0; i < grid.I*grid.J; i++{
		<-hold
	}

	var new Old_State
	for i := 0; i < grid.I*grid.J; i++{
		new = <-result
		grid.array[new.i][new.j].viva = new.vida
		grid.array[new.i][new.j].especie=new.especie
	}
}
