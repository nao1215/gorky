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
	"testing"
)

type boolExpectedCase struct {
	target   string
	expected bool
}

type stringExpectedCase struct {
	target   string
	expected string
}

func TestIsSnakeCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"foo_bar", true},
		{"foo1_bar2", true},
		{"foo_bar_1", true},
		{"foo_bar_1", true},
		{"111_foo_bar", true},
		{"foobar", true},
		{"foobar1", true},
		{"foo1bar", true},
		{"111foobar", true},

		{"", false},
		{"FOO_BAR", false},
		{"fooBar", false},
		{"FooBar", false},
		{"FOOBAR", false},
		{"foo_@bar", false},
		{"foo-bar", false},
		{"テスト", false},
		{"テスト_テスト", false},
	}

	for _, tc := range testCases {
		if actual := IsSnakeCase(tc.target); actual != tc.expected {
			t.Errorf("IsSnakeCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "foo_bar"},
		{"foo-bar", "foo_bar"},
		{"foo-bar_baz", "foo_bar_baz"},
		{"foo--bar__baz", "foo_bar_baz"},
		{"fooBar", "foo_bar"},
		{"FooBar", "foo_bar"},
		{"foo bar", "foo_bar"},
		{"   foo   bar   ", "foo_bar"},
		{"fooBar111", "foo_bar_111"},
		{"111FooBar", "111_foo_bar"},
		{"foo-111-Bar", "foo_111_bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToSnakeCase(tc.target); actual != tc.expected {
			t.Errorf("ToSnakeCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
}

func TestIsChainCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"foo-bar", true},
		{"foo1-bar2", true},
		{"foo-bar-1", true},
		{"foo-bar-1", true},
		{"111-foo-bar", true},
		{"foobar", true},
		{"foobar1", true},
		{"foo1bar", true},
		{"1foobar", true},

		{"", false},
		{"FOO-BAR", false},
		{"fooBar", false},
		{"FooBar", false},
		{"FOOBAR", false},
		{"foo-@bar", false},
		{"foo_bar", false},
		{"テスト", false},
		{"テスト-テスト", false},
	}

	for _, tc := range testCases {
		if actual := IsChainCase(tc.target); actual != tc.expected {
			t.Errorf("IsChainCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestToChainCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "foo-bar"},
		{"foo-bar", "foo-bar"},
		{"foo-bar_baz", "foo-bar-baz"},
		{"foo--bar__baz", "foo-bar-baz"},
		{"fooBar", "foo-bar"},
		{"FooBar", "foo-bar"},
		{"foo bar", "foo-bar"},
		{"   foo   bar   ", "foo-bar"},
		{"fooBar111", "foo-bar-111"},
		{"111FooBar", "111-foo-bar"},
		{"foo-111-Bar", "foo-111-bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToChainCase(tc.target); actual != tc.expected {
			t.Errorf("ToChainCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
}

func TestIsCamelCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"fooBar", true},
		{"fooBar1", true},
		{"foo1Bar", true},

		{"", false},
		{"1FooBar", false},
		{"1fooBar", false},
		{"foo@Bar", false},
		{"foo_bar", false},
		{"FOO_BAR", false},
		{"FooBar", false},
		{"Foo@Bar", false},
	}

	for _, tc := range testCases {
		if actual := IsCamelCase(tc.target); actual != tc.expected {
			t.Errorf("IsCamelCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestToCamelCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "fooBar"},
		{"foo-bar", "fooBar"},
		{"foo-bar_baz", "fooBarBaz"},
		{"foo--bar__baz", "fooBarBaz"},
		{"fooBar", "fooBar"},
		{"FooBar", "fooBar"},
		{"foo bar", "fooBar"},
		{"   foo   bar   ", "fooBar"},
		{"fooBar111", "fooBar111"},
		{"111FooBar", "111FooBar"},
		{"foo-111-Bar", "foo111Bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToCamelCase(tc.target); actual != tc.expected {
			t.Errorf("ToCamelCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
}

func TestIsPascalCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"FooBar", true},
		{"FooBar1", true},
		{"Foo1Bar", true},

		{"", false},
		{"1FooBar", false},
		{"1fooBar", false},
		{"Foo@Bar", false},
		{"foo_bar", false},
		{"FOO_BAR", false},
		{"fooBar", false},
		{"foo@Bar", false},
	}

	for _, tc := range testCases {
		if actual := IsPascalCase(tc.target); actual != tc.expected {
			t.Errorf("IsPascalCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestToPascalCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "FooBar"},
		{"foo-bar", "FooBar"},
		{"foo-bar_baz", "FooBarBaz"},
		{"foo--bar__baz", "FooBarBaz"},
		{"fooBar", "FooBar"},
		{"FooBar", "FooBar"},
		{"foo bar", "FooBar"},
		{"   foo   bar   ", "FooBar"},
		{"fooBar111", "FooBar111"},
		{"111FooBar", "111FooBar"},
		{"foo-111-Bar", "Foo111Bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToPascalCase(tc.target); actual != tc.expected {
			t.Errorf("ToPascalCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
}

func TestToFlatCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "foobar"},
		{"foo-bar", "foobar"},
		{"foo-bar_baz", "foobarbaz"},
		{"foo--bar__baz", "foobarbaz"},
		{"fooBar", "foobar"},
		{"FooBar", "foobar"},
		{"foo bar", "foobar"},
		{"   foo   bar   ", "foobar"},
		{"fooBar111", "foobar111"},
		{"111FooBar", "111foobar"},
		{"foo-111-Bar", "foo111bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToFlatCase(tc.target); actual != tc.expected {
			t.Errorf("ToFlatCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
}

func TestIsFlatCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"foobar", true},
		{"foo1bar", true},

		{"", false},
		{"1foobar", false},
		{"foo@bar", false},
		{"foo_bar", false},
		{"FOO_BAR", false},
		{"FooBar", false},
		{"fooBar", false},
		{"foo_bar", false},
	}

	for _, tc := range testCases {
		if actual := IsFlatCase(tc.target); actual != tc.expected {
			t.Errorf("IsFlatCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestIsUpperCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"FOO_BAR", true},
		{"FOO", true},
		{"FOO_1_BAR", true},

		{"", false},
		{"1FooBar", false},
		{"1fooBar", false},
		{"FOO@BAR", false},
		{"Foo_Bar", false},
		{"fooBar", false},
		{"foo@Bar", false},
	}

	for _, tc := range testCases {
		if actual := IsUpperCase(tc.target); actual != tc.expected {
			t.Errorf("IsUpperCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestToUpperCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "FOO_BAR"},
		{"foo-bar", "FOO_BAR"},
		{"foo-bar_baz", "FOO_BAR_BAZ"},
		{"foo--bar__baz", "FOO_BAR_BAZ"},
		{"fooBar", "FOO_BAR"},
		{"FooBar", "FOO_BAR"},
		{"foo bar", "FOO_BAR"},
		{"   foo   bar   ", "FOO_BAR"},
		{"fooBar111", "FOO_BAR_111"},
		{"111FooBar", "111_FOO_BAR"},
		{"foo-111-Bar", "FOO_111_BAR"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToUpperCase(tc.target); actual != tc.expected {
			t.Errorf("ToUpperCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
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

func TestToKebabCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "foo-bar"},
		{"foo-bar", "foo-bar"},
		{"foo-bar_baz", "foo-bar-baz"},
		{"foo--bar__baz", "foo-bar-baz"},
		{"fooBar", "foo-bar"},
		{"FooBar", "foo-bar"},
		{"foo bar", "foo-bar"},
		{"   foo   bar   ", "foo-bar"},
		{"fooBar111", "foo-bar-111"},
		{"111FooBar", "111-foo-bar"},
		{"foo-111-Bar", "foo-111-bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToKebabCase(tc.target); actual != tc.expected {
			t.Errorf("ToKebabCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
}

func TestIsKebabCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"foo-bar", true},
		{"foo1-bar2", true},
		{"foo-bar-1", true},
		{"foo-bar-1", true},
		{"111-foo-bar", true},
		{"foobar", true},
		{"foobar1", true},
		{"foo1bar", true},
		{"1foobar", true},

		{"", false},
		{"FOO-BAR", false},
		{"fooBar", false},
		{"FooBar", false},
		{"FOOBAR", false},
		{"foo-@bar", false},
		{"foo_bar", false},
		{"テスト", false},
		{"テスト-テスト", false},
	}

	for _, tc := range testCases {
		if actual := IsKebabCase(tc.target); actual != tc.expected {
			t.Errorf("IsKebabCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestIsFirstRuneUpper(t *testing.T) {
	testCases := ""
	if isFirstRuneUpper(testCases) {
		t.Error("isFirstRuneUpper() returns true, but expected false")
	}
}

func TestIsFirstRuneLower(t *testing.T) {
	testCases := ""
	if isFirstRuneLower(testCases) {
		t.Error("isFirstRuneLower() returns true, but expected false")
	}
}

func TestGetRune(t *testing.T) {
	testCases := ""
	if getRune(testCases) != 0 {
		t.Error("getRune() returns not 0, but expected 0")
	}
}
