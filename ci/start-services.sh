#!/bin/bash -e

echo "Starting services..."
docker-compose up -d

echo -n "Setting admin password for Nexus..."
i=1
until wget -t 1 http://localhost:8081 -O /dev/null -q
do
    sleep 1
    echo -n .
    if [[ $((i%3)) == 0 ]]; then echo -n ' '; fi
    (( i++ ))
done

CURRENT_NEXUS_ADMIN_PASSWORD=$(docker exec -ti nexus cat /nexus-data/admin.password)
curl -X PUT "http://127.0.0.1:8081/service/rest/beta/security/users/admin/change-password" -H \
  "accept: application/json" -H "Content-Type: text/plain" -d "p455w0rd" -u "admin:${CURRENT_NEXUS_ADMIN_PASSWORD}"
