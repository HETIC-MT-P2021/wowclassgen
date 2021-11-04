# Wow class generator

Simple generator to create maps in Go using [warcraft logs](https://www.warcraftlogs.com/api/v2/client) for Class and Spec Names.

It generates :

- A dict of current wow classes and their warlog IDs
- A list of Healer classes, with a point counter initialized to 0
- A list of DPS classes, with a point counter initialized to 0

## Usage

Install it by running:

```shell
go install github.com/HETIC-MT-P2021/wowclassgen@latest
```

Then run:

```shell
wowclassgen filename.go packagename
```

`filename.go` is the file (relative to the current directory) where the generator will write the code.
It does not need to exist.

`packagename` is the name of the package that will be placed at the top of the generated code file.
