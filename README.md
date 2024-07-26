# GO-Client-Extractor

[![Go Reference][go-reference-badge]][go-reference]
[![Build][ci-build-badge]][ci-build]

Go Client Extractor is a complete solution for efficient and fast import secrets from hashicorp vault server.

## Installation

```sh
go get -u github.com/olenindenis/vault-extractor
```


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

### If your -conf=.env and -file=.env equal than extractor will add new envs to old env file
