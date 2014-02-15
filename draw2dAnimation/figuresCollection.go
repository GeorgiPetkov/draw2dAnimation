package draw2dAnimation

import (
	"code.google.com/p/go-avltree/trunk"
)

func compareInts(first int, second int) int {
	if first < second {
		return -1
	} else if first > second {
		return 1
	}

	return 0
}

func compareFigures(first interface{}, second interface{}) int {
	firstFigure, secondFigure := first.(Figurer), second.(Figurer)
	compareResult := compareInts(firstFigure.GetDepth(), secondFigure.GetDepth())
	if compareResult != 0 {
		return compareResult
	}

	return compareInts(firstFigure.getId(), secondFigure.getId())
}

type figuresCollection struct {
	tree       *avltree.Tree
	dictionary map[string]Figurer
}

func newFiguresCollection() *figuresCollection {
	return &figuresCollection{
		avltree.New(compareFigures, avltree.AllowDuplicates),
		make(map[string]Figurer)}
}

func (this *figuresCollection) add(name string, figure Figurer) {
	this.tree.Add(figure)
	this.dictionary[name] = figure
}

func (this *figuresCollection) remove(name string) {
	figure := this.dictionary[name]
	delete(this.dictionary, name)
	this.tree.Remove(figure)
}

func (this *figuresCollection) getByName(name string) Figurer {
	return this.dictionary[name]
}

func (this *figuresCollection) traverse(operation func(Figurer)) {
	for _, figure := range this.tree.Data() {
		operation(figure.(Figurer))
	}
}
