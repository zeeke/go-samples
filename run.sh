#!/bin/bash

set -x

# Parameter `flake-attempts=N` overrides spec decorator
go test ./sginkgo
go test ./sginkgo --ginkgo.flake-attempts=3


go test -v ./ginkgofailhandler -ginkgo.v

exit 0
