
# Marque - a bookmark manager #

Marque is a bookmark manager which can be used through a web browser.

Requirements: go 1.5, node.js 1.10+, yarn or npm and optionally make.

## Installation ##

`git clone https://github.com/terotoi/marque.git`

Change to the marque directory and type:

`make`

or:

`yarn install`    (or `npm install`)

`npx webpack --mode=production`

`go build`

The resulting binary is **marque**.

## Running marque inside a docker image ##

To build the image, type:

`docker build -t marque .` or `make docker`

You can run the image with, for example:

`docker run --mount source=marque,target=/data -p 127.0.0.1:9000:9999 --name marque marque:latest`

This will start the container and redirect the port 9000 on localhost to marque running inside the container.

To launch the image in the background and have it restart automatically, type:

`docker run -d --restart=always --mount source=marque,target=/data -p 127.0.0.1:9000:9999 --name marque marque:latest`

Or just:

`make docker_launch`

## Using marque ##

Use your favorite browser to navigate to http://localhost:9000/

## TODO ##

- Authentication and support for multiple user accounts
- Prepared statements
- Better parsing of some page titles
