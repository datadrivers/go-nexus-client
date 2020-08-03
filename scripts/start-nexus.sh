#!/bin/bash -e

LOCAL_PORT=8081
NEXUS_VERSION=3.25.1

echo "Starting Nexus container..."
docker run -d --rm -v "${PWD}/nexus.properties":/nexus-data/etc/nexus.properties \
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

echo "Getting admin password..."
NEXUS_ADMIN_PASSWORD=$(docker exec -ti nexus cat /nexus-data/admin.password)

echo "Setting admin password..."
curl -X PUT "http://127.0.0.1:${LOCAL_PORT}/service/rest/beta/security/users/admin/change-password" -H "accept: application/json" -H "Content-Type: text/plain" -d "admin123" -u "admin:${NEXUS_ADMIN_PASSWORD}"