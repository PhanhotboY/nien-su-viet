#! /bin/sh

service=$1
tag=$2
shift 2
command="docker build -f local/$service-service/Dockerfile -t phanhotboy/nsv-$tag $@ ."
echo running command $command
$command
