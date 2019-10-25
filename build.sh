#!/usr/bin/env bash
set +x
set -u
set -e

go build -v -ldflags="-X 'github.com/gambledor/remote-connector/build.Version=v0.9.0' \
    -X 'github.com/gambledor/remote-connector/build.Build=`git rev-parse --short HEAD`' \
    -X 'github.com/gambledor/remote-connector/build.User=$(id -u -n)' \
    -X 'github.com/gambledor/remote-connector/build.Time=$(date)'" \
    cmd/remoteconnector.go
