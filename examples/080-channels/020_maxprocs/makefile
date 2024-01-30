KIND_CLUSTER := dev
VERSION := 1.0

all: docker-build

docker-build:
	docker build \
		-f k8s/Dockerfile \
		-t app:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

kind-up:
	kind create cluster \
		--image kindest/node:v1.24.12 \
		--name $(KIND_CLUSTER)
	kubectl config set-context --current --namespace=app

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	kind load docker-image app:$(VERSION) --name $(KIND_CLUSTER)

kind-apply:
	kubectl apply -f k8s/app.yaml

kind-update-apply: all kind-load kind-apply

kind-logs:
	kubectl logs -l app=app --all-containers=true -f --tail=100 --namespace=app