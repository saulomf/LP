package main


type Old_State struct{
	i,j int
	vida bool
}

type Rule [512]bool


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
    	posX float64
    	posY float64
}
