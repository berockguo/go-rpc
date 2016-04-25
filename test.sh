#!/bin/bash -
killall -9 $1server
./bin/$1server &
./bin/$1client
