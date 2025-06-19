package is

import (
	"fmt"
	"os"
)

func DirExists(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("directory %s doesn't exist\n: %v", dir, err)
		} else {
			return err
		}
	}

	return nil
}
