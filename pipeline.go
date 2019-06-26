package main

import ()

type Data struct{
	ReadUp 	   *SubMundo
	ReadDown   *SubMundo
	ReadLeft   *SubMundo
	ReadRight  *SubMundo
	Center 	   *SubMundo
}

func Reajusta_Mundo(Frame[escrita],I,J int) {
	Frame[escrita].Carnivoro = Frame[escrita].Carnivoro[:I]
	Frame[escrita].Herbivoro = Frame[escrita].Herbivoro[:I]
	Frame[escrita].Planta	 = Frame[escrita].Planta[:I]
	Frame[escrita].Agua	 = Frame[escrita].Agua[:I]
	Frame[escrita].Ar	 = Frame[escrita].Ar[:I]
	Frame[escrita].Relevo 	 = Frame[escrita].Relevo[:I]

	for i := range Frame[escrita].Ar {
		Frame[escrita].Carnivoro[i] 	= Frame[escrita].Carnivoro[i][:J]
		Frame[escrita].Herbivoro[i] 	= Frame[escrita].Herbivoro[i][:J]
		Frame[escrita].Planta[i]	= Frame[escrita].Planta[i][:J]
		Frame[escrita].Agua[i]	 	= Frame[escrita].Agua[i][:J]
		Frame[escrita].Ar[i]	 	= Frame[escrita].Ar[i][:J]
		Frame[escrita].Relevo[i] 	= Frame[escrita].Relevo[i][:J]
	}
}

func Estagio1(In <-chan Data, Out chan<- Data,Max int) {

	var Frame [2]Mundo
	Frame[0] = Aloca_Mundo(Max,Max)
	Frame[1] = Aloca_Mundo(Max,Max)

	escrita := 0
	leitura := 1

	for {
		// Recebe e executa	
		Data := <-In

		// Resize frame
		Reajusta_Mundo(Frame[escrita],len(Data.Center),len(Data.Center[0]))

		Atualiza_Carnivoros(Frame[escrita],Data)

		// Troca frame
		escrita ^= 1
		leitura ^= 1	

		// Proximo estagio
		Data.Center = Frame[leitura] 
		Out <- Data
	}
}

func Estagio2(In <-chan Data, Out chan<- Data,Max int) {

	var Frame [2]Mundo
	Frame[0] = Aloca_Mundo(Max,Max)
	Frame[1] = Aloca_Mundo(Max,Max)

	escrita := 0
	leitura := 1
	
	for {	
		// Recebe e executa	
		Data := <-In

		// Resize frame
		Reajusta_Mundo(Frame[escrita],len(Data.Center),len(Data.Center[0]))

		Atualiza_Herbivoros(Frame[escrita],Data)

		// Troca frame
		escrita ^= 1
		leitura ^= 1	

		// Proximo estagio
		Data.Center = Frame[leitura] 
		Out <- Data
	}	
}

func Estagio3(In <-chan Data, Out chan<- Data,Max int) {

	var Frame [2]Mundo
	Frame[0] = Aloca_Mundo(Max,Max)
	Frame[1] = Aloca_Mundo(Max,Max)

	escrita := 0
	leitura := 1
	
	for { 
		// Recebe e executa	
		Data := <-In

		// Resize frame
		Reajusta_Mundo(Frame[escrita],len(Data.Center),len(Data.Center[0]))

		Atualiza_Plantas(Frame[escrita],Data)

		// Troca frame
		escrita ^= 1
		leitura ^= 1	

		// Proximo estagio
		Data.Center = Frame[leitura] 
		Out <- Data
	}		
}

func Estagio4(In <-chan Data, Out chan<- Data,Max int) {

	var Frame [2]Mundo
	Frame[0] = Aloca_Mundo(Max,Max)
	Frame[1] = Aloca_Mundo(Max,Max)

	escrita := 0
	leitura := 1
	
	for { 
		// Recebe e executa	
		Data := <-In

		// Resize frame
		Reajusta_Mundo(Frame[escrita],len(Data.Center),len(Data.Center[0]))

		Atualiza_Agua(Frame[escrita],Data)

		// Troca frame
		escrita ^= 1
		leitura ^= 1	

		// Proximo estagio
		Data.Center = Frame[leitura] 
		Out <- Data
	}	
}

func Estagio5(In <-chan Data, Out chan<- Data,Max int) {

	var Frame [2]Mundo
	Frame[0] = Aloca_Mundo(Max,Max)
	Frame[1] = Aloca_Mundo(Max,Max)

	escrita := 0
	leitura := 1
	
	for { 
		// Recebe e executa	
		Data := <-In

		// Resize frame
		Reajusta_Mundo(Frame[escrita],len(Data.Center),len(Data.Center[0]))

		Atualiza_Agua(Frame[escrita],Data)

		// Troca frame
		escrita ^= 1
		leitura ^= 1	

		// Proximo estagio
		Data.Center = Frame[leitura] 
		Out <- Data
	}	
}

func Estagio6(In <-chan Data, Out chan<- Data,Max int) {

	var Frame [2]Mundo
	Frame[0] = Aloca_Mundo(Max,Max)
	Frame[1] = Aloca_Mundo(Max,Max)

	escrita := 0
	leitura := 1
	
	for { 
		// Recebe e executa	
		Data := <-In

		// Resize frame
		Reajusta_Mundo(Frame[escrita],len(Data.Center),len(Data.Center[0]))

		Atualiza_Agua(Frame[escrita],Data)

		// Troca frame
		escrita ^= 1
		leitura ^= 1	

		// Proximo estagio
		Data.Center = Frame[leitura] 
		Out <- Data
	}	
}





var QueueIn chan Mundo
var QueueOut chan Mundo
// O grid possui bordas com numero negativo
func Start_Pipeline(n_estagios int,I,J int) {
	QueueIn  =  make(chan Mundo)
	Stage2In := make(chan Mundo)
	Stage3In := make(chan Mundo)
	Stage4In := make(chan Mundo)
	QueueOut =  make(chan Mundo)

	go Estagio1(QueueIn,Stage2In,I,J)	// Atualiza Carnivoros
	go Estagio2(Stage2In,Stage3In,I,J)	// Atualiza Herbivoros
	go Estagio3(Stage3In,Stage4In,I,J)	// Atualiza Vegetais
	go Estagio4(Stage4In,QueueOut,I,J)	// Atualiza Agua
}