package pathutil

import (
	"fmt"
	"os"
)

// Exists checks if a file or directory exists.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// DirExists checks if a path exists and is a directory.
func DirExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil && fi.IsDir() {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsEmpty checks if a given path is empty.
func IsEmpty(path string) (bool, error) {
	if b, _ := Exists(path); !b {
		return false, fmt.Errorf("%q path does not exist", path)
	}
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if fi.IsDir() {
		f, err := os.Open(path)
		defer f.Close()
		if err != nil {
			return false, err
		}
		list, err := f.Readdir(-1)
		// f.Close() - see bug fix above
		return len(list) == 0, nil
	}
	return fi.Size() == 0, nil
}

// IsDir checks if a given path is a directory.
func isDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}
