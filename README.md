# Go Datasack

Datasack lets you store float64 numbers.

It also lets you find out the maximum and minimum float64 number in the datasack.

And it also lets you iterate through all the float64 number in the datasack.

**Note** that a datasack allows for duplicates.
I.e., you can put the same number into the datasack more than once and (unlike a set) each copy will be in there separately.


## Usage
Usage of this library is like in the following example:

```
package main


import "fmt"
import "github.com/reiver/go-datasack"


func main() {

    // Create a new datasack.
    datasack := datasack.New()


    // Insert some (float64) numbers into the datasack.
    datasack.Insert(55.555)
    datasack.Insert(-1000000.0)
    datasack.Insert(-0.10203040506070809)
    datasack.Insert(0.0)
    datasack.Insert(-0.00000400012)
    datasack.Insert(123456789)
    datasack.Insert(0.0)
    datasack.Insert(10.0)
    datasack.Insert(7.0)
    datasack.Insert(10.0)


    // Output some interesting information about the numbers in the datasack.
    max, _ := datasack.Max()
    min, _ := datasack.Min()

    fmt.Printf("The maximum value in the datasack is %v", max)
    fmt.Printf("The minimum value in the datasack is %v", min)


    // Iterate through the the numbers in the datasack.
    for number := range datasack.Data() {

        // ...
    }
}
```
