package main

type Color string

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	Gray   Color = "\033[37m"
	White  Color = "\033[97m"
)

func main() {
	sandglass(11, '@', Green)
}

func sandglass(size int, char rune, color Color) {
	sand := make([][]rune, size)
	sand[0] = makeRow(size, char)
	sand[size-1] = makeRow(size, char)
	half := (size + 1) / 2
	for i := 1; i < half; i++ {
		u := makeRow(size, ' ')
		d := makeRow(size, ' ')
		u[i] = char
		u[size-1-i] = char
		d[i] = char
		d[size-1-i] = char
		sand[i] = u
		sand[size-1-i] = d
	}

	for _, v := range sand {
		println(string(color) + string(v))
	}
	print(Reset)
}

func makeRow(size int, char rune) []rune {
	row := make([]rune, size)
	for i := range row {
		row[i] = char
	}
	return row
}
