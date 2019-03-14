#!/bin/bash

deploy() {
    bash ./build.sh

    echo "Updating dcode-gateway image on DockerHub..."
    docker push harshiakkaraju/dcode-gateway

    # TODO -- ssh into ec2 instance
    bash ./update.sh
}

deploy
