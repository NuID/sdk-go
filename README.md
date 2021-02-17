<p align="right"><a href="https://nuid.io"><img src="https://nuid.io/svg/logo.svg" width="20%"></a></p>

# NuID SDK for Go

[![](https://img.shields.io/github/go-mod/go-version/NuID/sdk-go/main?color=red&label=go.mod&logo=go&logoColor=white&style=for-the-badge)](https://github.com/NuID/sdk-go)
[![](https://img.shields.io/badge/docs-latest-blue?style=for-the-badge&logo=read-the-docs)](https://pkg.go.dev/github.com/NuID/sdk-go/api/auth)
[![](https://img.shields.io/badge/docs-platform-purple?style=for-the-badge&logo=read-the-docs)](https://portal.nuid.io/docs)

This repo provides a library for interacting with NuID APIs within go
applications.

Read the latest [package docs](https://pkg.go.dev/github.com/NuID/sdk-go/api/auth) or
checkout the [platform docs](https://portal.nuid.io/docs) for API docs, guides,
video tutorials, and more.

## Install

```sh
GO111MODULE=on go get github.com/NuID/sdk-go
```

## Usage

A fully working go server example can be found in our [examples
repo](https://github.com/NuID/examples/tree/main/go).

Visit the [Integrating with
NuID guide](https://portal.nuid.io/docs/guides/integrating-with-nuid) for a
detailed walk-through on your first NuID integration.

## Development

You'll want to download docker to run the tests, as we depend on the `@nuid/cli`
npm package to provide a CLI you can shell out to in the tests for generating zk
crypto. After checking out the repo, run `make build run` to install
dependencies and create the docker environment. Then, run `make test` to run the
tests inside the running container. You can also run `make shell` to get a
prompt in the container.

`make clean` will stop and destroy the container and image. `make build run`
will rebuild the image and run the container.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/NuID/sdk-go.

## License

The library is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
