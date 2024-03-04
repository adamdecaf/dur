# dur

[![GoDoc](https://godoc.org/github.com/adamdecaf/dur?status.svg)](https://godoc.org/github.com/adamdecaf/dur)
[![Build Status](https://github.com/adamdecaf/dur/workflows/Go/badge.svg)](https://github.com/adamdecaf/dur/actions)
[![Coverage Status](https://codecov.io/gh/adamdecaf/dur/branch/master/graph/badge.svg)](https://codecov.io/gh/adamdecaf/dur)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamdecaf/dur)](https://goreportcard.com/report/github.com/adamdecaf/dur)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/adamdecaf/dur/master/LICENSE)

CLI tool that prints time programs take to run.

## Install

Download the [latest release for your architecture](https://github.com/adamdecaf/dur/releases/latest).

## Usage

```
$ dur sleep 11
  command has taken 10s
  command took 11s
```

`dur` allows for a configurable `-tick` period.

```
$ dur -tick 5s sleep 12
  command has taken 5s
  command has taken 10s
  command took 12s
```

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
