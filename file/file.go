// Package file defines functions to determine the file contents.
// For example, IsFile() returns whether the path is a file or not with a bool value.
package file

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	// Unknown : Don't use this bits
	Unknown os.FileMode = 1 << (9 - iota)
	// Readable : readable bits
	Readable
	// Writable : writable bits
	Writable
	// Executable : executable bits
	Executable
)

const (
	// FileModeCreatingDir is used for creating directory
	FileModeCreatingDir fs.FileMode = 0750
	// FileModeCreatingFile is used for creating directory
	FileModeCreatingFile fs.FileMode = 0600
)

// IsFile reports whether the path exists and is a file.
func IsFile(path string) bool {
	stat, err := os.Stat(path)
	return (err == nil) && (!stat.IsDir())
}

// Exists reports whether the path exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return (err == nil)
}

// IsDir reports whether the path exists and is a directory.
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	return (err == nil) && (stat.IsDir())
}

// IsSymlink reports whether the path exists and is a symbolic link.
func IsSymlink(path string) bool {
	stat, err := os.Lstat(path)
	if err != nil {
		return false
	}
	if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true
	}
	return false
}

// IsZero reports whether the path exists and is zero size.
func IsZero(path string) bool {
	stat, err := os.Stat(path)
	return (err == nil) && (stat.Size() == 0)
}

// IsReadable reports whether the path exists and is readable.
func IsReadable(path string) bool {
	stat, err := os.Stat(path)
	return (err == nil) && ((stat.Mode() & Readable) != 0)
}

// IsWritable reports whether the path exists and is writable.
func IsWritable(path string) bool {
	stat, err := os.Stat(path)
	return (err == nil) && ((stat.Mode() & Writable) != 0)
}

// IsExecutable reports whether the path exists and is executable.
func IsExecutable(path string) bool {
	if runtime.GOOS == "windows" {
		return strings.HasSuffix(path, ".exe")
	}
	stat, err := os.Stat(path)
	return (err == nil) && ((stat.Mode() & Executable) != 0)
}

// IsHiddenFile reports whether the path exists and is included hidden file.
func IsHiddenFile(filePath string) bool {
	_, file := filepath.Split(filePath)
	if IsFile(filePath) && strings.HasPrefix(file, ".") {
		return true
	}
	return false
}

// Copy copy file to destination path
func Copy(src string, dest string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}
	return d.Close()
}
