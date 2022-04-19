## Install

### MacOS

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install storj.io/drpc/cmd/protoc-gen-go-drpc@latest
```

## Development

### Generate spec

```bash
protoc --go_out=pb --go_opt=paths=source_relative \
--go-drpc_out=pb --go-drpc_opt=paths=source_relative \
example.proto
```


