protoc -I . \
  --go_out . \
  --go_opt paths=source_relative \
  --go-grpc_out . --go-grpc_opt paths=source_relative \
  --swagger_out logtostderr=true:. \
  proto/echopb/echo_server.proto

protoc -I . \
  --go_out . \
  --go_opt paths=source_relative \
  --go-grpc_out . --go-grpc_opt paths=source_relative \
  --openapiv2_out . \
  --openapiv2_opt logtostderr=true \
  proto/echopb/echo_server.proto

protoc -I . \
    --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
  proto/echopb/echo_server.proto
