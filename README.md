# Seerlink

Node of the SeerLink decentralized oracle network, bridging on and off-chain computation. 


[Seerlink](https://seer.link/) is middleware to simplify communication with blockchains.
Here you'll find the Seerlink Golang node, currently in alpha.
This initial implementation is intended for use and review by developers,
and will go on to form the basis for Seerlink's decentralized oracle network.
Further development of the Seerlink Node and Seerlink Network will happen here.

## Features

- easy connectivity of on-chain contracts to any off-chain computation or API
- multiple methods for scheduling both on-chain and off-chain computation for a user's smart contract
- automatic gas price bumping to prevent stuck transactions, assuring your data is delivered in a timely manner
- push notification of smart contract state changes to off-chain systems, by tracking Ethereum logs
- translation of various off-chain data types into EVM consumable types and transactions
- easy to implement smart contract libraries for connecting smart contracts directly to their preferred oracles
- easy to install node, which runs natively across operating systems, blazingly fast, and with a low memory footprint


## Install

1. [Install Go 1.14](https://golang.org/doc/install?download=go1.14.9.darwin-amd64.pkg), and add your GOPATH's [bin directory to your PATH](https://golang.org/doc/code.html#GOPATH)
   - Example Path for macOS `export PATH=$GOPATH/bin:$PATH` & `export GOPATH=/Users/$USER/go`
2. Install [NodeJS 12.18](https://nodejs.org/en/download/package-manager/) & [Yarn](https://yarnpkg.com/lang/en/docs/install/)
   - It might be easier long term to use [nvm](https://nodejs.org/en/download/package-manager/#nvm) to switch between node versions for different projects: `nvm install 12.18 && nvm use 12.18`
3. Install [Postgres (>= 11.x)](https://wiki.postgresql.org/wiki/Detailed_installation_guides).
   - You should [configure Postgres](https://www.postgresql.org/docs/12/ssl-tcp.html) to use SSL connection
4. Download Seerlink: `git clone https://github.com/SeerLink/seerlink && cd seerlink`
5. Build Seerlink: `make`
   - If you got any errors regarding locked yarn package, try running `yarn install` before this step
6. Run the node: `seerlink help`


## Run

**NOTE**: By default, seerlink will run in TLS mode. For local development you can either disable this by setting SEERLINK_DEV to true.

To start your Seerlink node, simply run:

```bash
seerlink node start
```

By default this will start on port 6688.

Once your node has started, you can view your current jobs with:

```bash
seerlink job_specs list # v1 jobs
seerlink jobs list # v2 jobs
```

View details of a specific job with:

```bash
seerlink job_specs show "$JOB_ID # v1 jobs"
```

To find out more about the Seerlink CLI, you can always run `seerlink help`.

### Build your current version

```bash
go build -o seerlink ./core/
```

- Run the binary:

```bash
./seerlink
```

### Use of Go Generate

Go generate is used to generate mocks in this project. Mocks are generated with [mockery](https://github.com/vektra/mockery) and live in core/internal/mocks.

## Contributing

Seerlink's source code is [licensed under the MIT License](./LICENSE), and contributions are welcome.

Thank you!

## License

[MIT](https://choosealicense.com/licenses/mit/)
