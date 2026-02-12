#! /bin/sh

tag=$1
shift 1
command="docker image remove phanhotboy/nsv-$tag $@"
echo running command $command
$command
