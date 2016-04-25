#!/bin/bash - 
killall -9 server
./bin/server &
./bin/client
