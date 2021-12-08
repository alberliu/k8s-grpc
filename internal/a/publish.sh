CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
docker build -t a .
kind load docker-image a --name kind
kubectl delete deployment a-deployment
kubectl apply -f k8s.yaml