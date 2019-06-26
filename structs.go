package main

type Mundo struct {
	Planta	   [][]Planta
	Herbivoro  [][]Herbivoro
	Carnivoro  [][]Carnivoro
	Agua	   [][]Agua
	Ar	   [][]Ar
	Relevo     [][]Relevo
}

type Agua int
type Relevo int

type Planta struct {
	a int

}

type Herbivoro struct {
	a int

}

type Carnivoro struct {
	a int

}


type Ar struct {
	a int

}

type Output struct {
	a int

}