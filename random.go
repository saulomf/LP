package main

import (
	"math/rand"
	"time"
)


func rule_random () Rule {
    	s1 := rand.NewSource(time.Now().UnixNano())
    	r1 := rand.New(s1)
	var rule Rule
	// 000 000 000 até 111 111 111
	for domain := 0; domain < 512; domain++ {
		if(r1.Intn(2)==1){
			rule[domain] = true
		} else {
			 rule[domain] = false
		}
	}

	return rule
}
