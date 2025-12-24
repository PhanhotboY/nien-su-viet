#! /bin/sh

docker compose -f local/docker-compose.$1.yaml ${2:-up}
