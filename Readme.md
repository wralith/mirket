# Local Setup

Build docker images
```sh
docker build -t mirket-api-gateway:alpha ./src/userservice
docker build -t userservice:alpha ./src/api-gateway
```

I used `kind` for local development.
```sh
# Creates new kind cluster
kind create cluster
# Don't forget to change context to kind cluster

# Load docker images to kind
kind load docker-image mirket-api-gateway:alpha
kind load docker-image userservice:alpha

# Apply manifests
kubectl apply -f deploy/api-gateway.yaml
kubectl apply -f deploy/user-service.yaml

kubectl get pods
# There must be a pod running for each service
```

Api-Gateway should be exposed to Port 31000