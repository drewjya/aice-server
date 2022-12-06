package helper

import (
	"path/filepath"
	"regexp"
)

func ImageExtension(filename string) string {

	var extension = filepath.Ext(filename)

	return extension

}

func ValidateExtension(filename string) bool {
	re := regexp.MustCompile(`(?m)\.(jpe?g|png)$`)

	match := re.Match([]byte(filename))
	return match
}
