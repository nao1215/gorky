# golden - Golden test utility for Go
`golden` is a golden file test utility for Go projects. It's typically used for testing responses with larger data bodies. 
  
The concept is straight forward. Valid response data is stored in a "golden file". The actual response data will be byte compared with the golden file and the test will fail if there is a difference. Updating the golden file can be done by running `go test -update ./...`.

See the [GoDoc](https://godoc.org/github.com/nao1215/golden) for API reference and configuration options.
This project is forked from [github.com/sebdah/goldie](https://github.com/sebdah/goldie). Original project (== goldie) is not support windows. Also, there is deprecated function and not enough unit tests. So, I forked it.

# Support OS
- Windows
- Mac
- Linux
  
# Example usage
## Basic assertions
The below example fetches data from a REST API. The last line in the test is the
actual usage of `golden`. It takes the HTTP response body and asserts that it's
what is present in the golden test file.

```
func TestExample(t *testing.T) {
    recorder := httptest.NewRecorder()

    req, err := http.NewRequest("GET", "/example", nil)
    assert.Nil(t, err)

    handler := http.HandlerFunc(ExampleHandler)
    handler.ServeHTTP()

    g := golden.New(t)
    // Compare testdata/example.golden and recorder.Body.Bytes()
    g.Assert(t, "example", recorder.Body.Bytes())
}
```

## Assertions using templates
If some values in the golden file can change depending on the test, you can use
golang template in the golden file and pass the data to `AssertWithTemplate`.

### example.golden
```
This is a {{ .Type }} file.
```

### Test
```
func TestTemplateExample(t *testing.T) {
    recorder := httptest.NewRecorder()

    req, err := http.NewRequest("POST", "/example/Golden", nil)
    assert.Nil(t, err)

    handler := http.HandlerFunc(ExampleHandler)
    handler.ServeHTTP()

    data := struct {
        Type	string
    }{
        Type:	"Golden",
    }

    g := golden.New(t)
    g.AssertWithTemplate(t, "example", data, recorder.Body.Bytes())
}
```

Then run your test with the `-update` flag the first time to store the result.

`go test -update ./...`

For any consecutive runs where you actually want to compare the data, simply
drop the `-update` flag.

`go test ./...`

## Validating JSON and XML output

If you are asserting JSON and XML data, you can use the handy `AssertJson` and
`AssertXml` functions that will nicely indent the golden validation files for
better readability.

# Flags

## Clean output directory

Using `-update` along with `-clean` flag will clear the fixture directory before updating golden files.

`go test -update -clean ./...`


# Options

`golden` supports a number of configuration options that will alter the behavior
of the library.  These options should be passed into the `golden.New()` method.

```
func TestNewExample(t *testing.T) {
    g := golden.New(
        t,
        golden.WithFixtureDir("test-fixtures"),
        golden.WithNameSuffix(".golden.json"),
        golden.WithDiffEngine(golden.ColoredDiff),
        golden.WithTestNameForDir(true),
    )

    g.Assert(t, "example", []byte("my example data"))
}
```

## Available options

| Option                     | Comment                                                  | Default
|----------------------------|----------------------------------------------------------|-------------
| `WithFixtureDir`           | Set fixture dir name                                     | `testdata`
| `WithNameSuffix`           | Suffix for fixture files.                                | `.golden`
| `WithDirPerms`             | Directory permissions for fixtures                       | `0755`
| `WithFilePerms`            | File permissions for fixtures                            | `0644`
| `WithDiffEngine`           | Diff engine to use for diff output                       | `ClassicDiff`
| `WithDiffFn`               | Custom diff logic to be used                             | None
| `WithIgnoreTemplateErrors` | Ignore errors from templates                             | `false`
| `WithTestNameForDir`       | Create a folder with the tests name for the fixtures     | `false`
| `WithSubTestNameForDir`    | Create a folder with the sub tests name for the fixtures | `false`

## Diff output

golden has three output modes; classic diff (default), colored diffs and simple
mode.

You can select your preferred output using the `WithDiffEngine` option:

```
g.New(
    t,
    golden.WithDiffEngine(golden.ColoredDiff), // Simple, ColoredDiff, ClassicDiff
)
```

## New default fixture directory
There is a new default directory for fixtures, `testdata`. This directory is a better default as it is more widely used in the Go community (including the standard library). 

## New way to initialize golden
With the introduction of the functional options we also introduced `golden.New`, which is initializing golden. `Assert*` and other methods are now accessed like:

```
g := golden.New(t)
g.Assert(t, ...)
```

# License
The golden project is licensed under the terms of [MIT LICENSE](./LICENSE).
Original author is [Sebastian Dahlgren](https://github.com/sebdah/).
