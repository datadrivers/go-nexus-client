Golang Nexus Client
==========

# Introduction

Sonatype Nexus Golang Client

# Development and testing
                         
Start a Nexus Docker container pre-configured for the tests (plugins installed and password set to `admin123`)

```shell
$ scripts/start-nexus.sh
```

Set config as environment variables

```shell
$ export NEXUS_URL=http://127.0.0.1:8081
$ export NEXUS_USERNAME=admin
$ export NEXUS_PASSWORD=admin123
```

Run tests

```shell
$ make test
```

Stop nexus

```shell
$ docker rm -f nexus
```

*NOTE*: Some tests do not clean up reliably. You may need to restart Nexus to get a clean test run

# Author

[Datadrivers GmbH](https://www.datadrivers.de)
