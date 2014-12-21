package datasack


// Data returns a channel that lets you iterate over all the numbers in the datasack.
func (me *T) Data() <-chan float64 {

	// Create channel.
	//
	// We will return this to the caller.
	//
	// We will also spawn a goroutine and output the data from this datasack has onto it.
	//
		out := make(chan float64)

	// Spawn a goroutine that will output the data from this datasack onto the channel
	// we previously created.
	//
	// Note that this goroutine will probably block. But that's OK, since it is in
	// its own goroutine (and shouldn't take anything else down with it).
	//
		go func() {
			for _,value := range me.slice {
				out <- value
			}

			close(out)
		}()

	// Return.
		return out
}
