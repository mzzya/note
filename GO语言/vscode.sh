#! /bin/bash

go get -u golang.org/x/lint/golint
go get -u github.com/stamblerre/gocode
mv $GOPATH/bin/gocode $GOPATH/bin/gocode-gomod
go get -u github.com/mdempsky/gocode
go get -u github.com/uudashr/gopkgs/cmd/gopkgs
go get -u github.com/ramya-rao-a/go-outline
go get -u github.com/acroca/go-symbols
go get -u golang.org/x/tools/cmd/guru
go get -u golang.org/x/tools/cmd/gorename
go get -u github.com/go-delve/delve/cmd/dlv

go get -u github.com/rogpeppe/godef
go get -u golang.org/x/tools/cmd/goimports
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
go get -u golang.org/x/tools/cmd/gopls
go get -u github.com/cweill/gotests/gotests
go get -u github.com/fatih/gomodifytags
go get -u github.com/josharian/impl
go get -u github.com/davidrjenni/reftools/cmd/fillstruct
go get -u github.com/haya14busa/goplay/cmd/goplay
go get -u github.com/godoctor/godoctor
go get -u github.com/mgechev/revive
go get -u honnef.co/go/tools/cmd/staticcheck
go get -u github.com/zmb3/gogetdoc

ll $GOPATH/bin/