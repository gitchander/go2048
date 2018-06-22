package go2048

import (
	"bytes"
	"image"
)

type cellContenter interface {
	Size() image.Point
	CellValue(cell image.Point) (val int, ok bool)
}

type dummyCellContenter image.Point

func (c dummyCellContenter) Size() image.Point {
	return image.Point(c)
}

func (c dummyCellContenter) CellValue(cell image.Point) (val int, ok bool) {
	return 0, false
}

func encodePrintableTest(g *grid) []byte {
	return encodePrintable(g, "\t", BorderTable(5))
}

func encodePrintable(cc cellContenter, prefix string, ssr [][]rune) []byte {

	var cellSize = image.Point{X: 5, Y: 1}

	pe := newPrintEncoder(cellSize, prefix, cc, ssr)

	return pe.Encode()
}

type printEncoder struct {
	gridSize image.Point
	cellSize image.Point
	prefix   string
	cc       cellContenter
	ssr      [][]rune
}

func newPrintEncoder(cellSize image.Point, prefix string,
	cc cellContenter, ssr [][]rune) *printEncoder {

	if cellSize.X < 1 {
		cellSize.X = 1
	}
	if cellSize.Y < 1 {
		cellSize.Y = 1
	}

	gridSize := cc.Size()

	return &printEncoder{
		gridSize: gridSize,
		cellSize: cellSize,
		prefix:   prefix,
		cc:       cc,
		ssr:      ssr,
	}
}

func (pe *printEncoder) Encode() []byte {

	var buf = new(bytes.Buffer)

	yn := pe.gridSize.Y

	var (
		cellHeight = pe.cellSize.Y

		beforeHeight = (cellHeight - 1) / 2
		afterHeight  = (cellHeight - 1) - beforeHeight

		ssr = pe.ssr
	)

	for y := 0; y < yn; y++ {

		if y == 0 {
			pe.writeLine(buf, ssr[0])
		}

		for i := 0; i < beforeHeight; i++ {
			pe.writeLine(buf, ssr[1])
		}

		pe.writeLineVal(buf, ssr[1], y)

		for i := 0; i < afterHeight; i++ {
			pe.writeLine(buf, ssr[1])
		}

		if y < yn-1 {
			pe.writeLine(buf, ssr[2])
		} else {
			pe.writeLine(buf, ssr[3])
		}
	}

	return buf.Bytes()
}

func (pe *printEncoder) writeLine(buf *bytes.Buffer, sr []rune) {
	var (
		xn        = pe.gridSize.X
		cellWidth = pe.cellSize.X
	)

	buf.WriteString(pe.prefix)
	for x := 0; x < xn; x++ {
		if x == 0 {
			buf.WriteRune(sr[0])
		}
		for i := 0; i < cellWidth; i++ {
			buf.WriteRune(sr[1])
		}
		if x < xn-1 {
			buf.WriteRune(sr[2])
		} else {
			buf.WriteRune(sr[3])
		}
	}
	buf.WriteByte('\n')
}

func (pe *printEncoder) writeLineVal(buf *bytes.Buffer, sr []rune, y int) {
	var (
		xn        = pe.gridSize.X
		cellWidth = pe.cellSize.X
	)

	buf.WriteString(pe.prefix)
	for x := 0; x < xn; x++ {
		if x == 0 {
			buf.WriteRune(sr[0])
		}

		if val, ok := pe.cc.CellValue(image.Point{x, y}); ok {
			if cellWidth >= 4+2 {
				buf.WriteRune(sr[1])
				buf.WriteString(ItoaN(val, cellWidth-2, byte(sr[1])))
				buf.WriteRune(sr[1])
			} else {
				buf.WriteString(ItoaN(val, cellWidth, byte(sr[1])))
			}
		} else {
			for i := 0; i < cellWidth; i++ {
				buf.WriteRune(sr[1])
			}
		}

		if x < xn-1 {
			buf.WriteRune(sr[2])
		} else {
			buf.WriteRune(sr[3])
		}
	}
	buf.WriteByte('\n')
}

func repeatRune(r rune, n int) string {
	rs := make([]rune, n)
	for i := 0; i < n; i++ {
		rs[i] = r
	}
	return string(rs)
}

func ItoaN(x, n int, fillByte byte) string {

	data := make([]byte, n)
	i := n

	if (x == 0) && (i > 0) {
		data[i-1] = '0'
		i--
	}

	for (x > 0) && (i > 0) {
		quo, rem := quoRem(x, 10)
		x = quo
		data[i-1] = byte(rem + '0')
		i--
	}

	for i > 0 {
		data[i-1] = fillByte
		i--
	}

	return string(data)
}
