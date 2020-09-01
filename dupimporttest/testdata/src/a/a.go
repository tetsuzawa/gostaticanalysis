package a

import fmt1 "fmt"
import fmt2 "fmt" // want "fmt is duplicated import"

func f() {
	fmt1.Println("Hello")
	fmt2.Println("World")
}
