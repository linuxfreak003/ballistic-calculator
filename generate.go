package generate

// PROTOS
//go:generate go install github.com/bufbuild/buf/cmd/buf@latest
//go:generate env PATH=$PATH:$GOPATH/bin $GOPATH/bin/buf mod update
//go:generate env PATH=$PATH:$GOPATH/bin $GOPATH/bin/buf generate
