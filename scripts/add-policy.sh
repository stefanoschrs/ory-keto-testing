#!/bin/bash

docker run -it --rm \
  --network ketoguide \
  -v $(pwd)/policies:/policies \
  -e KETO_URL=http://ory-keto-example--keto:4466/ \
  oryd/keto:v0.5.7-alpha.1 \
  engines acp ory policies import exact /policies/example-policy.json

