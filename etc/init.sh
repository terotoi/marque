#!/bin/bash

# Azure Files storage requires journal_mode=wal

if [ ! -z "IS_AZURE" ]; then
    echo "== Running on azure"
    if [ ! -f "/data/marque.db" ]; then
        echo "== Setting up intial database"
        sqlite3 ./marque.db 'pragma journal_mode=wal;'
        mv ./marque.db /data/
    else
        echo "== Database already exists"
    fi
else
    echo "== Not running on azure"
fi

/dist/marque -c /dist/config.json serve

