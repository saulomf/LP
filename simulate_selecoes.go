package main

//import(
//	"github.com/hajimehoshi/ebiten"
//	"github.com/hajimehoshi/ebiten/inpututil"
//)



//func aloca_muxsimulacao(N int) muxSimulacao {
//	var muxsimulacao muxSimulacao
//	muxsimulacao.array = make([]Simulacao,N)
//	// Por padrao, i = 0
//	muxsimulacao.N = N
//
//	return muxsimulacao
//}

//func chose(simulacoes muxSimulacao){
//	for true {
//		if(ebiten.IsKeyPressed(ebiten.KeyLeft) && inpututil.KeyPressDuration(ebiten.KeyLeft) == 5) {
//			simulacoes.i = (simulacoes.i - 1 + simulacoes.N)%simulacoes.N							
//		} else if (ebiten.IsKeyPressed(ebiten.KeyRight) && inpututil.KeyPressDuration(ebiten.KeyRight) == 5) {
//			simulacoes.i = (simulacoes.i + 1)%simulacoes.N
//		}
//		tela.canal <- simulacoes.array[simulacoes.i].output
//	}
//}