#!/usr/bin/env bash

# https://regex101.com/r/ZRPKqy/2

ngrok http 9090 > /dev/null &

sleep 2

curl -s localhost:4040/api/tunnels | grep -Eo "\"public_url\":\"http:([^\"]+)" | grep -Eo "http:([^\"]+)"
