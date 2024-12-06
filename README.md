# advent-of-code-2024-4
advent-of-code-2024 day 4
<br>
By Francis Nickels - 12/5/24

This is my implementation of the [Advent of Code](https://adventofcode.com/) 2024 solution written in Go.  

https://adventofcode.com/2024


The solution requirements used to develop this implementation are defined in the attached [Problem Statement](./PROBLEM-STATEMENT.md).

The [input file](./input-file.txt) is saved here.

## Usage

There are executables for Linux, Windows, and MacOS (AMD64) that can be run to validate the solution.  Each expects the input via stdin.  To feed a preexisting file to the application simply redirect the input from a file. 

If you have successfully compiled the application using `make build` the `gol` executable should be native to your O/S.  Otherwise feel free to use one of the [3 precompiled binaries](https://gh.riotgames.com/fnickels/gol/releases) on the appropriate system.

```bash
./advent-of-code
```

## Results

The AoC stats from immediately after the part solution being posted are captured [here](./results.txt)

## Source Code

The source code for this application is located in the [src](./src) directory.

## How to Build executable

The following are a subset of the make targets created for building and testing the application.  

For a full list run `make help`.

### Run everything

```bash
make run
# or 
make goversion goinit golint gotest build exec
```

### Build & execute test data files

```bash
make build exec
```

### Build binary

```bash
make build
```

### Golang Unit & Integration Tests

```bash
make gotest
```

### Linter

```bash
make golint
```

### (Re-)Initialize Project

```bash
make goinit
```

### Help

```bash
make 
# or 
make help
```

### Prerequisites

* Go v1.22
* golangci-lint v1.55.2
* awk


## Debugging

There are two [debugging configs](./.vscode/launch.json) configured for use in VSCode. 

* **Debug Advent-of-Code App with test data**
  * *for general debugging of the application*
* **Debug Tests of Advent-of-Code App**
  * *for debugging Unit & Integration tests*
