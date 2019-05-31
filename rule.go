package main

import (
)

/*
	Pretendo criar uma função que associe cadeias de bits,
	correspondentes às células no alcance de uma específica
	célula, a um valor booleano, que define o próximo estágio dela.


	Uma função
	f: {x e N | 0 <= x < 512 } -> {0,1}

	-> Rule type array: vetor de 512 bits que associa uma conformação de células
	a um booleano de resultado
*/


func bitfy(pos_i int,pos_j int,N int,M int) int {
	bit := 1 // Bit atual
	ret := 0 // RETORNO
	// de 0.00 00.0 001 até 1.00 00.0 000
	// (i,j) = célula sendo analisada
	for i:= pos_i-1; i <= pos_i+1; i++ {
		for j:= pos_j-1; j <= pos_j+1; j++ {
			if(grid[(i+N)%N][(j+M)%M].viva){
				ret |= bit // obter bit caso tenha célula viva
			}
			bit <<= 1
		}
	}
	return ret
}

func apply_rule(rule Rule,domain int) bool {
	return rule[domain]
}
