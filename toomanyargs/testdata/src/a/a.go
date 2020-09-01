package a

func f1(a int, b int, c int, d bool, e string, f uint) { // want "too many number of args in the func declaration. should be less than 5"
}

func f2(a, b, c, d, e, f int) { // want "too many number of args in the func declaration. should be less than 5"
}

func f3(a, b string, c, d, e, f int) { // want "too many number of args in the func declaration. should be less than 5"
}
