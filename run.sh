#!/bin/bash

set -x

# Parameter `flake-attempts=N` overrides spec decorator
go test ./sginkgo
go test ./sginkgo --ginkgo.flake-attempts=3


go test -v ./ginkgofailhandler -ginkgo.v

go test -v ./multipleskips --ginkgo.skip "D1 A" --ginkgo.skip "D1 B"

exit 0
