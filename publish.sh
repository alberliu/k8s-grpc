# 异常中断
if [[ $? -ne 0 ]]; then
    exit 1
fi

version=`date "+%Y%m%d.%H%M%S"`
image_name=$1:$version
image_name_latest=$1:latest

echo 开始发布服务$1
# 构建镜像
cd internal/$1
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
docker build -t $image_name -t $image_name_latest .


# 加载镜像到kind
kind load docker-image $image_name --name kind
kind load docker-image $image_name_latest --name kind

# 发布
kubectl set image deployment/$1-deployment $1=$image_name

echo 成功发布服务$1