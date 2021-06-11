## Install 

### MacOS

```bash
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Development

### Generate spec

```bash
protoc -I . *.proto --go_out=plugins=grpc:.
```
