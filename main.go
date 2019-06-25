package main

import(
        //"fmt"
        "time"
        "log"
        "github.com/hajimehoshi/ebiten"
)


func update(screen *ebiten.Image) error {
    if ebiten.IsDrawingSkipped() {
        return nil
    }
    screen.DrawImage(fundo, nil)
    op := &ebiten.DrawImageOptions{}
    //x = 0.0
    for i := 0; i < 110; i++ {
        //op.GeoM.Translate(0, 11)
        //screen.DrawImage(img, op)
        for j := 0; j < 146; j++ {
            alimentacao(i, j, grid[i][j].tipo_Cel)

            op.GeoM.Translate(grid[i][j].posX, grid[i][j].posY)
            op.GeoM.Scale(0.5, 0.5)

            tempo_max(i, j)
            if(grid[i][j].viva == true){

                if(grid[i][j].tipo_Cel == "F"){
                    screen.DrawImage(img, op)
                }else if(grid[i][j].tipo_Cel == "H"){
                    screen.DrawImage(img2, op)
                }else if(grid[i][j].tipo_Cel == "C"){
                    screen.DrawImage(img3, op)
                }
            }else{
                grid[i][j].viva = false
            }
            op.GeoM.Reset()
            //x-=11.0
        }
        //op.GeoM.Translate(x, 0)
        //x = 0.0
    }
    if(teste < 10){
        teste++
    }
    if(teste == 10){
        AtualizaGrid(110,146)
        // Modificações para testar que o grid se altera
        teste = 0
    }

    return nil
}

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
    time_vida time.Time
    time_max time.Time
    tipo_Cel string // H para herbivor, C para carnivoro, F para fotossintetico
}


func main(){

    for i := range grid {
            grid[i] = make([]Celula,200,200)
            copia_grid[i] = make([]Celula,200,200)
    }

    //Teste(110,146)
    IniciaGrid(110,146, grid)

    if err := ebiten.Run(update, 439, 331, 2, "Novo jogo da vida"); err != nil {
        log.Fatal(err)
    }
    //IniciaGrid(100,100,grid)
    //MostraGrid(100,100,grid)

}
