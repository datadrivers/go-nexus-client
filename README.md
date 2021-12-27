# Golang Nexus Client

![codeql workflow](https://github.com/datadrivers/go-nexus-client/actions/workflows/codeql-analysis.yml/badge.svg)

## Introduction

Sonatype Nexus Golang Client

## Development and testing

**NOTE**: For testing Nexus Pro features, place the `license.lic` in `scripts/`.

For testing start a local Docker containers using make

```shell
make start-services
```

This will start a Docker and MinIO containers and expose ports 8081 and 9000.

Now start the tests

```shell
$ make test
```

The tests assume Nexus Pro features. If you do not have a Nexus Pro license you can skip the pro tests by setting the `SKIP_PRO_TESTS` environment variable:

```shell
$ SKIP_PRO_TESTS=true make test
```

To `SKIP_AZURE_TESTS` environment variable:

```shell
$ SKIP_AZURE_TESTS=true make test
```

## Author

[Datadrivers GmbH](https://www.datadrivers.de)
