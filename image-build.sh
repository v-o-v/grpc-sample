#!/bin/bash

if [ $# -ne 1 ]; then
  echo "指定された引数は$#個です。" 1>&2
  echo "実行するには1個の引数が必要です。" 1>&2
  exit 1
fi

PROJECT=grpc-sample

IMAGE=$PROJECT:$1

docker build -t $IMAGE -f docker/Dockerfile.stg .
docker tag $IMAGE asia.gcr.io/ignition-heart/$IMAGE
docker push asia.gcr.io/ignition-heart/$IMAGE

exit 0
