#!/bin/bash

PASS=$1

if [ -z "$PASS" ]; then
	echo "You must pass a password as an argument!"
	exit 1
fi

echo "Starting redis container: shopping"

if [ "$(docker ps -aq -f "name=shopping")" ]; then
	if [ "$(docker ps -aq -f "status=exited" -f "name=shopping")" ]; then
		docker container start shopping
	else
		echo "Redis instance (shopping) is already running!"
	fi
else
	docker run -d -p 6379:6379 -v ~/.redis_data/shopping:/data --name shopping redis \
		redis-server --appendonly yes --requirepass "$PASS"
fi
