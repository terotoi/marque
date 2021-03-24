
.PHONY: all prod devjs servedev serve clean_ui clean install watch

all: prod

servedev: marque
	./marque serve

marque: *.go */*.go
	go build

marque_prod: *.go

public/dist/ui.js: ui/*.js
	npx webpack ${JSFLAGS}

node_modules:
	yarn install

prod: JSFLAGS = --mode=production
prod: node_modules public/dist/ui.js marque

devjs: JSFLAGS = --mode=development -w
devjs: node_modules clean_ui public/dist/ui.js

clean_ui:
	rm -f ./public/dist/*.js ./public/fonts/*

clean: clean_ui 
	rm -rf ./node_modules ./marque 

watch:
	find . -iname '*.go' | entr -r make servedev

docker: prod
	docker build -t marque .

docker_run: docker
	docker stop marque || true
	docker rm marque || true
	docker run \
		-it \
	  --mount source=marque,target=/data \
		-p 127.0.0.1:9000:9999 \
		--name marque marque:latest

docker_launch: docker
	docker stop marque || true
	docker rm marque || true
	docker run -d --restart=always \
	  --mount source=marque,target=/data \
		-p 127.0.0.1:9000:9999 \
		--name marque marque:latest

