package main

func conway_game (vivos int, vida bool) bool {
	if(vivos == 3) {
            	vida = true //Celula nasce se morta ou permanece se viva
        } else if(vivos < 2) {
		vida = false // Morre de solidao se viva		
	} else if(vivos > 3) {
		vida = false // Morre de superpopulacao se viva
   	}
	return vida
}

func rule_conway () Rule {
	var rule Rule
	// 000 000 000 até 111 111 111
	for domain := 0; domain < 512; domain++ {
		vivos := 0	
		for i := domain; i > 0 ; i>>=1 {
			if(i&0x1 == 1){
				vivos++
			}	
		}	
		
		// xxx x#x xxx -> # = celula em si	
		if(domain&0x010 == 0) {
			rule[domain] = conway_game(vivos,false)	
		} else {
			rule[domain] = conway_game(vivos-1,true)	
		}
	}	
	
	return rule
}
