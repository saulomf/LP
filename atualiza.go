package main

import (
    "time"
    "math/rand"
)

/*type Old_State struct{
    i,j int
    vida bool
}

/*func Jogo(rule Rule,N int, M int, i int, j int, hold chan int, result chan Old_State) {
        //Verifica todos os vizinhos da celula
        domain := bitfy(i,j,N,M)

    //Abaixo estão sendo usada as regras do jogo da vida de conway, a ideia e tornar essas      regras flexiveis a outros valores
        var ret Old_State
        ret.i= i
        ret.j= j
        ret.vida = apply_rule(rule,domain)

        hold <- 1
        result <- ret
}*/

func Jogo(i int, j int)  {
    vivos, paiI, paiJ, maeI, maeJ := 0, -1, -1, -1, -1 //quantidade de vizinhos vivos
    //Verifica todos os vizinhos da celula
    if(i == 0){
        i++

    }
    if(j == 0){
        j++
    }
    if(i == 110){
        i--
    }
    if(j == 146){
        j--
    }
    for posI := i-1; posI < i+1; posI++ {
        for posJ := j-1; posJ < j+1; posJ++ {
            //fmt.Printf("\n%d %d", posI, posJ)

             if(grid[posI][posJ].viva == true){
                vivos++
                if(vivos == 1){
                    paiI = posI
                    paiJ = posJ
                }else if(vivos == 2){
                    maeI = posI
                    maeJ = posJ
                }
            }
        }
    }
    //Abaixo estão sendo usada as regras do jogo da vida de conway, a ideia e tornar essas regras flexiveis a outros valores
    if(grid[i][j].viva == true){
        vivos--
        if(vivos > 3) {
            grid[i][j].viva = false // Morre de superpopulacao se viva
        }

    }else{
        if(vivos == 2){
            if(grid[paiI][paiJ].tipo_Cel == grid[maeI][maeJ].tipo_Cel){
                grid[i][j].viva = true
                copia_grid[i][j].viva = true
                st := time.Now()
                s := st.Add(time.Second * 5)
                grid[i][j].time_vida = st
                grid[i][j].time_max = s
                grid[i][j].tipo_Cel = grid[paiI][paiJ].tipo_Cel
            }
        }else if(vivos == 0){//Impede a extinção
            s1 := rand.NewSource(time.Now().UnixNano())
            r1 := rand.New(s1)
            if(r1.Intn(50) == 1){
                grid[i][j].viva = true
                copia_grid[i][j].viva = true
                st := time.Now()
                s := st.Add(time.Second * 5)
                grid[i][j].time_vida = st
                grid[i][j].time_max = s
                grid[i][j].tipo_Cel = sorteia_especie()
            }
        }
    }
}

func AtualizaGrid(N int, M int){
    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {
            Jogo(i, j)
        }
    }
}

/*func AtualizaGrid(N int,M int){
        //cont := 0
        //continua := true

    conway := rule_conway()
    //conway := rule_random()

        //for continua != false {

            //MostraGrid(N,M,grid)

            //time.Sleep(100 * time.Millisecond)
    hold := make(chan int,0)
    result := make(chan Old_State)

    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {
            go Jogo(conway,N,M,i,j,hold,result)
        }
    }

    for i := 0; i < N*M; i++{
        <-hold
    }

    var new Old_State
    for i := 0; i < N*M; i++{
        new = <-result
        grid[new.i][new.j].viva = new.vida
        if(new.vida == true){
            st := time.Now()
            s := st.Add(time.Second * 5)
            grid[new.i][new.j].time_vida = st
            grid[new.i][new.j].time_max = s
            grid[new.i][new.j].tipo_Cel = sorteia_especie()
        }else{
            st := time.Now()
            grid[new.i][new.j].time_vida = st
            grid[new.i][new.j].time_max = st
        }
    }
}*/
