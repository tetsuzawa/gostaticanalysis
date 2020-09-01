package a

import (
	"os"
)

func f1() {
	os.Open("") //want "NG"

	_, _ = os.Open("") //want "NG"

}

func f2() {
	f, _ := os.Open("_xxx") //OK
	defer f.Close()
	f.WriteString("aaaa")
	return
}

func f3() {
	f, _ := os.Open("_xxx") //OK
	f.WriteString("aaaa")
	defer f.Close()
	return //OK
}

func f4() {
	f, _ := os.Open("_xxx") //want "NG"
	f.WriteString("aaaa")
	return //OK
}
