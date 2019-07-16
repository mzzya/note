#!/bin/bash
#代码检查
go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint
#swag生成
go get -u -v github.com/swaggo/swag/cmd/swag
#json转go结构体
go get -u -v github.com/ChimeraCoder/gojson/gojson
# go项目脚手架
go get -u -v github.com/spf13/cobra/cobra
# go 分析工具
go get -u -v github.com/google/pprof
# 测试
go get -u -v github.com/smartystreets/goconvey
# grpc
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u -v google.golang.org/grpc
# 分析go项目二进制文件大小
go get -u -v github.com/jondot/goweight