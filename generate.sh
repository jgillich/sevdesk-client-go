#!/bin/sh

curl -O https://api.sevdesk.de/OpenAPI/Index/output.yaml
openapi-generator-cli generate --git-repo-id sevdesk-client-go --git-user-id jgillich -g go -o out -i output.yaml  --additional-properties=packageName=sevdesk
rm output.yaml
