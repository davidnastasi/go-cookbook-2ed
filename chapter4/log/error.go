package log

import (
	"github.com/pkg/errors"
	"log"
)

func OriginalError() error {
	return errors.New("error occurred")
}

func PassThroughError() error {
	err := OriginalError()
	// no need to check error since this works with nil
	return errors.Wrap(err, "in passthrougherror")
}

func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// we log because an unexpected error occurred!
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
