package draw2dAnimation

const (
	Unlimitted int = -1
)

// Type representing figure's update method than can set properties like expiration time, wait until starting, how often to update and a set of functions used for updating.
type UpdateMethod struct {
	times int
	waitTimes int
	eachNth bool
	eachNthValue int
	funcs []func(Figurer)
}

// Creates new update method by given set of update functions.
func NewUpdateMethod(funcs ...func(Figurer)) *UpdateMethod {
		return newUpdateMethod(funcs)
}

// Creates new update method by given times to update and set of update functions.
func NewUpdateMethod2(times int, funcs ...func(Figurer)) *UpdateMethod {
	return newUpdateMethod2(times, funcs)
}

// Creates new update method by given times to update, wait time before updating and set of update functions.
func NewUpdateMethod3(times int, waitTimes int, funcs ...func(Figurer)) *UpdateMethod {
	return newUpdateMethod3(times, waitTimes, funcs)
}

// Creates new update method by given times to update, wait time before updating, how often to update and set of update functions.
func NewUpdateMethod4(times int, waitTimes int, eachNth int, funcs ...func(Figurer)) *UpdateMethod {
	return newUpdateMethod4(times, waitTimes, eachNth, funcs)
}

// Creates new update method by given set of update functions.
func newUpdateMethod(funcs []func(Figurer)) *UpdateMethod {
		return newUpdateMethod2(Unlimitted, funcs)
}

// Creates new update method by given times to update and set of update functions.
func newUpdateMethod2(times int, funcs []func(Figurer)) *UpdateMethod {
	return newUpdateMethod3(times, 0, funcs)
}

// Creates new update method by given times to update, wait time before updating and set of update functions.
func newUpdateMethod3(times int, waitTimes int, funcs []func(Figurer)) *UpdateMethod {
	return newUpdateMethod4(times, waitTimes, 0, funcs)
}

// Creates new update method by given times to update, wait time before updating, how often to update and set of update functions.
func newUpdateMethod4(times int, waitTimes int, eachNth int, funcs []func(Figurer)) *UpdateMethod {
	return newUpdateMethod5(times, waitTimes, eachNth == 0, eachNth, funcs)
}

// Creates new update method by given times to update, wait time before updating, how often to update and set of update functions.
func newUpdateMethod5(times int, waitTimes int, eachNth bool, eachNthValue int, funcs []func(Figurer)) *UpdateMethod {
	return &UpdateMethod {
		times,
		waitTimes,
		eachNth,
		eachNthValue,
		funcs}
}

func (this *UpdateMethod) Update(figure Figurer) {
	if (this.times == 0) {
		figure.SetUpdateMethod(nil)
		return
	}
	
	if (this.waitTimes > 0) {
		this.waitTimes--
		return
	}
	
	for _, function := range this.funcs {
		function(figure)
	}
	
	if this.times != Unlimitted {
		this.times--
	}
	
	if this.eachNth {
		this.waitTimes = this.eachNthValue - 1
		return
	}
}