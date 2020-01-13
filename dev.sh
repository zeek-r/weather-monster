#!/bin/bash
go test && \
go build -o weather-monster && \
set -o allexport; \
source ./run.env; \
set +o allexport; \
./weather-monster start

