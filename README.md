#### Document

##### 1.Download Gen Tool
```shell script
    go get -u github.com/golang/protobuf/protoc-gen-go
```
##### 2.Build Go Package
```shell script
    protoc --go_out=plugins=grpc,paths=source_relative:. ./user/user.proto
```