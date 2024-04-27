# predkube_operator

## Test Environment
- Ubuntu v22.04.3 LTS
- Go v1.21.6
- Kind v0.21.0
- Kubebuilder v3.14.0

## Test Commands
- Assume that the cluster where the target 'test-app' is running has already been created with the command below.
    - kind create cluster --name test --config cluster_config.yaml

- Build & Deploy
    - cd predkube
    - make generate
    - make docker-build IMG=test-image:0.1
    - make install
    - make deploy IMG=test-image:0.1
    - kind load docker-image test-image:0.1 -n test

- Create Custom Resource
    - From the sample yaml file
        - cd predkube
        - kubectl apply -k config/samples/
    - From the python
        - cd predkube_python
        - python3 cr_test.py

- Scaling Custom Resource Check
    - kubectl get scalings

- Replica Count Check
    - kubectl get pods --all-namespaces

## Instructions for use in Analyzer
- Use the contents of predkube_python/cr_test.py as a reference.
```
import cr_scaling

cr_scaling.create_scaling_resource('test-app', 'default', 2)
```