package main

import (
	"fmt"
	"strings"
)

type shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

type shBuilder struct {
	code string
}

// NewShieldBuilder : Constructor for a new shield builder
func NewShieldBuilder() *shBuilder {
	return new(shBuilder)
}

func (sh *shBuilder) RaiseFront() *shBuilder {
	sh.code += "F"
	return sh
}

func (sh *shBuilder) RaiseBack() *shBuilder {
	sh.code += "B"
	return sh
}

func (sh *shBuilder) RaiseRight() *shBuilder {
	sh.code += "R"
	return sh
}

func (sh *shBuilder) RaiseLeft() *shBuilder {
	sh.code += "L"
	return sh
}

func (sh *shBuilder) Build() *shield {
	code := sh.code
	return &shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		right: strings.Contains(code, "R"),
		left:  strings.Contains(code, "L"),
	}
}

func main() {
	builder := NewShieldBuilder()
	shield := builder.RaiseLeft().RaiseFront().Build()
	fmt.Printf("%+v \n", *shield)
}
