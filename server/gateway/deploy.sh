#!/bin/bash

deploy() {
    bash ./build.sh

    echo "Updating dcode-gateway image on DockerHub..."
    docker push my828/dcode-gateway

    # TODO -- ssh into ec2 instance
    bash ./update.sh
}

deploy
