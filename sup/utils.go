package sup

import (
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func getProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", os.ErrNotExist
}

func testFilePath() string {
	for skip := 0; ; skip++ {
		_, file, _, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		if strings.Count(file, "/tests/") > 0 {
			return file
		}
	}
	panic("Unable to detect test file path")
}

func ToSnakeCase(s string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")
	snake = strings.ToLower(snake)
	return snake
}
