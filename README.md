# Parallel

Parallel is a small tool that execute shell commands in parallel using workers.

## How it works

Using an input file, Parallel reads the list of shell commands to execute (one per line) and excute then using multiples workers.

*Example :
If we have 10 commands to execute using 2 workers, Parallel will make the first worker execute the first command and the second worker execute the second command. Then, whenever the first or the second finished, Parrallel will make the first avaible worker to execute the next command.*

## Usage

```sh
Usage of parallel:
  -f string
        The file containing the commands to play
  -w int
        The number of parallel processus to use (default 2)
```


## Installation

```sh
$ go get -u github.com/seblegall/parallel/cmd/parallel
```