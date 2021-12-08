CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
docker build -t b .
kind load docker-image b --name kind
kubectl delete deployment b-deployment
kubectl apply -f k8s.yaml