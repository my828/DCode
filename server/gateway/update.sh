#!/bin/bash

update() {
    docker image prune -f
    docker volume prune -f
    docker container prune -f

    docker "Removing existing decodeNetwork..."
    docker network rm decodeNetwork

    docker rm sessions -f
    docker rm dcodeGateway -f

    # pull most recent container images from Dockerhub
    docker pull my828/dcode-gateway

    docker network create dcodeNetwork

    # set environment variables
    GATEWAYADDRESS=":4000" # change later
    MONGOADDRESS=""
    RABBITADDRESS=""
    RABBITHOSTNAME=""
    RABBITMQNAME=""
    REDISADDRESS="sessions:6379"
    SIGNINGKEY=""
    TLSCERT=""
    TLSKEY=""

    echo "Running Redis container..."
    docker run -d --name sessions --network dcodeNetwork redis

    echo "Waiting for all dependencies to boot up..."
    sleep 5s

    #  -v /etc/letsencrypt:/etc/letsencrypt:ro \
    echo "Running decode-gateway container..."
    docker run -d --name dcodeGateway \
    -p 4000:4000 \
    -e TLSCERT=$TLSCERT \
    -e TLSKEY=$TLSKEY \
    -e SIGNINGKEY=$SIGNINGKEY \
    -e REDISADDRESS=$REDISADDRESS \
    -e GATEWAYADDRESS=$GATEWAYADDRESS \
    --network dcodeNetwork \
    my828/dcode-gateway

    echo "Done."
}

update