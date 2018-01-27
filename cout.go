package go2048

import (
	"bytes"
	"fmt"
	"image"
	"strings"
)

func coutMatrix(m *grid) {

	s := m.Size()
	dx := s.X
	dy := s.Y

	buf := new(bytes.Buffer)

	d := 7

	horisBorder := strings.Repeat("═", d)
	emptyFill := strings.Repeat(" ", d)

	for y := 0; y < dy; y++ {

		if y == 0 {
			for x := 0; x < dx; x++ {
				if x == 0 {
					buf.WriteRune('╔')
				} else {
					buf.WriteRune('╦')
				}
				buf.WriteString(horisBorder)
			}
			buf.WriteString("╗\n")
		}

		writeSticks(buf, dx, emptyFill)

		for x := 0; x < dx; x++ {

			if x == 0 {
				buf.WriteRune('║')
			}

			v := m.cellContent(image.Point{x, y})

			if v == nil {
				buf.WriteString(emptyFill)
			} else {
				/*
					if (newCell != nil) && ((newCell.X == x) && (newCell.Y == y)) {
						fmt.Fprintf(buf, ">%5d ", v)
					} else {
						fmt.Fprintf(buf, " %5d ", v)
					}
				*/
				fmt.Fprintf(buf, " %5d ", v.Value)
			}

			buf.WriteString("║")
		}
		buf.WriteByte('\n')

		writeSticks(buf, dx, emptyFill)

		if y < dy-1 {
			for x := 0; x < dx; x++ {
				if x == 0 {
					buf.WriteRune('╠')
				}
				buf.WriteString(horisBorder)
				if x < dx-1 {
					buf.WriteRune('╬')
				} else {
					buf.WriteRune('╣')
				}
			}
			buf.WriteRune('\n')
		} else {
			for x := 0; x < dx; x++ {
				if x == 0 {
					buf.WriteRune('╚')
				}
				buf.WriteString(horisBorder)
				if x < dx-1 {
					buf.WriteRune('╩')
				} else {
					buf.WriteRune('╝')
				}
			}
			buf.WriteRune('\n')
		}
	}

	fmt.Println(buf.String())
}

func writeSticks(buffer *bytes.Buffer, dx int, emptyFill string) {
	for x := 0; x < dx; x++ {
		if x == 0 {
			buffer.WriteRune('║')
		}
		buffer.WriteString(emptyFill)
		buffer.WriteRune('║')
	}
	buffer.WriteByte('\n')
}
