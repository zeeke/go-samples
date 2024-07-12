#!/bin/bash

set -x

# Parameter `flake-attempts=N` overrides spec decorator
go test ./sginkgo
go test ./sginkgo --ginkgo.flake-attempts=3

# Expected to fail with
# Ginkgo detected an issue with your spec structure
# var _ = BeforeSuite(func() {
# /home/apanatto/dev/github.com/zeeke/go-samples/ginkgonestedsuite/suite_test.go:18
#   It looks like you are trying to add a [BeforeSuite] node but
#   you already have a [BeforeSuite] node defined at:
#   /home/apanatto/dev/github.com/zeeke/go-samples/ginkgonestedsuite/nested/test_one.go:8.
# 
#   Ginkgo only allows you to define one suite setup node.
go test -v ./ginkgonestedsuite


go test -v ./ginkgofailhandler -ginkgo.v

go test -v ./multipleskips --ginkgo.skip "D1 A" --ginkgo.skip "D1 B"


exit 0
