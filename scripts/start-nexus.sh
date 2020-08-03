#!/bin/bash -e

LOCAL_PORT=8081
NEXUS_VERSION=3.24.0

echo "Starting Nexus container..."
docker run -d --rm -v "${PWD}"/nexus.properties:/nexus-data/etc/nexus.properties \
  --name nexus -p 127.0.0.1:${LOCAL_PORT}:8081 sonatype/nexus3:${NEXUS_VERSION}

echo -n "Waiting for Nexus to be ready "
i=1
until wget -t 1 http://127.0.0.1:${LOCAL_PORT} -O /dev/null -q
do
    sleep 1
    echo -n .
    if [[ $((i%3)) == 0 ]]; then echo -n ' '; fi
    (( i++ ))
done
