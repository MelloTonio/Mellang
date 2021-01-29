package FileMatch

import (
	"os"
	"path/filepath"
	"strings"
)

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func FindFilename() string {
	myDir, _ := os.Getwd()
	myString, _ := WalkMatch(myDir, "*.mel")

	newStr := strings.Split(myString[0], "/")

	return newStr[len(newStr)-1]

}
