# Simple Cluster Test

## Test Environment
- Ubuntu v22.04.3 LTS
- Go v1.21.6
- Kind v0.21.0
- Docker v25.0.2

## Test Process
### Docker Image Build
- Test App
    - ```cd docker```
    - ```docker build -t test-app:local .```
- Request App
    - ```cd docker_request```
    - Modify request.sh
        - HOST_DOCKER_IP=172.17.0.1
        - Replace this IP with the host IP of your Docker network.
    - ```docker build -t request-app:local .```

### Create Cluster & Deploy App
- Test Cluster
    - ```kind create cluster --name test --config cluster_config.yaml```
    - ```kind load docker-image test-app:local -n test```
    - ```kubectl config use-context kind-test```
    - ```kubectl apply -f deploy_test_app.yaml```
- Request Cluster
    - ```kind create cluster --name remote```
    - ```kind load docker-image request-app:local -n remote```
    - ```kubectl config use-context kind-remote```
    - ```kubectl apply -f deploy_request_app.yaml```

### Deploy Metric Server & Set HPA
- ```kubectl apply -f deploy_metrics_server.yaml```
- ```kubectl autoscale deployment test-app --min=1 --max=5 --cpu-percent=50```
- ```kubectl get hpa```

### Collect Data Metrics
- ```./collect.sh```
    - It collects data metrics every 1 second


## Manual Test
### Test Request
- ```curl http://localhost:9999/traffic_test```
    - It performs a CPU task and then responds

### Run Test Script
- ```./request.sh 5```
    - It sends a request every 5 seconds


## References
- https://kubernetes.io/docs/tutorials/kubernetes-basics/deploy-app/deploy-intro/
- https://kubernetes.io/docs/tasks/run-application/run-stateless-application-deployment/
- https://kubernetes.io/docs/tasks/administer-cluster/manage-resources/
- https://stackoverflow.com/questions/50333650/install-python-package-in-docker-file
- https://kubernetes.io/docs/concepts/services-networking/service/
- https://medium.com/@bravinwasike18/how-to-build-and-deploy-a-flask-application-to-kubernetes-cluster-e52d460fcb14
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
- https://github.com/kubernetes-sigs/metrics-server
- https://www.shellhacks.com/alpine-install-curl/