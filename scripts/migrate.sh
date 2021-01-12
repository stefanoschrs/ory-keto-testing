#!/bin/bash

DSN=postgres://keto:secret@ory-keto-example--postgres:5432/keto?sslmode=disable

docker run \
  -it --rm \
  --network ketoguide \
  -e DSN=${DSN} \
  oryd/keto:v0.5.7-alpha.1 \
  migrate sql -e

