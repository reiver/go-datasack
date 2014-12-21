package datasack


// Len returns the number of numbers added to the Datasack.
func (me *T) Len() int {
	return len(me.slice)
}
