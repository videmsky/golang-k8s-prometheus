## Summary
This repository includes a simple golang http server instrumented with the Prometheus `prometheus/client_golang` library to expose a custom metric via the `/metrics` endpoint. The k8s resources in `/manifests` create a Custom Resource `PodMonitor` that allows our go client's metric endpoint to be discoverable as a scrape target by our cluster's Prometheus operator. We first create the Prometheus operator in our k8s cluster then we apply the manifests in this repo.

Note: These steps are mostly my scratch and your mileage may vary on running these given your OS, Go, Docker, & local k8s environment.

### Run Go Client Locally

configure dependencies:
* `go mod init gkp`
* `go get -u -v ./...`

running the project:
* `go run .`
* http server listens on `localhost:8080`
* metrics endpoint available `localhost:8080/metrics`

testing the project:
* `go test -v -cover ./...`

building the project binary:
* `go build -o app`

### Tag Commit to Trigger Actions Workflow
create tag:
* `NEWTAG=v1.0.1; git tag $NEWTAG && git push origin $NEWTAG`

delete tag:
* `DELTAG=v1.0.1; git tag -d $DELTAG && git push origin :refs/tags/$DELTAG`

### Dockerize Go Client
build & Tag container:
* `docker build -t gkp .`
* `docker tag gkp user/gkp:latest`

login to Dockerhub:
* `docker login`

push local image to Dockerhub registry
* `docker push user/gkp:latest`

### Install Prometheus Operator to K8s
* https://github.com/prometheus-operator/kube-prometheus
install `kube-prometheus` operator and monitoring components:
* `git clone git@github.com:prometheus-operator/kube-prometheus.git`

remember to cd into the cloned directory ^^^ before running:
* `kubectl apply --server-side -f manifests/setup`

```bash
kubectl wait \
	--for condition=Established \
	--all CustomResourceDefinition \
	--namespace=monitoring
```

* `kubectl apply -f manifests/`

to delete `kube-prometheus`:
* `kubectl delete --ignore-not-found=true -f manifests/ -f manifests/setup`

### Manually Deploy `gkp` to K8s

create namespace:
* `kubectl create namespace gkp`

create deployment:
* `kubectl apply -f ./manifests --namespace=gkp`

delete deployment:
* `kubectl delete -f ./manifests --namespace=gkp`

### Access the UIs
a few local `port-forwards` to view GUIs

gkp application:
* `kubectl --namespace gkp port-forward svc/gkp-service 8080`
* `localhost:8080`
* `localhost:8080/metrics`

prometheus dashboard:
* `kubectl --namespace monitoring port-forward svc/prometheus-k8s 9090`
* `localhost:9090`

grafana dashboard:
* `kubectl --namespace monitoring port-forward svc/grafana 3000`
* `localhost:3000`

### Install ArgoCD
https://argo-cd.readthedocs.io/en/stable/getting_started/

apply manifests:
* `kubectl create namespace argocd`
* `kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml`

port forward
* `kubectl port-forward svc/argocd-server -n argocd 8080:443`

get initial client admin secret
* `kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo`

### Deploy to Armory CDaaS

install `armory-cli` w/ homebrew:
* `brew install armory-io/armory/armory-cli`

check version:
* `armory version`

login:
* `armory login`

install Remote Network Agent(RNA) in K8s cluster:
* `armory agent create`
- NOTE: replace `<my-agent-identifier>` below with string you chose in this step

check that agent is up & running:
* `kubectl get all -n armory-rna`

kickoff deployment:
* `armory deploy start -f https://raw.githubusercontent.com/videmsky/golang-k8s-prometheus/main/deploy/deployment.yml --account <my-agent-identifier>`

check staging namespace:
* `kubectl get pods -n gkp-staging --watch`

after manual approve... check prod namespace:
* `kubectl get pods -n gkp-prod --watch`
