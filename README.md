## Go Laniakea

Official Golang implementation of the Laniakea protocol.

Automated builds are available for stable releases and the unstable master branch. 

## Building the source

Building `lkea` requires both a Go (version 1.20 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make lkea
```

or, to build the full suite of utilities:

```shell
make all
```

## Executables

The go-laniakea project comes with several wrappers/executables found in the `cmd`
directory.

|    Command    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| :-----------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  **`lkea`**   | Our main Laniakea CLI client. It is the entry point into the Laniakea network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Laniakea network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. 
|   `clef`    | Stand-alone signing tool, which can be used as a backend signer for `lkea`.  |

## Running `lkea`

Going through all the possible command line flags is out of scope here (please consult our
[CLI Wiki page](https://github.com/LaniakeaChain/wiki)),
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `lkea` instance.

### Hardware Requirements

Minimum:

* CPU with 2+ cores
* 4GB RAM
* 1TB free storage space to sync the Mainnet
* 8 MBit/sec download Internet service

Recommended:

* Fast CPU with 4+ cores
* 16GB+ RAM
* High Performance SSD with at least 1TB free space
* 25+ MBit/sec download Internet service

### Full node on the main Laniakea network

By far the most common scenario is people wanting to simply interact with the Laniakea
network: create accounts; transfer funds; deploy and interact with contracts. For this
particular use-case the user doesn't care about years-old historical data, so we can
sync quickly to the current state of the network. To do so:

```shell
$ lkea console
```

## Contribution

Thank you for considering to help out with the source code! We welcome contributions
from anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to go-laniakea, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base. 
Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting)
   guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary)
   guidelines.
 * Pull requests need to be based on and opened against the `master` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "eth, rpc: make trace configs optional"

Please see the [Developers' Guide](https://github.com/LaniakeaChain/wiki)
for more details on configuring your environment, managing project dependencies, and
testing procedures.

## License

- TODO
