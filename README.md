
# Marque - a bookmark manager #

Marque is a bookmark manager which can be used through a web browser.

Requirements: go 1.14, node.js 1.10+, yarn or npm, optionally make.

## Installation ##

`go get github.com/terotoi/marque`

Change to the marque directory and type:

`make`

or:

`yarn install`    (or `npm install`)

`npx webpack --mode=production`

`go build`

The resulting binary is **marque** at it is a standalone executable.
All data files are embedded in the binary.

The binary can be copied to somewhere in your path.

## Running ##

`marque serve`

This starts the server, listening on 127.0.0.1 port 9999 by default.
The server can be accessed by a browser, for example:

`firefox http://localhost:9999/`

An example service script for systemd can be found in **docs/marque.service**

## Configuration ##

The configuration file is in **$HOME/.config/marque/config.json**

Sqlite3 database file for the server is **$HOME/.config/marque/marque.db**

## Development ##

Marque uses pkger to bundle static files inside the binary. If you make
modifications to any files in public/ you have to rebuild the pkged.go using pkger
before building the production executable.

`go get github.com/markbates/pkger/cmd/pkger`

`pkger`

## TODO ##

- Importing and exporting of firefox/chrome bookmarks
- Some layout fixes for mobile
- Authentication and support for multiple user accounts (partially done)
- Prepared statements
- Better parsing of page titles

