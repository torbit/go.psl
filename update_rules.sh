#!/bin/bash

OUT="psl/psl_rules.go"

set -e

cd $(dirname $0)

echo ": build rulegen"
(cd rulegen; go build)
echo ": rulegen"
./rulegen/rulegen -out="$OUT"
echo ": fmt"
go fmt $OUT

