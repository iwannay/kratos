package main

import (
	"os/exec"
)

const (
	_getGRPCGenFast = "go get -u github.com/gogo/protobuf/protoc-gen-gofast"
	_grpcProtocFast = `protoc --proto_path=%s --proto_path=%s --proto_path=%s --gofast_out=plugins=grpc:.`

	_getGRPCGen = "go get -u github.com/gogo/protobuf/protoc-gen-gogo"
	_grpcProtoc = `protoc --proto_path=%s --proto_path=%s --proto_path=%s --gogo_out=plugins=grpc:.`

	_getGRPCGenOfficial = "go get -u github.com/golang/protobuf/protoc-gen-go"
	_grpcProtocOfficial = `protoc --proto_path=%s --proto_path=%s --proto_path=%s --go_out=plugins=grpc:.`
)

func installGRPCGen() error {
	if _, err := exec.LookPath("protoc-gen-gofast"); err != nil {
		if err := goget(_getGRPCGenFast); err != nil {
			return err
		}
	}
	if _, err := exec.LookPath("protoc-gen-go"); err != nil {
		if err := goget(_getGRPCGenOfficial); err != nil {
			return err
		}
	}
	if _, err := exec.LookPath("protoc-gen-gogo"); err != nil {
		if err := goget(_getGRPCGen); err != nil {
			return err
		}
	}
	return nil
}

func genGRPC(files []string) error {
	return generate(_grpcProtoc, files)
}
