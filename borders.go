package go2048

/*

┌──────────────────┬───────┬─────────┐
│                  │ Best  │   10124 │
│  golang 2048     ├───────┼─────────┤
│                  │ Score │     544 │
└──────────────────┴───────┴─────────┘

╔══════════════════╦═══════╤═════════╗
║                  ║ Best  │   10124 ║
║  Go 2048         ╠═══════╪═════════╣
║                  ║ Score │     544 ║
╚══════════════════╩═══════╧═════════╝

*/

type borderRunes [][]rune

var b_table = []borderRunes{
	{
		[]rune("+-++"),
		[]rune("| ||"),
		[]rune("+-++"),
		[]rune("+-++"),
	},
	{
		[]rune("┌─┬┐"),
		[]rune("│ ││"),
		[]rune("├─┼┤"),
		[]rune("└─┴┘"),
	},
	{
		[]rune("╓─╥╖"),
		[]rune("║ ║║"),
		[]rune("╟─╫╢"),
		[]rune("╙─╨╜"),
	},
	{
		[]rune("╒═╤╕"),
		[]rune("│ ││"),
		[]rune("╞═╪╡"),
		[]rune("╘═╧╛"),
	},
	{
		[]rune("╔═╦╗"),
		[]rune("║ ║║"),
		[]rune("╠═╬╣"),
		[]rune("╚═╩╝"),
	},
	{
		[]rune("╔═╤╗"),
		[]rune("║ │║"),
		[]rune("╟─┼╢"),
		[]rune("╚═╧╝"),
	},
	{
		[]rune("┌─╥┐"),
		[]rune("│ ║│"),
		[]rune("╞═╬╡"),
		[]rune("└─╨┘"),
	},
}

//var tableRune1 = [][]rune{
//	[]rune("+-++"),
//	[]rune("| ||"),
//	[]rune("+-++"),
//	[]rune("+-++"),
//}

//var tableRune2 = [][]rune{
//	[]rune("┌─┬┐"),
//	[]rune("│ ││"),
//	[]rune("├─┼┤"),
//	[]rune("└─┴┘"),
//}

//var tableRune3 = [][]rune{
//	[]rune("╓─╥╖"),
//	[]rune("║ ║║"),
//	[]rune("╟─╫╢"),
//	[]rune("╙─╨╜"),
//}

//var tableRune4 = [][]rune{
//	[]rune("╒═╤╕"),
//	[]rune("│ ││"),
//	[]rune("╞═╪╡"),
//	[]rune("╘═╧╛"),
//}

//var tableRune5 = [][]rune{
//	[]rune("╔═╦╗"),
//	[]rune("║ ║║"),
//	[]rune("╠═╬╣"),
//	[]rune("╚═╩╝"),
//}

//var tableRune6 = [][]rune{
//	[]rune("╔═╤╗"),
//	[]rune("║ │║"),
//	[]rune("╟─┼╢"),
//	[]rune("╚═╧╝"),
//}

//var tableRune7 = [][]rune{
//	[]rune("┌─╥┐"),
//	[]rune("│ ║│"),
//	[]rune("╞═╬╡"),
//	[]rune("└─╨┘"),
//}

//var tableRune8 = [][]rune{
//	[]rune("+=++"),
//	[]rune("I.|I"),
//	[]rune("+-++"),
//	[]rune("+=++"),
//}

func BorderTable(i int) [][]rune {
	return ([][]rune)(b_table[mod(i, len(b_table))])
}

/*

╔═══════╦═══════╦═══════╦═══════╗
║  2048 ║       ║       ║       ║
╠═══════╬═══════╬═══════╬═══════╣
║       ║       ║       ║       ║
╠═══════╬═══════╬═══════╬═══════╣
║     4 ║       ║       ║     2 ║
╠═══════╬═══════╬═══════╬═══════╣
║       ║       ║       ║       ║
╚═══════╩═══════╩═══════╩═══════╝

*/