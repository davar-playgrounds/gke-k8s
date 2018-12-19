#!/usr/bin/env bash

echo "Waiting for file:"
while ! [[ -f "/app/data.json" ]]; do
    echo -ne "#"
    sleep 1
done
echo " - File loaded."

queryMongoVersion() {
  mongo \
     --authenticationDatabase "admin" \
     --username "${MONGO_USERNAME:-admin}" \
     --password "${MONGO_PASSWORD:-password}" \
     --host "${MONGO_DOMAIN:-127.0.0.1}" \
     --port "${MONGO_PORT:-27017}" \
     --eval "db.version()"
}

while ! queryMongoVersion > /dev/null 2>&1; do sleep 1; done

mongoimport \
 --authenticationDatabase "admin" \
 --username "${MONGO_USERNAME:-admin}" \
 --password "${MONGO_PASSWORD:-password}" \
 --host "${MONGO_DOMAIN:-127.0.0.1}" \
 --port "${MONGO_PORT:-27017}" \
 --db "${MONGO_DB:-default}" \
 --collection "${MONGO_COLLECTION:-app}" \
 --type "json" \
 --file "/app/data.json" \
 --jsonArray
