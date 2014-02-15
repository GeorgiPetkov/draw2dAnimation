package draw2dAnimation

type UpdateType int

const (
	None        UpdateType = 0
	Translation UpdateType = 1
	Rotation    UpdateType = 2
	Custom      UpdateType = 4
)
