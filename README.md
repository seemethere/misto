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

# Building:

With Docker:
```
make -f docker.Makefile build
```

Without Docker:
```
make build
```

# TODO:

* Have a parser that determines the most common indentation in a file and
finds lines that do not conform to that indentation style
* Have command line options to:
  * Only print filenames if wanted
* Add some tests
* Maybe break the parsing rules into a library for other utilities to use?
* Work with goroutines
