#!/usr/bin/env bash

set -eo pipefail

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
    protoc \
    -I "proto" \
    -I "third_party/proto" \
    --gocosmos_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
    --grpc-gateway_out=logtostderr=true:. \
    $(find "${dir}" -maxdepth 1 -name '*.proto')
done

protoc \
-I "proto" \
-I "third_party/proto" \
--descriptor_set_out="injectived.protoset" \
--include_imports \
$(find ./proto ./third_party -maxdepth 16 -name 'query.proto')

# move proto files to the right places
cp -r github.com/InjectiveLabs/injective-core/* ./
rm -rf github.com
