# Seerlink
Seerlink Core is the API backend that Seerlink client contracts on Ethereum 
make requests to. The backend utulizes Solidity contract ABIs to generate types 
for interacting with Ethereum contracts.

## Features

* Headless API implementation
* CLI tool providing conveniance commands for node configuration, administration,
  and CRUD object operations (e.g. Jobs, Runs, and even the VRF)

## Installation

See the [root README](../README.md#install)
for instructions on how to build the full Seerlink node.


## Common Commands

**Install:**

By default `go install` will install this directory under the name `core`.
You can instead, build it, and place it in your path as `seerlink`:

```
go build -o ./../build/bin/seerlink
```

**Test:**

```
# A higher parallel number can speed up tests at the expense of more RAM.
go test -p 1 ./...
```

The golang testsuite is almost entirely parallelizable, and so running the default
`go test ./...` will commonly peg your processor. Limit parallelization with the
`-p 2` or whatever best fits your computer: `go test -p 4 ./...`.
