package datasack


import "testing"

import "math"
import "math/rand"
import "time"


func TestMaxWithEmptyDatasack(t *testing.T) {

	// Create a new datasack.
		datasack := New()
		if nil == datasack {
			t.Errorf("Got back 'nil' when called datasack.New().")
			return
		}


	// Make sure that if we call Max() on a datasack with nothing in it, then we receive
	// an error.
		if max, err := datasack.Max(); nil == err {
			t.Errorf("Excepted Max() to return an error, but it did not. It returned a max of [%v]", max)
		}

}


func TestMax(t *testing.T) {

	// Seed the random number generator.
	//
	// It is import to seed the random number generator, so that
	// we get different random number each time we run this test.
	//
	// If we don't seed it, then we get the same "random" numbers
	// over and over again!
	//
		rand.Seed( time.Now().UTC().UnixNano() )

	// We will do 10 tests. ("10" was arbitrarily chosen.)
	//
	// These tests contain a degree of randomness.
	// Thus, theoretically each time you run these tests, you could get a different result.
	// This is being done because there are too many possiblities to test everything.
	//
	// So... for each test...
		for test_number:=0; test_number<10; test_number++ {

			// Create a new datasack.
				datasack := New()
				if nil == datasack {
					t.Errorf("Got back 'nil' when called datasack.New().")
					return
				}


			// Figure out how many numbers we are going to add to the datasack for this test.
			//
			// This is randomly chosen.
			//
			// Note that this cannot be zero. (We deal with that case in another test.)
				number_of_numbers := 1 + rand.Intn(127)

			// Initialize the variable that will hold the max to negative infinity.
			//
			// If this still holds negative infinity at the end of this, there is something wrong.
				max := math.Inf(-1)

			// Generate "number_of_numbers" random numbers and put them in the datasack.
			//
			// Also figure out what the max should be.
				for i:=0; i<number_of_numbers; i++ {

					// Generate a new random number.
					//
					// This random number will (eventually) get added to the datasack.
						random_number := rand.Float64()
						if 0 == rand.Intn(5) {
							random_number *= 67108864.0
						}

					// Figure out the max.
					//
					// Check to see if this new random number is the new max.
					// If not, keep the old max.
						if random_number > max {
							max = random_number
						}

					// Add the new random number of the datasack.
						datasack.Insert(random_number)
				}

			// Make that "max" isn't still negative infinity.
				if math.Inf(-1) == max {
					t.Errorf("Our expected max was initialized to -∞ [%v]. However, we expected it be changed to something else.", max)
				}

			// Make sure the max that the datasack tells us is what we expect.
				datasack_max, err := datasack.Max()
				if nil != err {
					t.Errorf("Did not except Max() to return an error, but it did: [%v]", err)
				}
				if max != datasack_max {
					t.Errorf("The datasack.Max() [%v] does NOT equal the expected max [%v]", datasack_max, max)
				}
	}
}
