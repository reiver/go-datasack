package datasack


import "errors"
import "math"


// Max return the current maximum number the Datasack has "seen".
// This could change as new numbers are added to the Datasack.
func (me *T) Max() (float64, error) {
	if 0 == me.Len() {
		return math.Inf(-1), errors.New("The maximum is undefined when the datasack is empty.")
	}

	return me.max, nil
}

// proposeMax checks to see if we have a new max.
func (me *T) proposeMax(number float64) {

	if number > me.max {
		me.max = number
	}
}
