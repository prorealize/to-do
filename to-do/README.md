# To do API

A simple to-do API built with Go language using 
`REST` and `gRPC` for a proof of concept only.

There are multiple edge cases that are not being handled on this API.

## Update the proto file
```bash
apt update && apt install -y protobuf-compiler
protoc --go_out=./api --go_opt=paths=source_relative \
        --go-grpc_out=./api --go-grpc_opt=paths=source_relative ../proto/notification.proto
```


## API Endpoints
- GET /items: Fetch all to-do items
- GET /items/:id: Fetch a single to-do item by its ID
- POST /items: Create a new to-do item
- PUT /items/:id: Update a to-do item by its ID
- DELETE /items/:id: Delete a to-do item by its ID
