#!/bin/bash

. .env

docker run \
  --rm \
  --name ory-keto-example--keto \
  --network ketoguide \
  -p 4466:4466 \
  -e DSN=${DSN} \
  oryd/keto:v0.5.7-alpha.1 \
  serve

