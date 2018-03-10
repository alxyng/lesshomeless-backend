#!/bin/bash

set -e

declare -a endpoints=("me" "offers" "offers/offer-id")

for i in "${endpoints[@]}"
do
	cd endpoints/$i
	GOOS=linux go build -o main

	if [[ $(uname) == "MINGW"* ]]; then
		~/go/bin/build-lambda-zip.exe -o deployment.zip main
	else
		zip deployment.zip main
	fi

	cd -
done
