// Package path extracts specified string from a file path string.
// This package define only those features not provided by package path/filepath.
package path

import (
	"path/filepath"
	"strings"
)

// Ext extracts file extension from path.
// If path does not have extension, Ext return "".
func Ext(path string) string {
	base := filepath.Base(path)
	pos := strings.LastIndex(base, ".")
	if pos <= 0 {
		return ""
	}
	// hidden file
	if strings.HasPrefix(path, ".") && pos == 0 {
		return ""
	}
	return base[pos:]
}
