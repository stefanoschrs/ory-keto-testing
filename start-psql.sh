#!/bin/bash

. .env

docker run \
  --rm \
  --name ory-keto-example--postgres \
  --network ${network} \
  -e POSTGRES_USER=${dbUser} \
  -e POSTGRES_PASSWORD=${dbPass} \
  -e POSTGRES_DB=${dbName} \
  postgres:latest
