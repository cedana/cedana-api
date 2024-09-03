#/bin/bash

# Use this to generate from proto files after updating the submodules.
# You can update the submodule with git submodule update --init --recursive

PROTO_DIR=".."

# Generate Go code for gpu.proto
protoc --go_out=gpu --go_opt=paths=source_relative \
    --go-grpc_out=gpu --go-grpc_opt=paths=source_relative \
    --go_opt=Mgpu.proto=github.com/cedana/cedana-api/go/gpu \
    --go-grpc_opt=Mgpu.proto=github.com/cedana/cedana-api/go/gpu \
    -I$PROTO_DIR \
    $PROTO_DIR/gpu.proto

# Generate Go code for image.proto
protoc --go_out=image --go_opt=paths=source_relative \
    --go-grpc_out=image --go-grpc_opt=paths=source_relative \
    --go_opt=Mimage.proto=github.com/cedana/cedana-api/go/image \
    --go-grpc_opt=Mimage.proto=github.com/cedana/cedana-api/go/image \
    -I$PROTO_DIR \
    $PROTO_DIR/image.proto

# Generate Go code for img-streamer.proto
protoc --go_out=img_streamer --go_opt=paths=source_relative \
    --go-grpc_out=img_streamer --go-grpc_opt=paths=source_relative \
    --go_opt=Mimg_streamer.proto=github.com/cedana/cedana-api/go/img-streamer \
    --go-grpc_opt=Mimg_streamer.proto=github.com/cedana/cedana-api/go/img-streamer \
    -I$PROTO_DIR \
    $PROTO_DIR/img-streamer.proto

# Generate Go code for rpc.proto
protoc --go_out=rpc --go_opt=paths=source_relative \
    --go-grpc_out=rpc --go-grpc_opt=paths=source_relative \
    --go_opt=Mrpc.proto=github.com/cedana/cedana-api/go/rpc \
    --go-grpc_opt=Mrpc.proto=github.com/cedana/cedana-api/go/rpc \
    -I$PROTO_DIR \
    $PROTO_DIR/rpc.proto

# Generate Go code for task.proto
protoc --go_out=task --go_opt=paths=source_relative \
    --go-grpc_out=task --go-grpc_opt=paths=source_relative \
    --go_opt=Mgpu.proto=github.com/cedana/cedana-api/go/gpu \
    --go_opt=Mtask.proto=github.com/cedana/cedana-api/go/task \
    --go-grpc_opt=Mgpu.proto=github.com/cedana/cedana-api/go/gpu \
    --go-grpc_opt=Mtask.proto=github.com/cedana/cedana-api/go/task \
    -I$PROTO_DIR \
    $PROTO_DIR/task.proto
