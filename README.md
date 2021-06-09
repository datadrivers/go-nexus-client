Golang Nexus Client
==========

# Introduction

Sonatype Nexus Golang Client

# Development and testing

Start a Nexus Docker container

```shell
$ docker run -d -p 8081:8081 --name nexus sonatype/nexus3
```

Get the Admin password

```shell
$ docker exec -ti nexus /bin/bash -c 'cat /nexus-data/admin.password'
```

Set config as environment variables

```shell
$ export NEXUS_URL=http://127.0.0.1:8081
$ export NEXUS_USERNAME=admin
$ export NEXUS_PASSWORD=<random-password-from-above>
```

Run tests

```shell
$ make test
```
  
To test the Nexus Pro features you will need to install a Nexus Pro license, restart nexus and then enable the additional tests:

```shell
$ export SKIP_PRO_TESTS=false
```

# Author

[Datadrivers GmbH](https://www.datadrivers.de)
