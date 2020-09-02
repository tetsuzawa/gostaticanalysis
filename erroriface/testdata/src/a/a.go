package a

import "fmt"

type MyInt int

type MyError struct {
	Name string
	error
}

func (m MyError) Error() string {
	return fmt.Sprintf("%v error", m.Name)
}

type E2 struct {
	MyError
}

type E3 string

func (e E3) Error() string {
	return string(e)
}
