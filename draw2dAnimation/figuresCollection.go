package draw2dAnimation

import (
	"code.google.com/p/go-avltree/trunk"
)

// Used as compare method for type int.
func compareInts(first int, second int) int {
	if first < second {
		return -1
	} else if first > second {
		return 1
	}

	return 0
}

// Used as compare method for interface Figurer
func compareFigures(first interface{}, second interface{}) int {
	firstFigure, secondFigure := first.(Figurer), second.(Figurer)
	compareResult := compareInts(firstFigure.GetDepth(), secondFigure.GetDepth())
	if compareResult != 0 {
		return compareResult
	}

	return compareInts(firstFigure.getId(), secondFigure.getId())
}

// used as inner collection of type ComposedFigure
type figuresCollection struct {
	tree       *avltree.Tree
	dictionary map[string]Figurer
}

// Initializes new empty collection of figures.
func newFiguresCollection() *figuresCollection {
	return &figuresCollection{
		avltree.New(compareFigures, avltree.AllowDuplicates),
		make(map[string]Figurer)}
}

// Adds figure with string key to the collection.
func (this *figuresCollection) add(name string, figure Figurer) {
	this.tree.Add(figure)
	this.dictionary[name] = figure
}

// Removes figure by string key from the collection.
func (this *figuresCollection) remove(name string) {
	figure := this.dictionary[name]
	delete(this.dictionary, name)
	this.tree.Remove(figure)
}

// Gets the figure corresponding to the given string key in the collection or nil if not found.
func (this *figuresCollection) getByName(name string) Figurer {
	return this.dictionary[name]
}

// Traverse the collection calling a given function with each one of the figures.
func (this *figuresCollection) traverse(operation func(Figurer)) {
	for _, figure := range this.tree.Data() {
		operation(figure.(Figurer))
	}
}
