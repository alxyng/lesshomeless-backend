#!/bin/bash

set -e

cd endpoints/me

GOOS=linux go build -o main
# zip deployment.zip main

cd -
