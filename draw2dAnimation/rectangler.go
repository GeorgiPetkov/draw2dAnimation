package draw2dAnimation

type Rectangler interface {
	Figurer
	GetWidth() float64
	SetWidth(value float64)
	GetHeight() float64
	SetHeight(value float64)
}