
.PHONY: all prod prodjs watchdev \
	docker_demo docker_demo_db demo demo_gz

all: public/dist/ui.js marque

prod: prodjs marque

marque: *.go
	go build -o marque

rundev: all
	./marque serve

watchdev:
	find . \( \
		-path './ui/*.js' -or \
		-path './ui/*.jsx' -or \
		-path './ui/*/*.jsx' -or \
		-path './*.go' \) | \
		entr -r make rundev

public/dist/ui.js: node_modules ui/*.js ui/*.jsx ui/*/*.jsx
	npm run build

node_modules: package.json
	npm install

clean:
	rm -rf \
		public/dist/* \
		node_modules \
		marque

### Docker images ###


docker: prod
	etc/create_docker_config.sh
	docker build -t marque .

docker_azure: prod
	etc/create_docker_config.sh
	docker build --build-arg azure=true -t marque .

docker_run: docker
	docker stop marque || true
	docker rm marque || true
	docker run \
		-it \
	  --mount source=marque,target=/data \
		-p 127.0.0.1:9998:9999 \
		--name marque marque:latest

docker_launch: docker
	docker stop marque || true
	docker rm marque || true
	docker run -d --restart=always \
	  --mount source=marque,target=/data \
		-p 127.0.0.1:9998:9999 \
		--name marque marque:latest

