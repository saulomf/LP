package main

import "fmt"

package water_flow

import (
	
	"image/color"
	_ "image/png"
	"log"
	"math"
	//"fmt"

	"github.com/hajimehoshi/ebiten"
)

type Submatrizes struct{
	frameWrite [][]SubMundo
	frameRead  [][]SubMundo
}
func (matriz *Submatrizes) swap_frame() {
	aux := matriz.frameRead
	matriz.frameRead = matriz.frameWrite
	matriz.frameWrite = aux 
}

type SubMundo struct{
	posi,posj int	// Posicao do inicio no grid real
	mundo Mundo 	// grid
}

const (
	max_size = 100
)

func subdivide(mundo *Mundo) Submatrizes {

	frameRead := recursive_subdivide(0,0,mundo)		

	frameWrite := make([] []SubMundo,len(frameRead))
	for i,linha:= range frameWrite {
		linha = make([] SubMundo,len(frameRead[i]))
	}

	submatrizes := &Submatrizes{frameWrite: frameWrite,frameRead: frameRead}

	return submatrizes
}

func recursive_subdivide(posi,posj int,mundo *Mundo) [][]SubMundo {
	I,J := len(mundo),len(mundo[0])
	
	if num := I*J; num <= max_size {
		// Ponto de parada: tamanho menor ou igual ao desejado
		submundos := make([] []Submundo,1)	
		submundos[0] := make([]Submundo,1)

		submundos[0][0].mundo = mundo
		submundos[0][0].posi = posi
		submundos[0][0].posj = posj

		return submundos
	}
	
	
	if(I >= J) {
		// Corte horizontal - Separar linhas - Dividir o vetor linha
		meio := I/2 + 1
			
		mundoNorte,mundoSul := corte_horizontal(meio,mundo)
			
		submundosNorte := recursive_subdivide(posi,posj,mundoNorte) 
		submundosSul :=   recursive_subdivide(posi + meio,posj,mundoSul)

		return append(submundosNorte,submundosSul...)

	} else {
		// Corte vertical - Separar colunas - Dividir vetores colunas e criar vetor linha para conecta-los
		meio := J/2 + 1
		
		mundoLeste,mundoOeste := corte_vertical(meio,mundo)
			
		submundosLeste := recursive_subdivide(posi,posj, mundoLeste) 
		submundosOeste := recursive_subdivide(posi,posj + meio, mundoOeste)

		submundosReturn := submundosLeste

		for i:= range submundosReturn {
			submundosReturn[i] = append(submundosLeste[i],submundosOeste[i]...)
		}

		return submundosReturn
	}
}

func corte_horizontal(meio int,mundo *Mundo) (*Mundo,*Mundo) {
	var mundoNorte,mundoSul *Mundo
	
	mundoNorte.Planta,	mundoSul.Planta 	= mundo.Planta[:meio],		mundo.Planta[meio:]
	mundoNorte.Herbivoro,	mundoSul.Herbivoro 	= mundo.Herbivoro[:meio],	mundo.Herbivoro[meio:]
	mundoNorte.Carnivoro,	mundoSul.Carnivoro 	= mundo.Carnivoro[:meio],	mundo.Carnivoro[meio:]
	mundoNorte.Agua,	mundoSul.Agua 		= mundo.Agua[:meio],		mundo.Agua[meio:]
	mundoNorte.Ar,		mundoSul.Ar 		= mundo.Ar[:meio],		mundo.Ar[meio:]
	mundoNorte.Relevo,	mundoSul.Relevo 	= mundo.Relevo[:meio],		mundo.Relevo[meio:]

	return mundoNorte,mundoSul
}

func corte_vertical(meio int,mundo *Mundo) (*Mundo,*Mundo) {
	var mundoLeste,mundoOeste *Mundo
	// Formar linhas do Oeste
	mundoOeste.Planta 	:= make([] []Planta)
	mundoOeste.Herbivoro 	:= make([] []Herbivoro)
	mundoOeste.Carnivoro 	:= make([] []Carnivoro)
	mundoOeste.Agua 	:= make([] []Agua)
	mundoOeste.Ar 		:= make([] []Ar)
	mundoOeste.Relevo 	:= make([] []Relevo)
	
	for i:= range Ar {
		mundoLeste.Planta,	mundoOeste.Planta 	= mundo.Planta[i][:meio],	mundo.Planta[i][meio:]
		mundoLeste.Herbivoro,	mundoOeste.Herbivoro 	= mundo.Herbivoro[i][:meio],	mundo.Herbivoro[i][meio:]
		mundoLeste.Carnivoro,	mundoOeste.Carnivoro 	= mundo.Carnivoro[i][:meio],	mundo.Carnivoro[i][meio:]
		mundoLeste.Agua,	mundoOeste.Agua 	= mundo.Agua[i][:meio],		mundo.Agua[i][meio:]
		mundoLeste.Ar,		mundoOeste.Ar 		= mundo.Ar[i][:meio],		mundo.Ar[i][meio:]
		mundoLeste.Relevo,	mundoOeste.Relevo 	= mundo.Relevo[i][:meio],	mundo.Relevo[i][meio:]
	}
	
	return mundoLeste,mundoOeste
}