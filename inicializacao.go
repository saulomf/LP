package main

func Inicializa_Mundo(I,J int) Mundo {
	mundo := Aloca_Mundo(I,J)
	
	for i:= range mundo {
		for j:= range mundo[i] {
			mundo.Relevo[i][j]= sencons(i,j)
			mundo.Agua[i][j]= 20*sencons(i+5,j+5)
		} 
	}

	return mundo
}
/*
func sencons (i,j int) int {
	// ret(x,y)=a:z=1+sen(sqrt(x^(2)+y^(2)))
	
	if (i!=0 && j!=0 && i!=cell_num_y-1 && j!=cell_num_x-1) {
		x,y := float64(j),float64(i)
		x = (x/10)-cell_size_x/2
		y = (y/10)-cell_size_y/2
		f := math.Sin(math.Sqrt(x*x+y*y))
		f = (f+1)*2 

		return int(math.Min(altitude_max,f))
	} else {
		return altitude_max
	}
}
*/