#!/bin/bash -xe
#
# Curl example showing token auth header
curl -i -X POST \
  -H 'Content-Type: application/json' \
  -H 'Api-Key: YOUR_API_KEY_HERE' \
  https://controller.YOUR_ENVIRONMENT.pinecone.io/databases \
  -d '{
    "name": "auth-guide",
    "dimension": 8,
    "metric": "euclidean"
  }'