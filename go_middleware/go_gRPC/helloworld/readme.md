# go_gRPC hello-world steps

### 1. Protocol Buffers3 编译器
<blockquote>https://github.com/protocolbuffers/protobuf/releases</blockquote>
安装完毕后，查看版本
<blockquote>$ protoc --version</blockquote>

### 2. golang grpc插件

#### 2.1 设置下代理(非必须)，视网络情况而定
export GOPROXY="https://goproxy.cn,direct"

### 2.2 grpc插件 安装
go get google.golang.org/grpc

### 3. golang版本proto编译器插件
<blockquote>go get -u github.com/golang/protobuf/{proto,protoc-gen-go}</blockquote>

### 4. 生成(根据protoc参数调整)
<blockquote>
1：protoc -I . --go_out=plugins=grpc:. ./example/data.proto <br />
2：protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./example/data.proto
</blockquote>

### 5. helloworld.proto 内容如下
```protobuf
syntax = "proto3"; //指定proto版本

package proto;

option go_package="./;proto";

//定义请求结构
message HelloRequest{
    string name=1;
}

//定义响应结构
message HelloReply{
    string message=1;
}

//定义Hello服务
service Hello{
    //定义服务中的方法
    rpc SayHello(HelloRequest) returns (HelloReply){}
}
```
### 6. 生成 helloworld.pb.go 文件
protoc -I . --go_out=plugins=grpc:. ./hellworld.proto

### 7. 服务端代码
```go
package main

import (
	"fmt"
	pb "go_gRPC/helloworld/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

const (
	//gRPC服务地址
	Address = "127.0.0.1:5050"
)

//定义一个helloServer并实现约定的接口
type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "hello" + in.Name + "."
	return resp, nil
}

var HelloServer = helloService{}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}

	//实现gRPC Server
	s := grpc.NewServer()
	//注册helloServer为客户端提供服务
	pb.RegisterHelloServer(s, HelloServer) //内部调用了s.RegisterServer()
	fmt.Println("Listen on" + Address)

	s.Serve(listen)
}

```
### 8. 客户端代码
```go
package main

import (
	pb "go_gRPC/helloworld/protocol"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:5050"
)

func main() {
	//连接gRPC服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	//调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Message)
}
```
### 9. 运行测试
```go
服务端运行：
└─(20:41:28 on main ✹ ✚ ✭)──> go run server/main.go                                                                                                                                               1 ↵ ──(五,1126)─┘
Listen on127.0.0.1:5050

客户端运行：
└─(20:41:28 on main ✹ ✚ ✭)──> go run client/main.go                                                                                                                                               1 ↵ ──(五,1126)─┘
hellogRPC.
```