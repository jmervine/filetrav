travis:
	go test -test.v

test: fmt
	go test

fmt:
	gofmt -tabs=false -tabwidth=4 -w -l -s *.go

readme: test
	# generating readme
	godoc -ex -v -templates "$(PWD)/_support" . | tee README.md

.PHONY: travis test fmt readme
