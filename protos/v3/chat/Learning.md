google/api.http.proto에 기록된  example을 확인 해 본다. 
Sub Learning Objectives
===================
proto v3 에서는 아래의 사항을 확인 한다.
1. Golang에서는 interface{} Type을  return 할수 있다. 
   gRPC에서 interface{} 타입으로 return 되는 Golang Handler를 매핑 시킬 수 있는지 확인 한다.

2. HTTP의 특정 Header 정보가 필요 할 경우 이를 protobuf에 정의하는 방법을 확인 한다.
   gRPC Service 구현 시 Header 정보를 Gettering 하는 방법에 대해서 확인 한다.

3. gRPC Stream 방식 사용 시 gRPC Gateway에서 어떻게 동작하는지 확인 한다.

4. GUI `->` API-SERVER `->` ServiceApps 구현 시 API-SERVER에서 ServiceApps에 구현된 gRPC를 바로 Routing 하는 방법을 확인 한다.

5. Protobuf 정의 시 http rules 를 사용 하는 방법을 알아 본다.
Garbage Text
===================
1. HTTP Header 정보를 protobuf에 정의 하는 방법
2. HTTP Header 정보와 rpc 서비스간 매핑되는 정보 
 > grpc에서는 metadata or ctx에서 확인 할 수 있다?? 
3. gRPC Stream 방식이 gRPC Gateway 사용 할 경우 동일하게 사용 할 수 있는지
4. gRPC에서 하나의 URL에서 여러개의 Body 타입을 가질 수 있는 방법
5. Message를 바로 Routing 하는 방법
