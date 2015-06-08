#!/bin/bash

go test -coverprofile=client.cover.out -coverpkg=./... ./client
go test -coverprofile=enablers.cover.out -coverpkg=./... ./enablers
go test -coverprofile=obj_oma.cover.out -coverpkg=./... ./objdefs/oma
go test -coverprofile=objects.cover.out -coverpkg=./... ./objects
go test -coverprofile=registry.cover.out -coverpkg=./... ./registry
go test -coverprofile=request.cover.out -coverpkg=./... ./request
go test -coverprofile=resources.cover.out -coverpkg=./... ./resources
go test -coverprofile=response.cover.out -coverpkg=./... ./response
go test -coverprofile=utils.cover.out -coverpkg=./... ./utils
go test -coverprofile=values.cover.out -coverpkg=./... ./values
go test -coverprofile=values.cover.out -coverpkg=./... ./values/tlv

echo "mode: set" > coverage.out && cat *.cover.out | grep -v mode: | sort -r | awk '{if($1 != last) {print $0;last=$1}}' >> coverage.out
go tool cover -html=coverage.out
rm *.out