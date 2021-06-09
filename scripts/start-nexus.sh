#!/bin/bash -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

LOCAL_PORT=8081
NEXUS_VERSION=3.29.0
GOOGLE_BLOBSTORE_VERSION=0.19.1

if [ ! -d "${SCRIPT_DIR}/deploy" ] ; then
  mkdir "${SCRIPT_DIR}/deploy"
fi

if [ ! -f  "${SCRIPT_DIR}/deploy/nexus-blobstore-google-cloud-${GOOGLE_BLOBSTORE_VERSION}.kar" ]; then
  rm -f "${SCRIPT_DIR}/deploy/nexus-blobstore-google-cloud-"*.kar
  echo "Downloading Nexus Google Blobstore Plugin..."
  curl -fsSL -o "${SCRIPT_DIR}/deploy/nexus-blobstore-google-cloud-${GOOGLE_BLOBSTORE_VERSION}.kar" https://repo.maven.apache.org/maven2/org/sonatype/nexus/plugins/nexus-blobstore-google-cloud/$GOOGLE_BLOBSTORE_VERSION/nexus-blobstore-google-cloud-$GOOGLE_BLOBSTORE_VERSION.kar
fi


echo "Starting Nexus container..."
docker run -d --rm -v "${SCRIPT_DIR}/nexus.properties":/nexus-data/etc/nexus.properties \
  -v "${SCRIPT_DIR}/deploy":/opt/sonatype/nexus/deploy \
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
