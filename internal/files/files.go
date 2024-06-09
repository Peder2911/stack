
package files

import (
	"os"
	"path"
)

type Files struct {
	Database string
}

func DefaultFiles() (*Files, error) {

	cache,err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}
	files := Files {
		Database: path.Join(cache, "stack.bin"),
	}
	return &files,nil
}
