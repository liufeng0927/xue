### api聚合层开发

#### go-zero网关概念
go-zero架构往大的说主要有2部分组成，一个是api，一个是rpc。
api主要是http对外访问的，rpc主要就是内部业务交互，使用的是protobuf+grpc。

该gateway其实并不是真正意义上的网关，它只是一个聚合服务，实际上go-zero官网的例子是每个服务都拆成了api 和 rpc 两个入口，
当项目服务模块多的时候总不能每个服务解析个域名对应api，这时候api前面就要有一个网关了，这个网关才是真正意义上的网关，比如nginx、kong、apisix，
但该项目为了减少后期部署复杂度，这里就不用nginx做网关了，而是把所有服务的api部分聚合在一起，到时候部署的时候可以只将这个聚合服务对外暴露而不需要再加上一层网关。


#### 开发环境安装

1、安装goctl

```shell
# for Go 1.15 and earlier
GO111MODULE=on go get -u github.com/zeromicro/go-zero/tools/goctl@latest

# for Go 1.16 and later
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

验证是否安装成功

```shell
$ goctl --version
```


2、安装protoc

链接：https://github.com/protocolbuffers/protobuf/releases

验证是否安装成功

```shell
protoc --version
```


3、安装protoc-gen-go

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
```

查看$GOPATH/bin下是否有protoc-gen-go即可


4、安装protoc-gen-go-grpc

```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```


#### 开发

1、api 描述文件
```api
//========================> user v1 <========================
//no need auth and login
@server(
	prefix: v1/sys
	group: sys
)
service gateway {
	@doc "login"
	@handler login
	post /login (LoginReq) returns (LoginResp)
	
	@doc "captcha"
	@handler captcha
	get /captcha returns (CaptchaResp)
}
```

语法介绍请查阅：[https://legacy.go-zero.dev/cn/api-grammar.html](https://legacy.go-zero.dev/cn/api-grammar.html)

修改完api描述文件之后，生成api业务代码，进入"service/gateway/api/desc"目录下，执行下面命令：
```shell
goctl api go -api gateway.api -dir ../  --style=goZero
```


