package golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var resultString string = Length[string]("Masred")
	assert.Equal(t, "Masred", resultString)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}
