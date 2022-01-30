package gobasic

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func variableName(name string) string {
	return name
}

func variableConversionIntToString(from int) string {
	var result string = strconv.Itoa(from)
	return result
}

func TestVariable(t *testing.T) {
	fmt.Println(variableName("erik"))
}

func TestVariableConversionIntToString(t *testing.T) {
	varname := variableConversionIntToString(1)
	assert.Equal(t, "1", varname)
}
