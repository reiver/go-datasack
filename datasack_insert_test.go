package datasack


import "testing"
import "math/rand"


func TestNoInsert(t *testing.T) {

	// Create a new datasack.
		datasack := New()
		if nil == datasack {
			t.Errorf("Got back 'nil' when called datasack.New().")
			return
		}


	// Make sure that if we call len(me.slice) with nothing in the slice, then we receive
	// zero (0).
		if length := len(datasack.slice); 0 != length {
			t.Errorf("Excepted len(me.slice) to return zero (0), but it did not. It returned a length of [%v]", length)
		}
}


func TestInsert(t *testing.T) {

	// Create a new datasack.
		datasack := New()
		if nil == datasack {
			t.Errorf("Got back 'nil' when called datasack.New().")
			return
		}

	// We do 1,000,000 tests.
	//
	// Which means that:
	// * we make sure len(me.slice) returns what expected and we see what we added in the slice with 1 thing  in the datasack,
	// * we make sure len(me.slice) returns what expected and we see what we added in the slice with 2 things in the datasack,
	// * we make sure len(me.slice) returns what expected and we see what we added in the slice with 3 things in the datasack,
	// * we make sure len(me.slice) returns what expected and we see what we added in the slice with 4 things in the datasack,
	// * ...
	// * we make sure len(me.slice) returns what expected and we see what we added in the slice with 1,000,000 things in the datasack,
	// * ...
	//
		for i:=0; i<1000000; i++ {

			// Calculate the expecte length that datastack.Len() should return,
			// once we add the next number to the datasack.
			// (Which we do later in this iteration of the for-loop.)
				expected_len := i + 1

			// Generate a random number to add to the datasack.
			//
			// In this case, it shouldn't matter what the number actually is.
				random_number := rand.Float64()
				if 0 == rand.Intn(5) {
					random_number *= 67108864.0
				}

			// Add the random number to the datasack.
				datasack.Insert(random_number)

			// Make sure that the len(datasack.slice) is what we expect.
				if actual_len := len(datasack.slice); expected_len != actual_len {
					t.Errorf("Expected len(me.slice) to return [%v], but instead it returned [%v]", expected_len, actual_len)
				}

			// Make sure what we just put in to the datasack is in there, where we expect it to be.
				if x := datasack.slice[i]; random_number != x {
					t.Errorf("The value at datasack.slice[%v], we expected to be [%v] but it actually was [%v]", i, random_number, x)
				}
		}

}
