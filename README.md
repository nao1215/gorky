[![MultiVersionUnitTest](https://github.com/nao1215/gorky/actions/workflows/multi_ver_test.yml/badge.svg)](https://github.com/nao1215/gorky/actions/workflows/multi_ver_test.yml)
[![PlatformTests](https://github.com/nao1215/gorky/actions/workflows/platform_test.yml/badge.svg)](https://github.com/nao1215/gorky/actions/workflows/platform_test.yml)
[![reviewdog](https://github.com/nao1215/gorky/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/gorky/actions/workflows/reviewdog.yml)
[![codecov](https://codecov.io/gh/nao1215/gorky/branch/main/graph/badge.svg?token=QlXh0Q9Cxt)](https://codecov.io/gh/nao1215/gorky)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/gorky)](https://goreportcard.com/report/github.com/nao1215/gorky)
[![Go Reference](https://pkg.go.dev/badge/github.com/nao1215/gorky.svg)](https://pkg.go.dev/github.com/nao1215/gorky)
# gorky - powerful & bitter utility library
gorky is a library of generic utility code for golang. When the autho r(it's me) was creating OSS projects, I copied the code into the project many times. I felt this process was a waste of time. For this reason, I have put together a generic code as the gorky library.

# Packages within gorky
|package name| description|
|:--|:--|
|file|check the file type and checks for the existence of files.|
|[nameconv](,/../doc/en/README_nameconv.md)|converts string naming conventions(e.g. snake case to camel case)|
|path|extracts specified string from a file path string. This package define only those features not provided by package path/filepath.|

# Support OS
- Linux (main target)
- Mac
- Windows

# Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

[GitHub Issue](https://github.com/nao1215/gup/gorky)

# Contributing
First off, thanks for taking the time to contribute! Contributions are not only related to development. For example, GitHub Star motivates me to develop!
[![Star History Chart](https://api.star-history.com/svg?repos=nao1215/gorky&type=Date)](https://star-history.com/#nao1215/gorky&Date)


# LICENSE
The nameconv project is the mixed-license.

- MIT License（casee*.go and camelcase*.go）
- Apache License Version 2.0（All codes other than the above）

The authors of the MIT license source code are [pinzolo](https://github.com/pinzolo) and [Fatih Arslan](https://github.com/fatih). The code written by each author clearly states the full MIT license and Copyright.

# Origin of name (gorky)
gorky was taken from Pokémon and a Russian writer (Maxim Gorky). gorky evolves by exchanging Pokémon with other people's Pokémon. Similarly, I hoped that the use of the gorky library in other codes would make them better.

In Russia, gorky is both the name of the writer and a word meaning "bitter".  I thought gorky meant very well, as there will be hard times on the road to a better library.

