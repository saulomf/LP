package main

import(
        //"fmt"
        "time"
        "math/rand"
)

//Inicia cada celula do grid
func IniciaGrid(N int,M int, grid [][]Celula){
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    mutacao := -1
    eixoX, eixoY := 1.0, 1.0

    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {

            grid[i][j].posX=eixoX
            grid[i][j].posY=eixoY
            eixoX+=6.0
            //Seta a vida da celula com 5% de chance dela nascer
            if(r1.Intn(40) == 1){
                grid[i][j].viva = true
                grid[i][j].tipo_Cel = sorteia_especie()
                st := time.Now()
                s := st.Add(time.Second * 5)
                if(grid[i][j].tipo_Cel == "F"){
                    s = s.Add(time.Second * 5)
                }
                grid[i][j].time_vida = st
                grid[i][j].time_max = s

                //Caso a celula esteja viva sorteia com 8% de chance cada gene de mutacao
                for l := 0; l < 6; l++ {
                    mutacao = r1.Intn(25)
                    switch l {
                        case 0:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.agua = true
                            }
                            break
                        case 1:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.calor = true
                            }
                            break
                        case 2:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.frio = true
                            }
                            break
                        case 3:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.altitude = true
                            }
                            break
                        case 4:
                            if((mutacao >= 0) && (mutacao <=1)){
                                grid[i][j].mutacoes.comida = true
                            }
                    }
                }
            }else {
                grid[i][j].viva = false
            }

        }//end j

        eixoX=1.0
        eixoY+=6.0
    }//end i
}
