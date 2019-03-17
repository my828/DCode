#!/bin/bash

update() {
    docker image prune -f
    docker volume prune -f
    docker container prune -f

    docker "Removing existing decodeNetwork..."
    docker network rm decodeNetwork

    docker rm rabbit -f
    docker rm sessions -f
    docker rm dcodeGateway -f

    # pull most recent container images from Dockerhub
    docker pull harshiakkaraju/dcode-gateway

    docker network create dcodeNetwork

    # set environment variables
    GATEWAYADDRESS=":4000" # change later
    MONGOADDRESS=""
    RABBITNAME="rabbit"
    RABBITADDRESS="amqp://$RABBITNAME:5672/"
    REDISADDRESS="sessions:6379"
    SIGNINGKEY=""
    TLSCERT=""
    TLSKEY=""
    NETWORK="dcodeNetwork"

    docker run -d \
    --name $RABBITNAME \
    --network $NETWORK rabbitmq

    echo "Running Redis container..."
    docker run -d --name sessions --network $NETWORK redis

    echo "Waiting for all dependencies to boot up..."
    sleep 20s

    #  -v /etc/letsencrypt:/etc/letsencrypt:ro \
    echo "Running decode-gateway container..."
    docker run -d --name dcodeGateway \
    -p 4000:4000 \
    -e TLSCERT=$TLSCERT \
    -e TLSKEY=$TLSKEY \
    -e SIGNINGKEY=$SIGNINGKEY \
    -e REDISADDRESS=$REDISADDRESS \
    -e GATEWAYADDRESS=$GATEWAYADDRESS \
    -e RABBITADDRESS=$RABBITADDRESS \
    -e RABBITNAME=$RABBITNAME \
    --network $NETWORK \
    harshiakkaraju/dcode-gateway

    echo "Done."
}

update