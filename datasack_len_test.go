package datasack


import "testing"
import "math/rand"
import "time"


func TestLenWithEmptyDatasack(t *testing.T) {

	// Create a new datasack.
		datasack := New()
		if nil == datasack {
			t.Errorf("Got back 'nil' when called datasack.New().")
			return
		}


	// Make sure that if we call Len() on a datasack with nothing in it, then we receive
	// zero (0).
		if length := datasack.Len(); 0 != length {
			t.Errorf("Excepted Len() to return zero (0), but it did not. It returned a length of [%v]", length)
		}
}


func TestLen(t *testing.T) {

	// Seed the random number generator.
	//
	// It is import to seed the random number generator, so that
	// we get different random number each time we run this test.
	//
	// If we don't seed it, then we get the same "random" numbers
	// over and over again!
	//
		rand.Seed( time.Now().UTC().UnixNano() )

	// Create a new datasack.
		datasack := New()
		if nil == datasack {
			t.Errorf("Got back 'nil' when called datasack.New().")
			return
		}

	// We do 1,000,000 tests.
	//
	// Which means that:
	// * we make sure Len() returns what expected with 1 things in the datasack,
	// * we make sure Len() returns what expected with 2 things in the datasack,
	// * we make sure Len() returns what expected with 3 things in the datasack,
	// * we make sure Len() returns what expected with 4 things in the datasack,
	// * ...
	// * we make sure Len() returns what expected with 1,000,000 things in the datasack.
	//
		for i:=0; i<1000000; i++ {

			// Calculate the expecte length that datastack.Len() should return.
				expected_len := i + 1

			// Generate a random number to add to the datasack.
			//
			// In this case, it doesn't matter what it is.
				random_number := rand.Float64()
				if 0 == rand.Intn(5) {
					random_number *= 67108864.0
				}

			// Add the random number to the datasack.
				datasack.Insert(random_number)

			// Make sure that the Len() of the datasack is what we expect.
				if actual_len := datasack.Len(); expected_len != actual_len {
					t.Errorf("Expected Len() to return [%v], but instead it returned [%v]", expected_len, actual_len)
				}		
		}

}
