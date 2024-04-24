# Advent of code 2021

![Advent of Code](./AOC_LOGO.jpg)

## About
I am using `Advent of Code` to learn [Go](https://go.dev/). Let's... go!
So far it has been beneficial for learning:
- File i/o in go
- Handling multiple files and packages
- Using command line arguments to run certain days
- And more to come

[Advent of Code](https://adventofcode.com/)

## Setup
Input files are expected to be located in the `inputs` folder under the correct folder for the day, and to have the extension `.txt`. Example `inputs/01/test.txt`. 

Feel free to add more `.txt` input files in the appropriate folder if you want to try different inputs. You will simply need to adjust the run command accordingly. See next section.

## Running the Go solutions

- Make sure that Go is installed (instructions [here](https://go.dev/doc/install)) and that it is possible to call the `go` command in a terminal
- Then, to get the solution for `day01` and for the input file  `inputs/01/input.txt`, run the following:

```
cd go
go run main.go input 01
```
Alternatively, to solve today's problem with the corresponding input `input.txt`:
```
go run main.go input
```
Note that the last parameter in the command above corresponds to the file name without the extension. So for example if you want to run the solution for the file `test.txt` you can run:
```
go run main.go test
```
  
