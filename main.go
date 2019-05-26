package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var s = rand.NewSource(time.Now().UnixNano())
var r = rand.New(s)

type Environment struct {
	agua int
	calor int
	gases int
	luz int
}

type Genoma struct{
	agua int
	calor int
	gases int
	luz int
	v_agua int
	v_calor int
	v_gases int
	v_luz int
	variabilidade [6] int
}

type Especie struct {
	nome string
	agua int
	calor int
	gases int
	luz int
}

//TODO Montar catalogo de especies de acordo com seus atributos
//func catalogaEspecie (nomeEspecies []string, catalogoEspecies map[string]int, especie Especie)map[string]int{
//	var existeEspecie bool
//	tmp := nomeiaEspecie(nomeEspecies)
//	if len(catalogoEspecies) == 0{
//		especie.nome = tmp[len(tmp) - 1]
//		catalogoEspecies[tmp] = especie
//	} else {
//		for _, each := range catalogoEspecies{
//			if especie.agua == each.agua
//		}
//	}
//}

func nomeiaEspecie(nomeEspecies []string)[]string{
	var alfabeto string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var ultimoNome string
	var flag int = 0
	if len(nomeEspecies) == 0{
		nomeEspecies = append(nomeEspecies, "A")
	} else {
		for _, ultimo := range nomeEspecies{
			ultimoNome = ultimo
		}
		tmp := []byte(ultimoNome)
		if ultimoNome[len(ultimoNome) - 1] == 90{
			for _, letter := range tmp{
				if letter != 90 {
					flag = 1
				}
			}
			if flag == 0{
				tmp = append(tmp, 65)
				tmp = []byte(strings.ReplaceAll(string(tmp), "Z", "A"))
				nomeEspecies = append(nomeEspecies, string(tmp))
			} else {
				tmp[len(tmp) - 2] = tmp[len(tmp) - 2] + 1
				tmp[len(tmp) - 1] = 65
				nomeEspecies = append(nomeEspecies, string(tmp))
			}
		}else {
			index := strings.Index(alfabeto, string(ultimoNome[len(ultimoNome) - 1])) + 1
			tmp[len(tmp) - 1] = alfabeto[index]
			nomeEspecies = append(nomeEspecies, string(tmp))
		}
	}
	return nomeEspecies
}

func mostraIndividuo(especie Genoma){
	fmt.Println(especie)
}

func geraVariabilidade(genes Genoma) Genoma{
	genes.variabilidade[0] = r.Intn(10) + 1
	genes.variabilidade[1] = r.Intn(10) + 1
	genes.variabilidade[2] = r.Intn(10) + 1
	genes.variabilidade[3] = r.Intn(10) + 1
	genes.variabilidade[4] = r.Intn(10) + 1
	genes.variabilidade[5] = r.Intn(10) + 1
	return genes
}

func geraIndividuo(genes Genoma) Genoma{
	genes.v_agua = r.Intn(10) + 1
	genes.v_calor = r.Intn(10) + 1
	genes.v_gases = r.Intn(10) + 1
	genes.v_luz = r.Intn(10) + 1
	return genes
}

func geraEspecie() Genoma{
	var genes Genoma
	genes.agua = r.Intn(10) + 1
	genes.calor = r.Intn(10) + 1
	genes.gases = r.Intn(10) + 1
	genes.luz = r.Intn(10) + 1
	return genes
}

func mostraAgua(grid [10][10]Environment){
	fmt.Println("RECURSOS HIDRICOS\n")
	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++{
			fmt.Printf("%d\t", grid[i][j].agua)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func mostraCalor(grid [10][10]Environment){
	fmt.Println("MAPA DE CALOR\n")
	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++{
			fmt.Printf("%d\t", grid[i][j].calor)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func mostraGases(grid [10][10]Environment){
	fmt.Println("DISPERSAO DE GASES\n")
	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++{
			fmt.Printf("%d\t", grid[i][j].gases)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func mostraLuz(grid [10][10]Environment){
	fmt.Println("INCIDENCIA DE LUZ\n")
	for i := 0; i < 10; i++{
		for j := 0; j < 10; j++{
			fmt.Printf("%d\t", grid[i][j].luz)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func geraAmbiente()[10][10]Environment{
	var grid [10][10]Environment
	for i:= 0; i < 10; i++{
		for j:=0; j < 10; j++{
			grid[i][j].agua = r.Intn(10) + 1
			grid[i][j].calor = r.Intn(10) + 1
			grid[i][j].gases = r.Intn(10) + 1
			grid[i][j].luz = r.Intn(10) + 1

		}
	}
	return grid
}

func geraVida()[10][10]int{
	var grid [10][10]int
	for i:= 0; i < 10; i++{
		for j:=0; j < 10; j++{
			if r.Intn(3) == 1{
				grid[i][j] = 1
			}
		}
	}
	return grid
}

func main(){
	var nomeEspecies []string
	catalogoEspecies := make(map[string]Environment)
	//ambiente := geraAmbiente()
	//vida := geraVida()
	//especie := geraEspecie()
	//mostraIndividuo(especie)
	//especie = geraVariabilidade(especie)
	//mostraIndividuo(especie)
	//for _, valor := range vida{
	//	fmt.Println(valor)
	//}
	//fmt.Printf("\n")
	//mostraAgua(ambiente)
	//mostraCalor(ambiente)
	//mostraGases(ambiente)
	//mostraLuz(ambiente)
}
