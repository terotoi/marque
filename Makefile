
.PHONY: all prod jsdev servedev serve clean_ui clean install watch

all: prod

servedev: marque
	./marque serve

marque: *.go */*.go
	go build

marque_prod: *.go

public/dist/ui.js: ui/*.js
	npx webpack ${JSFLAGS}

node_modules:
	npm install

prod: JSFLAGS = --mode=production
prod: node_modules public/dist/ui.js marque

jsdev: node_modules
	rm -rf public/dist/*.js
	npx webpack --mode=development -w

jsprod: node_modules
	rm -f public/dist/ui.js
	npx webpack --mode=production

clean:
	rm -rf ./public/dist/*.js ./public/fonts/* marque node_modules

watchdev:
	find . -iname '*.go' | entr -r make servedev

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

