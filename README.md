# Exercism's Go Representer
This is Exercism's representer for the Go track. See the
[docs](https://exercism.org/docs/building/tooling/representers) for more
information on representers.

## Build Executable
This will create an executable called `represent`.

```
go build -tags build -o represent .
```

## Executing the Representer
The representer takes three arguments:

* The slug of the exercise (e.g. `two-fer`).
* A path to a directory containing the submitted file(s) (with a trailing slash).
* A path to an output directory (with a trailing slash). This directory is writable.

The representer will write a `representation.txt` and `mapping.json`
file to the output directory.

With the binary built above:
```
represent two-fer ./representer/testdata/two-fer/1/ ./
```

## Current Normalizations

The basic normalizations recommended in the [docs](https://exercism.org/docs/building/tooling/representers/normalization) are implemented.

* Use placeholders for variable names
* Remove comments
* Apply standard formatting
* Consolidate multiple files
* Sort top level declarations (imports, types, constants, variables, functions)

## Batch Analysis

See [Batch Analysis README](https://github.com/exercism/go-representer/blob/main/script/README.md) for information on how to run the representer on a bigger data set.