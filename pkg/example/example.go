package example

import "github.com/charlesakalugwu/codecov/pkg/awesome"

var result string

func Setup() {
	result = awesome.Smile()
}

func Result() string {
	return result
}
