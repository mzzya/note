# 代码检查
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
# swag文档生成
go get -u github.com/swaggo/swag/cmd/swag
# 性能分析工具
go get -u github.com/google/pprof
# go脚手架
go get -u github.com/spf13/cobra/cobra
# 分析go项目二进制文件大小
go get -u github.com/jondot/goweight
# 静态检查工具 https://staticcheck.io
go get -u honnef.co/go/tools/cmd/staticcheck
# grpc
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
# linter
go get -u github.com/mgechev/revive
# Simple utility for extracting a JSON representation of the declarations in a Go source file.
go get -u github.com/ramya-rao-a/go-outline

go get -u github.com/smartystreets/goconvey