package confgorm

import "os"

func MustMkdir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if err == os.ErrNotExist {
			return false
		}
	}
	return true
}
