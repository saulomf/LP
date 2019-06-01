package main

import(
     	"github.com/hajimehoshi/ebiten"
)

type Old_State struct{
	i,j int
	vida bool
	especie int
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
}

// Coleção de células
type vetCelula struct{
	I int
	J int
	array [][]Celula
}

//Estrutura para cada unidade de output
type Output struct {
	imagem *ebiten.Image
}

type vetOutput struct {
	I int
	J int
	array [][]Output
}

// Struct para representar simulações
type Simulacao struct{
	grid vetCelula
	output vetOutput
}

// Struct para selecionar simulacao
type muxSimulacao struct{
	array []Simulacao
	i int	// Selecao
	N int 	// Tamanho atual
}	
