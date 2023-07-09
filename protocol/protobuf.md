
#protubuf使用

[参考文章](https://blog.csdn.net/zhoupenghui168/article/details/130923516)

[下载](https://github.com/protocolbu%EF%AC%80ers/protobuf/releases)

1. 安装protobuf的go语言插件protoc-gen-go插件
```shell
go install github.com/golang/protobuf/protoc-gen-go@latest
```
2. 相关目录下，把proto 文件编译成go文件
   protoc --go_out=./ *.proto
3. [使用](https://developers.google.com/protocol-buffers/docs/proto3)
