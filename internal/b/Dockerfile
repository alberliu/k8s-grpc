FROM scratch as final
COPY main .
CMD ["/main"]

# 编译 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
# 构建 docker build -t servera .
# 运行docker docker run -p 80:80 servera

# 镜像加载到集群 kind load docker-image servera --name kind
# 部署dployment kubectl apply -f deployment.yaml
# 流量转发
# 转发到pod kubectl port-forward servera-deployment-bcbcc8db9-8rpff 30080:80  前面是宿主机端口，后面是容器端口
# 转发到service kubectl port-forward service/servera 30080:80


# plan
# 构建镜像，执行镜像            finish
# service                    finish
# 日志解决策略,打印到宿主机      finish
# k8s项目管理最佳实践
# deployment 配置可以简化吗
# 服务发现
