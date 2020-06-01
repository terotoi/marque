.PHONY: prod devjs servedev serve clean_ui clean install watch

prod: JSFLAGS = --mode=production
prod: node_modules public/dist/ui.js marque

devjs: JSFLAGS = --mode=development -w
devjs: node_modules clean_ui public/dist/ui.js

servedev: PKG = rm -f ./pkged.go
servedev: serve

serve: marque
	./marque serve

marque: *.go */*.go
	${PKG}
	go build

marque_prod: *.go

public/dist/ui.js: ui/*.js
	npx webpack ${JSFLAGS}

node_modules:
	yarn install

clean_ui:
	rm -f ./public/dist/*.js ./public/fonts/*

clean: clean_ui 
	rm -rf ./node_modules ./marque

watch:
	rm -f ./pkged.go
	find . -iname '*.go' | entr -r make servedev
