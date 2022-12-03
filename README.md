[![MultiVersionUnitTest](https://github.com/nao1215/gorky/actions/workflows/multi_ver_test.yml/badge.svg)](https://github.com/nao1215/gorky/actions/workflows/multi_ver_test.yml)
[![LinuxUnitTest](https://github.com/nao1215/gorky/actions/workflows/linux_test.yml/badge.svg)](https://github.com/nao1215/gorky/actions/workflows/linux_test.yml)
[![MacUnitTest](https://github.com/nao1215/gorky/actions/workflows/mac_test.yml/badge.svg)](https://github.com/nao1215/gorky/actions/workflows/mac_test.yml)
[![WindowsUnitTest](https://github.com/nao1215/gorky/actions/workflows/windows.yml/badge.svg)](https://github.com/nao1215/gorky/actions/workflows/windows.yml)
[![reviewdog](https://github.com/nao1215/gorky/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/gorky/actions/workflows/reviewdog.yml)
[![codecov](https://codecov.io/gh/nao1215/gorky/branch/main/graph/badge.svg?token=QlXh0Q9Cxt)](https://codecov.io/gh/nao1215/gorky)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/gorky)](https://goreportcard.com/report/github.com/nao1215/gorky)
[![Go Reference](https://pkg.go.dev/badge/github.com/nao1215/gorky.svg)](https://pkg.go.dev/github.com/nao1215/gorky)
# gorky - powerful & bitter utility library
gorky is a library of generic utility code for golang. When the author (it's me) was creating OSS projects, I copied the code into the project many times. I felt this process was a waste of time. For this reason, I have put together a generic code as the gorky library.

# Packages within gorky
|package name| description|
|:--|:--|
|errfmt| format the error message|
|file|check the file type and checks for the existence of files.|
|[nameconv](doc/en/README_nameconv.md)|converts string naming conventions(e.g. snake case to camel case)|
|path|extracts specified string from a file path string. This package define only those features not provided by package path/filepath.|
|str|implement string manipulations not provided by the golang standard package (strings package)|
|[golden](doc/en/README_golden.md)|golden file test utility|

# Support OS & go version
- Linux (main target)
- Mac
- Windows
- go version 1.16 to 1.19

# Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

[GitHub Issue](https://github.com/nao1215/gorky/issues)

# Contributing
First off, thanks for taking the time to contribute! Contributions are not only related to development. For example, GitHub Star motivates me to develop!
[![Star History Chart](https://api.star-history.com/svg?repos=nao1215/gorky&type=Date)](https://star-history.com/#nao1215/gorky&Date)


# LICENSE
### nameconv
The nameconv project is the mixed-license.

- MIT License（casee*.go and camelcase*.go）
- Apache License Version 2.0（All codes other than the above）

The authors of the MIT license source code are [pinzolo](https://github.com/pinzolo) and [Fatih Arslan](https://github.com/fatih). The code written by each author clearly states the full MIT license and Copyright.

### golden 
The golden project is licensed under the terms of [MIT LICENSE](./LICENSE).
Original author is [Sebastian Dahlgren](https://github.com/sebdah/).

### others
[MIT LICENSE](./LICENSE)

# Origin of name (gorky)
gorky was taken from Pokémon and a Russian writer (Maxim Gorky). gorky evolves by exchanging Pokémon with other people's Pokémon. Similarly, I hoped that the use of the gorky library in other codes would make them better.

In Russia, gorky is both the name of the writer and a word meaning "bitter".  I thought gorky meant very well, as there will be hard times on the road to a better library.
