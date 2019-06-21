package main

import(
    "time"
    "math/rand"
)

func sorteia_especie() string {//Sorteia um tipo de especie com maior probabilidade para plantas
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    tipo := 0

    tipo = r1.Intn(10)
    if(tipo >= 0 && tipo <= 4){
        return "F"//fotossintetico
    }else if(tipo >= 5 && tipo <= 7){
        return "H"//Herbivoro
    }else {
        return "C"//carnivoro
    }

}

func alimentacao(indice_i int, indice_j int, tipo string) {
    if(indice_i == 0){
        indice_i++
    }
    if(indice_j == 0){
        indice_j++
    }
    if(indice_i == 110){
        indice_i--
    }
    if(indice_j == 146){
        indice_j--
    }
    if(tipo == "H"){
        for i := indice_i-1; i < indice_i+1; i++ {
            for j:= indice_j-1; j < indice_j+1; j++ {
                if((grid[i][j].viva == true) && (grid[i][j].tipo_Cel == "F")){
                    grid[i][j].viva = false
                    grid[indice_i][indice_j].time_max.Add(time.Second * 5)
                    j = indice_j+1
                    i = indice_i+1
                }
            }
        }
    } else if (tipo == "C"){
        for i := indice_i-1; i < indice_i+1; i++ {
            for j:= indice_j-1; j < indice_j+1; j++ {
                if((grid[i][j].viva == true) && (grid[i][j].tipo_Cel == "H")){
                    grid[i][j].viva = false
                    grid[indice_i][indice_j].time_max.Add(time.Second * 5)
                    j = indice_j+1
                    i = indice_i+1
                }
            }
        }

    }
}
