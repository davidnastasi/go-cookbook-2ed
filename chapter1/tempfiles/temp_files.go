package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WorkWithTemp will give some basic patterns for working with temporary files and directories
func WorkWithTemp() error {
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}
	defer os.RemoveAll(t)

	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}
	fmt.Println(tf.Name())
	return nil
}
