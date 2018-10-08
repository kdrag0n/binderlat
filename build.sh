#!/usr/bin/env bash
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" && goupx binderlat > /dev/null 2>&1
