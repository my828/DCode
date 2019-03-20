#!/bin/bash

build () {
    echo "Building Go executable for Linux..."
    GOOS=linux go build

    echo "Building Docker container for gateway..."
    docker build -t my828/dcode-gateway .

    echo "Cleaning Go executable for Linux"
    go clean
}

build