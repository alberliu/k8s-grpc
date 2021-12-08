# k8s-grpc
k8s集成grpc,使用traefik作为网关
### 前置条件
首先保证安装docker和k8s集群，我本地使用kind搭建k8s集群  
macOS使用kind搭建k8s集群  
1.安装kind  
brew install kind  
2.创建k8s集群  
kind create cluster  
### 部署demo服务a和服务b
一.为pod中的default赋予服务发现以及服务变化监听的权限，切换目录到项目的k8s下，执行  
kubectl apply -f cluster_role.yaml  
二.分别部署服务a和服务b  
分别切换到internal/a和internal/b目录下，执行  
./publish.sh  
三.部署traefik,并且使用traefik为服务a和服务b代理流量  
1.创建traefik的自定义资源  
kubectl apply -f traefik_crd.yaml  
2.部署traefik的deployment  
kubectl apply -f traefik_deployment.yaml  
3.创建ingress_router  
kubectl apply -f ingress_router.yaml  
### 现在集群已经创建完毕了，可以测试一下  
将本地端口隐射到traefik的端口上，记得pod名称要改为你本地的traefik名称  
kubectl port-forward traefik-95fcd5b9c-scw57 30080:8000   
执行k8s-grpc/pkg/grpclib/resolver/k8s下的TestClient测试方法，检查代理是否成功  

