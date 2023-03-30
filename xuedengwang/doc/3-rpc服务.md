

##### 1、注册rpc服务

这里以usercenter服务举例
- 定义protobuf文件

  在service/usercenter/rpc/pb中新建usercenter.proto，写入注册方法

```protobuf
  //req 、resp

message CaptchaReq {
  int64  height = 1;
  int64  width = 2;
  int64  length = 3;
  float  maxSkew = 4;
  int64  dotCount = 5;
}

message CaptchaResp {
  string uuid = 1;
  string img = 2;
  int64 captchaLength = 3;
}

message VerifyCaptchaReq {
  string uuid = 1;
  string answer = 2;
  bool clear = 3;
}

message VerifyCaptchaResp {
  bool success = 1;
}



message GenerateTokenReq {
  int64 userId = 1;
}

message GenerateTokenResp {
  string accessToken = 1;
  int64  accessExpire = 2;
  int64  refreshAfter = 3;
}

//service
service usercenter {
  rpc generateCaptcha(CaptchaReq) returns(CaptchaResp);
  rpc verifyCaptcha(VerifyCaptchaReq) returns(VerifyCaptchaResp);
  rpc generateToken(GenerateTokenReq) returns(GenerateTokenResp);
}
```

  

- 使用goctl生成代码

  1）命令行进入service/usercenter/rpc/pb目录下。

  2）在命令行中执行（命令行要切换到service/usercenter/rpc/pb目录）

  ```shell
    goctl rpc protoc user.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero
  ```

  ```shell
    sed -i "" 's/,omitempty//g'  ./rpc/pb/*.pb.go
  ```
【注】建议在生成rpc文件时候，把protobuf生成的omitempty给删除掉，不然字段为nil就不返回了











