# Simple Cluster Test

## Test Environment
- Ubuntu v22.04.3 LTS
- Go v1.21.6
- Kind v0.21.0

## Test Process
### Docker Image Build
- ```cd docker```
- ```docker build -t test-app:local .```

### Create Cluster & Deploy App
- ```kind create cluster --name test --config cluster_config.yaml```
- ```kind load docker-image test-app:local -n test```
- (optional)```kubectl config use-context kind-test```
    - This is necessary only if there is another cluster
- ```kubectl apply -f deploy_test_app.yaml```

### Test Request
- ```curl http://localhost:9999/traffic_test```
    - It performs a CPU task and then responds

### Deploy Metric Server & Set HPA
- ```kubectl apply -f deploy_metrics_server.yaml```
- ```kubectl autoscale deployment test-app --min=1 --max=5 --cpu-percent=50```
- ```kubectl get hpa```

### Run Test Script
- ```./request.sh 5```
    - It sends a request every 5 seconds

### Collect Data Metrics
- ```./collect.sh```
    - It collects data metrics every 1 second