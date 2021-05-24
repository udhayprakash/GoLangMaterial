package main

/*
By using the pkg/errors package instead of the stdlib copy,
we’re able to Wrap() and Unwrap() errors as they bubble up.

This gives the great advantage of adding context along the way.
If you see a generic and specific error in a large code base they
can be difficult to track down without this added context.

*/
import (
	"log"
	// Note: we're using an 'enhanced' errors package
	"github.com/pkg/errors"
)

func main() {
	err := WholeJob()
	if err != nil {
		log.Fatal(err)
	}
}

func WholeJob() error {
	err := SpecificThing()
	if err != nil {
		// Wrap() accepts the original error, then a custom message
		return errors.Wrap(err, "The whole job failed")
	}
	return nil
}

func SpecificThing() error {
	return errors.New("The specific thing failed")
}
