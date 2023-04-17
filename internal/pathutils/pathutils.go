package pathutils

import (
	"path/filepath"
)

func PathDepth(path string) (depth int) {
	for path != filepath.Dir(path) && path != "" {
		depth++
		path = filepath.Dir(path)
	}
	return
}
