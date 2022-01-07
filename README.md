# 1. 목표

## grpc-example
Golang으로 gRPC Server 를 구현해본다 [O]

## proto-gateway
proto-gateway를 통해 Restapi와 연동하기 [O]

1. not found google/api/annotaions.proto 에러 메시지 해결 방법
```
:> git clone https://github.com/googleapis/googleapis.git
:> googleapis/api/annotaions.proto  => protos/google/api/annotaions.proto
:> googleapis/api/http.proto => protos/google/api/http.proto
```

2. not found google/protobuf/descriptor.proto 에러 메시지 해결 방법 
  1) local dir 확인
```
:> cd /usr/local/include/google
   해당 파일의 권한 변경 해주기
:> chmod -R 755 /usr/local/include/google

```

  2) /usr/local/include/google 이 없을 경우 
```
:> git clone https://github.com/protocolbuffers/protobuf.git  
:> cp -rp protobuf/src/google /usr/local/include/google
```

## proto-middleware
proto-middleware를 통해 Req/Res Log 남기기 [X]


## protos
protos/v1/user [None Gateway]

protos/v2/user [Use Gateway]
