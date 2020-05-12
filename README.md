# mkn-grpc-gw

REST API using `gRPC-gateway` example repository

## Preparation for gRPC

### Usage

1. Define your [gRPC](https://grpc.io/docs/) service using protocol buffers

   `your_service.proto`:

   ```protobuf
   syntax = "proto3";
   package example;
   message StringMessage {
     string value = 1;
   }

   service YourService {
     rpc Echo(StringMessage) returns (StringMessage) {}
   }
   ```

2. Add a [`google.api.http`](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L46)
   annotation to your .proto file

   `your_service.proto`:

   ```diff
    syntax = "proto3";
    package example;
   +
   +import "google/api/annotations.proto";
   +
    message StringMessage {
      string value = 1;
    }

    service YourService {
   -  rpc Echo(StringMessage) returns (StringMessage) {}
   +  rpc Echo(StringMessage) returns (StringMessage) {
   +    option (google.api.http) = {
   +      post: "/v1/example/echo"
   +      body: "*"
   +    };
   +  }
    }
   ```

   See [a_bit_of_everything.proto](examples/internal/proto/examplepb/a_bit_of_everything.proto)
   for examples of more annotations you can add to customize gateway behavior
   and generated Swagger output.

   If you do not want to modify the proto file for use with grpc-gateway you can
   alternatively use an external
   [gRPC Service Configuration](https://cloud.google.com/endpoints/docs/grpc/grpc-service-config) file.
   [Check our documentation](https://grpc-ecosystem.github.io/grpc-gateway/docs/grpcapiconfiguration.html)
   for more information.

3. Generate gRPC stub

   The following generates gRPC code for Golang based on `path/to/your_service.proto`:

   ```sh
   protoc -I/usr/local/include -I. \
     -I$GOPATH/src \
     -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --go_out=plugins=grpc:. \
     path/to/your_service.proto
   ```

   It will generate a stub file `path/to/your_service.pb.go`.

4. Implement your service in gRPC as usual

   1. (Optional) Generate gRPC stub in the [other programming languages](https://grpc.io/docs/).

   For example, the following generates gRPC code for Ruby based on `path/to/your_service.proto`:

   ```sh
   protoc -I/usr/local/include -I. \
     -I$GOPATH/src \
     -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --ruby_out=. \
     path/to/your_service.proto

   protoc -I/usr/local/include -I. \
     -I$GOPATH/src \
     -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --plugin=protoc-gen-grpc=grpc_ruby_plugin \
     --grpc-ruby_out=. \
     path/to/your_service.proto
   ```

   2. Add the googleapis-common-protos gem (or your language equivalent) as a dependency to your project.
   3. Implement your gRPC service stubs

5. Generate reverse-proxy using `protoc-gen-grpc-gateway`

   ```sh
   protoc -I/usr/local/include -I. \
     -I$GOPATH/src \
     -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --grpc-gateway_out=logtostderr=true:. \
     path/to/your_service.proto
   ```

   It will generate a reverse proxy `path/to/your_service.pb.gw.go`.

## Request sample
```sh
curl --location --request POST 'http://127.0.0.1:8081/v1/mkn/echo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "mikan": {
        "Name": "Request Best Mikan",
        "Kind": "Good ++ Best",
        "Quality": 200
    }
}'
```
