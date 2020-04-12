#!/bin/bash

VERSION=0.0.1

docker build -t alexandriaeam/k8s-pharos:"$VERSION" .
docker push alexandriaeam/k8s-pharos:"$VERSION"