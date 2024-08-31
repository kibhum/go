package types

type Color int

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "BLUE"
	case ColorBlack:
		return "BLACK"
	case ColorYellow:
		return "YELLOW"
	case ColorPink:
		return "Pink"
	default:
		panic("unsupported color")
	}
}

func getColor() Color {
	return ColorPink
}