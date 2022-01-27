### 编译协议
 protoc --go_out=plugins=grpc:../ *.proto

 编译 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
 构建 docker build -t servera .
 运行docker docker run -p 80:80 servera

 镜像加载到集群 kind load docker-image servera --name kind
 部署dployment kubectl apply -f deployment.yaml
 流量转发
 转发到pod kubectl port-forward b-deployment-5c845465f9-4w4xv 30080:80  前面是宿主机端口，后面是容器端口
 转发到service kubectl port-forward service/b 30080:80


 plan
 构建镜像，执行镜像             finish
 service                     finish
 日志解决策略,打印到宿主机       finish
 k8s项目管理最佳实践
 deployment 配置可以简化吗      没必要
 服务发现                      finish
 deployment 回滚              目前结论，通过镜像自己管理历史版本以及回滚

 监控：
 普罗米修斯
 接口监控
 单秒请求次数
 单秒请求平均耗时
 成功率（单秒）


 获取当前pod的IP地址，通过环境变量的形式

 最佳实践
 公共配置
 1.角色对象,角色绑定对象
 2.网关(ingress)
 应用配置
 1.deployment
 2.service

 发布流程
 编译二进制包
 打包镜像
 推送镜像
 启动 deployment

 todo grpc的完整k8s部署
 尝试 treafik router 部署方式


kubectl port-forward etcd-kind-control-plane 30080:2379 -n kube-system