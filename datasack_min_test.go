package datasack


import "testing"

import "math"
import "math/rand"


func TestMinWithEmptyDatasack(t *testing.T) {

	// Create a new datasack.
		datasack := New()
		if nil == datasack {
			t.Errorf("Got back 'nil' when called datasack.New().")
			return
		}


	// Make sure that if we call Min() on a datasack with nothing in it, then we receive
	// an error.
		if min, err := datasack.Min(); nil == err {
			t.Errorf("Excepted Min() to return an error, but it did not. It returned a min of [%v]", min)
		}

}


func TestMin(t *testing.T) {


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

			// Initialize the variable that will hold the min to positive infinity.
			//
			// If this still holds negative infinity at the end of this, there is something wrong.
				min := math.Inf(1)

			// Generate "number_of_numbers" random numbers and put them in the datasack.
			//
			// Also figure out what the min should be.
				for i:=0; i<number_of_numbers; i++ {

					// Generate a new random number.
					//
					// This random number will (eventually) get added to the datasack.
						random_number := rand.Float64()
						if 0 == rand.Intn(5) {
							random_number *= 67108864.0
						}

					// Figure out the min.
					//
					// Check to see if this new random number is the new min.
					// If not, keep the old min.
						if random_number < min {
							min = random_number
						}

					// Add the new random number of the datasack.
						datasack.Insert(random_number)
				}

			// Make that "min" isn't still negative infinity.
				if math.Inf(1) == min {
					t.Errorf("Our expected min was initialized to +∞ [%v]. However, we expected it be changed to something else.", min)
				}

			// Make sure the min that the datasack tells us is what we expect.
				datasack_min, err := datasack.Min()
				if nil != err {
					t.Errorf("Did not except Min() to return an error, but it did: [%v]", err)
				}
				if min != datasack_min {
					t.Errorf("The datasack.Min() [%v] does NOT equal the expected min [%v]", datasack_min, min)
				}
	}
}
