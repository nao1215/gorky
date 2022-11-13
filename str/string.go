// Package str implement string manipulations not provided by the golang standard package (strings package)
package str

import "strings"

// TrimGaps Remove white space at the beginning/end of a
// string and single out multiple white spaces between characters.
// Whitespace includes tabs and line feed.
// " Hello,    World  ! "         --> "Hello, World !"
// "Hello,\tWorld ! "             --> "Hello, World !"
// " \t\n\t Hello, \n\t World \n ! \n\t " --> "Hello, World !"
func TrimGaps(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Contains returns whether the specified string is contained in the slice.
func Contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
