# gostatic

Simple static site generator.

## Why another static site generator?

Because I can. This is a fun project, just for my personal site. Feel free to use it and extend it
if you also want some go coding fun, but if you are looking for a static site generator for
productive use, take a look at https://jekyllrb.com/.

## Usage

You can generate your static site using `go run . <path to your site data>`.

You can build the binary using `go build -o gs .`

## Tests

You can run the tests with `go test ./...` from the root directory.

The subfolder `test_data` contains the test data used by the tests contained in gostatic/*_test.go.

## Further information

For futher information, see the `doc` folder or [my blog](https://thomux.eu).