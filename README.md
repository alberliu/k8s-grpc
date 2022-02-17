# k8s-grpc
k8s集成grpc,使用traefik作为网关,使用Helm作为Kubernetes的包管理器
### 前置条件
首先保证安装docker和k8s集群，我本地使用kind搭建k8s集群  
macOS使用kind搭建k8s集群  
1.安装kind  
brew install kind  
2.创建k8s集群  
kind create cluster  
### 部署demo服务a和服务b
一.构建镜像  
分别执行./publish.sh a 和 ./publish.sh b构建服务a和b的镜像  
二.部署chart  
helm install k8s-grpc ./chart  
### 现在集群已经创建完毕了，可以测试一下  
将本地端口隐射到traefik的端口上，记得pod名称要改为你本地的traefik名称  
kubectl port-forward traefik-95fcd5b9c-scw57 30080:8000   
执行k8s-grpc/pkg/grpclib/resolver/k8s下的TestClient测试方法，检查代理是否成功