# 异常中断
if [[ $? -ne 0 ]]; then
    exit 1
fi

version=`date "+%Y%m%d.%H%M%S"`
image_name=$1:$version
image_name_latest=$1:latest

echo $image_name
echo 开始发布服务$1
# 构建镜像
cd internal/$1
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
docker build -t $image_name .

# 加载镜像到kind,如果是生产环境，需要推送到docker仓库
kind load docker-image $image_name --name kind

cd ../..
helm --set server.$1.image=$image_name upgrade k8s-grpc ./chart --reuse-values

echo 成功发布服务$1