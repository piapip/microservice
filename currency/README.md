# Lesson 14: gRPC Client Connections

# Lesson 13: gRPC

In a proto file, we will define services and methods for services and messages for methods.
<!-- protoc -I protos/ --go-grpc_out=protos/currency protos/currency.proto -->
protoc -I protos/ --go-grpc_out=protos/currency --go_out=protos/currency protos/currency.proto
If I remove the "-I protos/", it still works, but instead of generating file to protoS/currency, it'll create another protos folder in protoS/currency and generate the files there...

Test server: 
grpcurl --plaintext localhost:9092 list
grpcurl --plaintext localhost:9092 list Currency
grpcurl --plaintext localhost:9092 describe Currency.GetRate
grpcurl --plaintext localhost:9092 describe .RateRequest
<!-- For some reason Idk, that's how I need to write in the cmd in Windows... -->
grpcurl --plaintext -d '{\"Base\": \"GBP\", \"Destination\": \"USD\"}' localhost:9092 Currency.GetRate
