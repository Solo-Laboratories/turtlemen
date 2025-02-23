# Protobufs

## Requirements
- `npm install ts-protoc-gen` required to compile to typescript
- `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest` required for compiling into golang
- `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

## Compiling
Be mindful and change the commands below to match the protobuf and location as to where they should be compiled to.

### Golang
From the Go module root folder. For example the below is for Authentication Mock so this will be in `{root}/backend/auth-mock`
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ../../protofuls/authentication/mock-user-pass.proto`

### Typscript
From the Frontend root folder. For example the below is for Authentication Mock so this will be in `{root}/frontend/web`
`protoc --plugin="/home/coder/node_modules/.bin/protoc-gen-ts" --ts_opt=esModuleInterop=true --ts_out="." ../../protofuls/authentication/mock-user-pass.proto`

## Usage
### Golang

### Typescript
