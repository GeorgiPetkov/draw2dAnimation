package draw2dAnimation

// A flag enumeration representing the type of the update for a figure. Setted flags mean that this type of update will be performed on the current figure.
type updateType int

const (
	None        updateType = iota
	Translation updateType = 1 << iota
	Rotation    updateType = 1 << iota
	Custom      updateType = 1 << iota
)
