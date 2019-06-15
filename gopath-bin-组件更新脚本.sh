#!/bin/bash
#代码检查
go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint
#swag生成
go get -u -v github.com/swaggo/swag/cmd/swag
#json转go结构体
go get -u -v github.com/ChimeraCoder/gojson/gojson

go get -u -v github.com/spf13/cobra/cobra

go get -u -v github.com/google/pprof

go get -u -v github.com/smartystreets/goconvey