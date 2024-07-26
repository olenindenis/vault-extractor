# GO-Vault-Extractor

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/olenindenis/vault-extractor.svg)](https://pkg.go.dev/github.com/olenindenis/vault-extractor)
![Build](https://github.com/olenindenis/vault-extractor/actions/workflows/main.yml/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/olenindenis/vault-extractor)](https://goreportcard.com/report/github.com/olenindenis/vault-extractor)

Go Vault Extractor is a complete solution for efficient and fast import secrets from hashicorp vault server.

## Installation

```sh
go get -u github.com/olenindenis/vault-extractor
```

## Examples

### Getting Started

1. To import envs as env file use:
```sh
extractor env
```

2. To import envs as json file use:
```sh
extractor json
```

3. To run extractor with .env file as vault connection config:
```sh
extractor json -conf=.env
```

4. To run extractor with OS environment variables as vault connection config do not use -conf param

5. To save imported envs as env or json file use -file param:
```sh
extractor json -conf=.env -file=tst.json
```
or
```sh
extractor env -conf=.env -file=.env.dev
```

If your -conf=.env and -file=.env equal than extractor will add new envs to old env file
