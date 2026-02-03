#! /bin/sh

type=$1
command=$2
shift 2
docker compose -f local/docker-compose.$type.yaml ${command:-up} $@
