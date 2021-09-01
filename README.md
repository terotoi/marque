
# Marque - a bookmark manager #

Marque is a bookmark manager which can be used through a web browser.

Requirements: go 1.17, node.js 1.10+, yarn or npm and optionally make.

## Installation ##

`git clone https://github.com/terotoi/marque.git`

Change to the marque directory and type:

`make`

The resulting binary is **marque**.

## Local usage ##

Run make with the following command in order to create an initial configuration file:

`marque createconfig`

Edit the configuration file (fill in initial_user and initial_password) and the start the program using:

`./marque serve`

The initial_user and initial_password can be cleared after the first run of the software.

## Running marque inside a docker image ##

To build the image, type:

`docker build -t marque .` or `make docker`

You can run the image with, for example:

`docker run --mount source=marque,target=/data -p 127.0.0.1:9998:9999 --name marque marque:latest`

This will start the container and redirect the port 9998 on localhost to marque running inside the container.

To launch the image in the background and have it restart automatically, type:

`docker run -d --restart=always --mount source=marque,target=/data -p 127.0.0.1:9998:9999 --name marque marque:latest`

Or just:

`make docker_launch`

## Building a docker image for Azure ##

It is possible to build a docker image that is compatible with Azure's storage accounts. Type:

`make docker_azure`

And push the marque:latest image to azure. The image expects a writable storage to be mounted in /data.

## Using marque ##

Use your favorite browser to navigate to http://localhost:9998/


