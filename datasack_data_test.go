package datasack


import "testing"
import "math/rand"
import "time"


func TestDataWithEmptyDatasack(t *testing.T) {

	// Create a new datasack.
		datasack := New()
		if nil == datasack {
			t.Errorf("Got back 'nil' when called datasack.New().")
			return
		}

	// Since there is nothing in the datasack, there should be no iterations of this loop.
	//
	// And the "number_of_iterations" variable should stay at 0 (zero).
	//
		number_of_iterations := 0

		for _ = range datasack.Data() {

			number_of_iterations++
		}


		if 0 != number_of_iterations {
			t.Errorf("When iterating through the data in an empty datasack, there should not have actually be any interactions. But had [%d]", number_of_iterations)
		}
}


func TestDataByCheckingNumberOfIterations(t *testing.T) {

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

	// We do 1,000 tests.
	//
	// Which means that:
	// * we make sure calculated length, based on the Data() method, is what is expected with 1 thing  in the datasack,
	// * we make sure calculated length, based on the Data() method, is what is expected with 2 things in the datasack,
	// * we make sure calculated length, based on the Data() method, is what is expected with 3 things in the datasack,
	// * we make sure calculated length, based on the Data() method, is what is expected with 4 things in the datasack,
	// * ...
	// * we make sure calculated length, based on the Data() method, is what is expected with 1,000 things in the datasack,
	//
		for i:=0; i<1000; i++ {

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

			// Calculate the length of the datasack based on the Data() method.
				calculated_length := 0

				for _ = range datasack.Data() {
					calculated_length++
				}

			// Make sure that the calculated length of the datasack is what we expect.
				if expected_len != calculated_length {
					t.Errorf("Expected the calculated length to be [%v], but instead it was [%v]", expected_len, calculated_length)
				}		
		}
}


func TestDataByCheckingWhatIsInTHere(t *testing.T) {

	// Define tests.
		tests := []struct{
			numbers []float64
		}{
			{
				numbers: []float64{1.0},
			},
			{
				numbers: []float64{1.0, 2.0},
			},
			{
				numbers: []float64{1.0, 2.0, 3.0},
			},
			{
				numbers: []float64{1.0, 2.0, 3.0, 4.0},
			},

			{
				numbers: []float64{3.14, 0.001592, 0.0000006535, 0.00000000008979, 0.00000000000000323, 0.00000000000000000846264338, 0.00000000000000000000000000327950},
			},
		}

	//
		for test_number, test := range tests {


			// Create a new datasack.
				datasack := New()
				if nil == datasack {
					t.Errorf("Got back 'nil' when called datasack.New().")
					return
				}

			// Put numbers into the datasack.
				for _,number := range test.numbers {
					datasack.Insert(number)
				}

			// Confirm that we get back the numbers we expect and in the order we expect.
				index := 0
				for returned_number := range datasack.Data() {
					if expected := test.numbers[index]; expected != returned_number {
						t.Errorf("In test #%d, for number with index %d, expected %v but instead got %v", test_number, index, expected, returned_number)
					}

					index++
				}
		}
}
