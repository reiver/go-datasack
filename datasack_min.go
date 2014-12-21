package datasack


import "errors"
import "math"


// Min return the current minimum number the Datasack has "seen".
// This could change as new numbers are added to the Datasack.
func (me *T) Min() (float64, error) {
	if 0 == me.Len() {
		return math.Inf(1), errors.New("The minimum is undefined when the datasack is empty.")
	}

	return me.min, nil
}

// proposeMin checks to see if we have a new min.
func (me *T) proposeMin(number float64) {

	if number < me.min {
		me.min = number
	}
}
