package main

import (
	"fmt"
	"strconv"
)

func main() {
}

type Stringer interface {
	String() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float:
		return strconv.Ftoa(v, 'g', -1)
	}
	return "???"
}

type error interface {
	Error() string
}

func New(text string) error {return &errorString{text}}

type errorString struct {
	text string
}

func (e *errorString) Error() string {
	return e.text
}