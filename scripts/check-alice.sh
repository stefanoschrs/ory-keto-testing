#!/bin/bash

docker run -it --rm \
  --network ketoguide \
  -e KETO_URL=http://ory-keto-example--keto:4466/ \
  oryd/keto:v0.5.7-alpha.1 \
  engines acp ory allowed exact alice blog_posts:my-first-blog-post delete

