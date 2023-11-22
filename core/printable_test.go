package core

import (
	"bytes"
	"fmt"
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

	cc := dummyCellContenter(MakePoint(3, 3))

	var cellSize = MakePoint(7, 3)

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
