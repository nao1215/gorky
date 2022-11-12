package nameconv

//
// This file is forked from https://github.com/fatih/camelcase
//
// The MIT License (MIT)
//
// Copyright (c) 2015 Fatih Arslan
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
//

import "fmt"

func ExampleSplit() {

	for _, c := range []string{
		"",
		"lowercase",
		"Class",
		"MyClass",
		"MyC",
		"HTML",
		"PDFLoader",
		"AString",
		"SimpleXMLParser",
		"vimRPCPlugin",
		"GL11Version",
		"99Bottles",
		"May5",
		"BFG9000",
		"BöseÜberraschung",
		"Two  spaces",
		"BadUTF8\xe2\xe2\xa1",
	} {
		fmt.Printf("%#v => %#v\n", c, split(c))
	}

	// Output:
	// "" => []string{}
	// "lowercase" => []string{"lowercase"}
	// "Class" => []string{"Class"}
	// "MyClass" => []string{"My", "Class"}
	// "MyC" => []string{"My", "C"}
	// "HTML" => []string{"HTML"}
	// "PDFLoader" => []string{"PDF", "Loader"}
	// "AString" => []string{"A", "String"}
	// "SimpleXMLParser" => []string{"Simple", "XML", "Parser"}
	// "vimRPCPlugin" => []string{"vim", "RPC", "Plugin"}
	// "GL11Version" => []string{"GL", "11", "Version"}
	// "99Bottles" => []string{"99", "Bottles"}
	// "May5" => []string{"May", "5"}
	// "BFG9000" => []string{"BFG", "9000"}
	// "BöseÜberraschung" => []string{"Böse", "Überraschung"}
	// "Two  spaces" => []string{"Two", "  ", "spaces"}
	// "BadUTF8\xe2\xe2\xa1" => []string{"BadUTF8\xe2\xe2\xa1"}
}
