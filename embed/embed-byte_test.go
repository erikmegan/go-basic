package embed

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed laptop.jpg
var laptop []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("laptop_new.jpg", laptop, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
