package datasack


// Insert lets you add another number of the datasack.
func (me *T) Insert(number float64) {

	// Add the new number to the slice.
		me.slice = append(me.slice, number)

	// Check to see if we have a new max.
		me.proposeMax(number)

	// Check to see if we have a new min.
		me.proposeMin(number)
}
