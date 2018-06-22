package go2048

import (
	"bytes"
	"image"
	"testing"

	"encoding/hex"
)

func TestPrintable(t *testing.T) {

	var tableRuneTest = [][]rune{
		[]rune("ABCD"),
		[]rune("EFGH"),
		[]rune("IJKL"),
		[]rune("MNOP"),
	}

	result := []byte("\tABBBBBBBCBBBBBBBCBBBBBBBD\n" +
		"\tEFFFFFFFGFFFFFFFGFFFFFFFH\n" +
		"\tEFFFFFFFGFFFFFFFGFFFFFFFH\n" +
		"\tEFFFFFFFGFFFFFFFGFFFFFFFH\n" +
		"\tIJJJJJJJKJJJJJJJKJJJJJJJL\n" +
		"\tEFFFFFFFGFFFFFFFGFFFFFFFH\n" +
		"\tEFFFFFFFGFFFFFFFGFFFFFFFH\n" +
		"\tEFFFFFFFGFFFFFFFGFFFFFFFH\n" +
		"\tIJJJJJJJKJJJJJJJKJJJJJJJL\n" +
		"\tEFFFFFFFGFFFFFFFGFFFFFFFH\n" +
		"\tEFFFFFFFGFFFFFFFGFFFFFFFH\n" +
		"\tEFFFFFFFGFFFFFFFGFFFFFFFH\n" +
		"\tMNNNNNNNONNNNNNNONNNNNNNP\n")

	cc := dummyCellContenter(image.Point{X: 3, Y: 3})

	data := encodePrintable(cc, "\t", tableRuneTest)

	if !bytes.Equal(data, result) {
		t.Log(hex.Dump(data))
		t.Log(hex.Dump(result))

		t.Fatal("not equal")
	}
}
