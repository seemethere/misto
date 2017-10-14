[![Build Status](https://travis-ci.org/seemethere/misto.svg?branch=master)](https://travis-ci.org/seemethere/misto)

```
            _     _
           (_)   | |
  _ __ ___  _ ___| |_ ___
 | '_ ` _ \| / __| __/ _ \
 | | | | | | \__ \ || (_) |
 |_| |_| |_|_|___/\__\___/
```

# misto :eyes:
> misto (*italian*), mixed (*english*)

A project about finding mixed indentation within files.

# Usage:

```shell
misto <filename>
```

**NOTE**: Misto currently only accepts filenames, use something like `find` or globs
to specify multiple files!

## Options
`--file-names-only` - Prints file names only

## Error codes:
MST1: Leading tabs with spaces after
MST2: Leading spaces with tabs after
MST3: Indentation that does not match the majority indentation style

## Exit status:
The exit status indicates how many errors were found through the entire linting process

# Installation:

With `go install` (from `master`)
```
go install -u github.com/seemethere/misto
```

# Building:

## Singular binaries

With Docker:
```
make -f docker.Makefile build
```

Without Docker:
```
make build
```

## Cross compile binaries

```
make cross
```

