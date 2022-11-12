// Package nameconv is library that converts string naming conventions.
// Convertible formats are Snake Case, Chain Case, Kebab Case, Camel Case,
// Pascal Case, Flat Case, and Upper Case.
package nameconv

//
// This file is forked from https://github.com/pinzolo/casee
//
// The MIT License (MIT)
//
// Copyright (c) 2016 pinzolo
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this
// software and associated documentation files (the "Software"), to deal in the Software
// without restriction, including without limitation the rights to use, copy, modify, merge,
// publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons
// to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or
// substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
// PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE
// FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

import (
	"strings"
	"unicode"
)

// ToSnakeCase convert string to snake_case style.
// If argument is empty, return itself.
func ToSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	return strings.Join(fields, "_")
}

// IsSnakeCase check whether string is snake_case style or not.
// true: string is snake case, false: other
func IsSnakeCase(s string) bool {
	if strings.Contains(s, "_") {
		fields := strings.Split(s, "_")
		for _, field := range fields {
			if !isMadeByLowerAndDigit(field) {
				return false
			}
		}
		return true
	}
	return isMadeByLowerAndDigit(s)
}

// ToChainCase convert string to chain-case style.
// If argument is empty, return itself.
func ToChainCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	return strings.Join(fields, "-")
}

// IsChainCase check whether string is chain-case style.
// true: string is chain case, false: other
func IsChainCase(s string) bool {
	if strings.Contains(s, "-") {
		fields := strings.Split(s, "-")
		for _, field := range fields {
			if !isMadeByLowerAndDigit(field) {
				return false
			}
		}
		return true
	}
	return isMadeByLowerAndDigit(s)
}

// ToCamelCase convert string to camelCase style.
// If argument is empty, return itself
func ToCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	for i, f := range fields {
		if i != 0 {
			fields[i] = toUpperFirstRune(f)
		}
	}
	return strings.Join(fields, "")
}

// IsCamelCase check whether string is camelCase style or not.
// true: string is camel case, false: other.
// If first character is digit, always returns false
func IsCamelCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	}
	return isMadeByAlphanumeric(s) && isFirstRuneLower(s)
}

// ToPascalCase convert string to PascalCase style.
// If argument is empty, return itself
func ToPascalCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	for i, f := range fields {
		fields[i] = toUpperFirstRune(f)
	}
	return strings.Join(fields, "")
}

// IsPascalCase check whether string is PascalCase style or not.
// true: string is pascal case, false: other.
// If first character is digit, always returns false
func IsPascalCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	}
	return isMadeByAlphanumeric(s) && isFirstRuneUpper(s)
}

// ToFlatCase convert string to flatcase style.
// If argument is empty, return itself
func ToFlatCase(s string) string {
	if len(s) == 0 {
		return s
	}

	fields := splitToLowerFields(s)
	return strings.Join(fields, "")
}

// IsFlatCase check whether string is flatcase style or not.
// true: string is flat case, false: other.
// If first character is digit, always returns false
func IsFlatCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	}
	return isMadeByLowerAndDigit(s)
}

// ToUpperCase convert string to UPPER_CASE style.
// If argument is empty, return itself.
func ToUpperCase(s string) string {
	return strings.ToUpper(ToSnakeCase(s))
}

// IsUpperCase check whether string is UPPER_CASE style or not.
// true: string is flat case, false: other.
func IsUpperCase(s string) bool {
	if strings.Contains(s, "_") {
		fields := strings.Split(s, "_")
		for _, field := range fields {
			if !isMadeByUpperAndDigit(field) {
				return false
			}
		}
		return true
	}
	return isMadeByUpperAndDigit(s)
}

func isMadeByLowerAndDigit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isMadeByUpperAndDigit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isMadeByAlphanumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) && !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isFirstRuneUpper(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsUpper(getRune(s))
}

func isFirstRuneLower(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsLower(getRune(s))
}

func isFirstRuneDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	return unicode.IsDigit(getRune(s))
}

func getRune(s string) rune {
	if len(s) == 0 {
		return 0
	}

	rs := []rune(s)
	return rs[0]
}

func splitToLowerFields(s string) []string {
	defaultCap := len([]rune(s)) / 3
	fields := make([]string, 0, defaultCap)

	for _, sf := range strings.Fields(s) {
		for _, su := range strings.Split(sf, "_") {
			for _, sh := range strings.Split(su, "-") {
				for _, sc := range split(sh) {
					fields = append(fields, strings.ToLower(sc))
				}
			}
		}
	}
	return fields
}

func toUpperFirstRune(s string) string {
	rs := []rune(s)
	return strings.ToUpper(string(rs[0])) + string(rs[1:])
}

//------The following source code does not exist in the original version----------------------
//
// The above code is the copyright below.
//
// Copyright 2022 Naohiro CHIKAMATSU
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ToKebabCase convert string to kebab-case style. It's same as ToChainCase().
// If argument is empty, return itself.
func ToKebabCase(s string) string {
	return ToChainCase(s)
}

// IsKebabCase check whether string is kebab-case style. It's same as IsChainCase().
// true: string is chain case, false: other
func IsKebabCase(s string) bool {
	return IsChainCase(s)
}
