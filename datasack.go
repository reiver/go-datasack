package datasack


import "math"


type T struct {
	slice []float64
	max     float64
	min     float64
}


// New returns a new datasack.
func New() (*T) {

	// There doesn't seem to be a literal for the +∞ and -∞ for float64 in Golang.
	// So we need to call these functions to get those values.
	//
	// We use these to initialize the "max" and "min" fields of the datasack.
	//
		positiveInfinity := math.Inf(1)
		negativeInfinity := math.Inf(-1)

	// Here we create a slice with a capacity of 1,024
	// and the length of 0 (zero).
	//
	// If we go past our capacity, it will automagically get extended.
	//
	// I'm not sure if this is a "hard" rule, but when I did a little testing
	// I found that the slice gets extended by additions of the original capacity.
	//
	// So, if your capacity is originally 1,024
	// then the capacity grows as:
	//
	//     1,024  (= 1 x 1,024)
	//     2,048  (= 2 x 1,024)
	//     3,072  (= 3 x 1,024)
	//     4,096  (= 4 x 1,024)
	//     5,120  (= 5 x 1,024)
	//     6,144  (= 6 x 1,024)
	//     7,168  (= 7 x 1,024)
	//      ...
	//
	// I do not know if Golang guarantees this behavior or not.
	//
	// This note was more of an interesting FYI.
	//
	// And that I probably don't need to write special code to extend the slice
	// myself (in a sane manner, since it already seems to be sane).
	//
	//
	// Also NOTE that a slice was chosen over something like a linked list, so that
	// this can take advantage of CPU caching of memory and things such as
	// cache lines.
	//
		new_slice := make([]float64, 0, 1024)

	// Create the datasack.
		me := T{
			slice: new_slice,
			max: negativeInfinity,
			min: positiveInfinity,
		}

	// Return.
		return &me
}
