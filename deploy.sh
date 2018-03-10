#!/bin/bash

set -e

aws lambda update-function-code \
    --function-name lhl-me \
    --zip-file fileb://endpoints/me/deployment.zip

aws lambda update-function-code \
    --function-name lhl-offers \
    --zip-file fileb://endpoints/offers/deployment.zip

aws lambda update-function-code \
    --function-name lhl-offers-id \
    --zip-file fileb://endpoints/offers/offer-id/deployment.zip

aws lambda update-function-code \
    --function-name lhl-offers-reservation \
    --zip-file fileb://endpoints/offers/offer-id/reservation/deployment.zip
