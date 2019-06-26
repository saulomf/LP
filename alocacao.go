package main

func Aloca_arrayMundo(num,I,J int) []Mundo {

	arrayMundo := make([]Mundo,num,num)
	for _,linha:= range arrayMundo {
		linha = Aloca_Mundo(I,J)
	} 
	return arrayMundo
}

func Aloca_Mundo(I,J int) Mundo {
	var mundo Mundo	

	mundo.Planta	= Aloca_Planta(I,J)
	mundo.Herbivoro	= Aloca_Herbivoro(I,J)
	mundo.Carnivoro	= Aloca_Carnivoro(I,J)
	mundo.Agua	= Aloca_Agua(I,J)
	mundo.Ar	= Aloca_Ar(I,J)
	mundo.Relevo	= Aloca_Relevo(I,J)

	return mundo
}

func Aloca_Planta(I,J int) [][]Planta {
	Ret := make([][]Planta,I,I)
	for _,linha:= range Ret {
		linha = make([]Planta,J,J)
	}
	return Ret
}
func Aloca_Herbivoro(I,J int) [][]Herbivoro {
	Ret := make([][]Herbivoro,I,I)
	for _,linha:= range Ret {
		linha = make([]Herbivoro,J,J)
	}
	return Ret
}
func Aloca_Carnivoro(I,J int) [][]Carnivoro {
	Ret := make([][]Carnivoro,I,I)
	for _,linha:= range Ret {
		linha = make([]Carnivoro,J,J)
	}
	return Ret
}
func Aloca_Agua(I,J int) [][]Agua {
	Ret := make([][]Agua,I,I)
	for _,linha:= range Ret {
		linha = make([]Agua,J,J)
	}
	return Ret
}
func Aloca_Relevo(I,J int) [][]Relevo {
	Ret := make([][]Relevo,I,I)
	for _,linha:= range Ret {
		linha = make([]Relevo,J,J)
	}
	return Ret
}
func Aloca_Ar(I,J int) [][]Ar {
	Ret := make([][]Ar,I,I)
	for _,linha:= range Ret {
		linha = make([]Ar,J,J)
	}
	return Ret
}
func Aloca_Output(I,J int) [][]Output {
	Ret := make([][]Output,I,I)
	for _,linha:= range Ret {
		linha = make([]Output,J,J)
	}
	return Ret
}