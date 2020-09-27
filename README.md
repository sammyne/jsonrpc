# jsonrpc

This is a fork of [gorilla/rpc].

![build](https://github.com/sammyne/jsonrpc/workflows/build/badge.svg)
[![codecov](https://codecov.io/gh/sammyne/jsonrpc/branch/master/graph/badge.svg)](https://codecov.io/gh/sammyne/jsonrpc)
[![LICENSE](https://img.shields.io/badge/license-ISC-blue.svg)](LICENSE)

## Installation

```bash
go get -u github.com/sammyne/jsonrpc
```

## Examples
- [server][server-example]: a server providing API as `HelloService`
- [client][client-example]: a client request service provided by [server][server-example]

[client-example]: examples/client.go
[server-example]: examples/server.go
[gorilla/rpc]: https://pkg.go.dev/github.com/gorilla/rpc/v2
