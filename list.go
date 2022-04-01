package main

//declare List
type list []*game

//declare method for List Print
func (l list) print() {
	for _, item := range l {
		item.print()
	}
}
