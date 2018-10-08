#!/usr/bin/env bash
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" && goupx binderlat > /dev/null 2>&1 && adb push binderlat /data/local/tmp > /dev/null && adb shell chmod 755 /data/local/tmp/binderlat && adb shell su -c /data/local/tmp/binderlat
