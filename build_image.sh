# 异常中断
if [[ $? -ne 0 ]]; then
    exit 1
fi

version=`date "+%Y%m%d.%H%M%S"`
image_name=$1:$version
image_name_latest=$1:latest

echo $image_name
echo 开始发布服务$1

cd internal/$1
# 打包可执行文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
# 构建镜像
docker build -t $image_name -t $image_name_latest .

# 加载镜像到kind,如果是生产环境，需要推送到docker仓库
kind load docker-image $image_name --name kind
kind load docker-image $image_name_latest --name kind

echo $image_name
echo $image_name_latest