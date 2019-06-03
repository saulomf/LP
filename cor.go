package main


//Retorna a cor onde determinada celula esta
func cor_pixel (X int, Y int) string {

    if Y < 36 || Y > 280 {//Se for Branco
        if X > 97 && X < 106 {
            if Y > 4 && Y < 17 {
                return "a"
            }
        } else if X > 104 && X < 324 {
            if Y > 0 && Y < 23 {
                return "a"
            }
        } else if X > 131 && X < 295 {
            if Y > 21 && Y < 32 {
                return "a"
            }
        } else if X > 174 && X < 257 {
            if Y > 30 && Y < 36 {
                return "a"
            }
        } else if X > 322 && X < 332 {
            if Y > 5 && Y < 16 {
                return "a"
            }
        } else if X > 162 && X < 263 {//////////////////////////////////
            if Y > 278 && Y < 285 {
                return "a"
            }
        } else if X > 121 && X < 303 {
            if Y > 284 && Y < 296 {
                return "a"
            }
        } else if X > 87 && X < 102 {
            if Y > 300 && Y < 326 {
                return "a"
            }
        } else if X > 100 && X < 321 {
            if Y > 294 && Y < 332 {
                return "a"
            }
        } else if X > 318 && X < 333 {
            if Y > 301 && Y < 323 {
                return "a"
            }
        }

    } else {//Se for Verde
        if X > 28 && X < 122 {
            if Y > 69 && Y < 96 {
                return "v"
            }
        } else if X > 21 && X < 106 {
            if Y > 60 && Y < 68 {
                return "v"
            }
        } else if X > 23 && X < 83 {
            if Y > 49 && Y < 62 {
                return "v"
            }
        } else if X > 17 && X < 110 {/////////////////////////////////
            if Y > 122 && Y < 163 {
                return "v"
            }
        } else if X > 25 && X < 108 {
            if Y > 163 && Y < 176 {
                return "v"
            }
        } else if X > 42 && X < 90 {
            if Y > 174 && Y < 201 {
                return "v"
            }
        } else if X > 49 && X < 74 {
            if Y > 199 && Y < 224 {
                return "v"
            }
        } else if X > 54 && X < 76 {
            if Y > 223 && Y < 233 {
                return "v"
            }
        } else if X > 63 && X < 73 {
            if Y > 231 && Y < 244 {
                return "v"
            }
        } else if X > 109 && X < 131 {
            if Y > 134 && Y < 162 {
                return "v"
            }
        } else if X > 185 && X < 335 {/////////////////////////////
            if Y > 96 && Y < 125 {
                return "v"
            }
        } else if X > 335 && X < 354 {
            if Y > 107 && Y < 127 {
                return "v"
            }
        } else if X > 238 && X < 330 {
            if Y > 123 && Y < 142 {
                return "v"
            }
        } else if X > 284 && X < 307 {
            if Y > 141 && Y < 152 {
                return "v"
            }
        } else if X > 217 && X < 312 {
            if Y > 82 && Y < 96 {
                return "v"
            }
        } else if X > 375 && X < 425 {///////////////////////////////////
            if Y > 73 && Y < 100 {
                return "v"
            }
        } else if X > 393 && X < 419 {
            if Y > 99 && Y < 112 {
                return "v"
            }
        } else if X > 184 && X < 246 {
            if Y > 174 && Y < 204 {
                return "v"
            }
        } else if X > 244 && X < 281 {
            if Y > 187 && Y < 203 {
                return "v"
            }
        } else if X > 191 && X < 283 {
            if Y > 202 && Y < 213 {
                return "v"
            }
        } else if X > 203 && X < 254 {
            if Y > 212 && Y < 234 {
                return "v"
            }
        } else if X > 253 && X < 266 {
            if Y > 212 && Y < 226 {
                return "v"
            }
        } else if X > 213 && X < 253 {
            if Y > 231 && Y < 253 {
                return "v"
            }
        } else if X > 235 && X < 250 {
            if Y > 232 && Y < 243 {
                return "v"
            }
        } else if X > 336 && X < 389 {////////////////////////////
            if Y > 179 && Y < 187 {
                return "v"
            }
        } else if X > 331 && X < 415 {
            if Y > 186 && Y < 227 {
                return "v"
            }
        } else if X > 331 && X < 361 {
            if Y > 226 && Y < 239 {
                return "v"
            }
        } else if X > 379 && X < 391 {
            if Y > 226 && Y < 241 {
                return "v"
            }
        }
    }
    return "b"
}
