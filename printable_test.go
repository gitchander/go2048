package go2048

import (
	"bytes"
	"fmt"
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

	var cellSize = image.Point{X: 7, Y: 3}

	pe := newPrintEncoder(cellSize, "\t", cc, tableRuneTest)
	data := pe.Encode()

	if !bytes.Equal(data, result) {
		fmt.Println(string(data))
		fmt.Println(string(result))

		t.Log(hex.Dump(data))
		t.Log(hex.Dump(result))

		t.Fatal("not equal")
	}
}
