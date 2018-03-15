#!/usr/bin/env bash
set -o errexit -o pipefail -o verbose
go version
go test -c -o failfast_testbin
./failfast_testbin -test.v -test.failfast
