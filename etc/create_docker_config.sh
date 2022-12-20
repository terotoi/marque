#!/bin/bash
CONFIG="etc/config.json"

if [ -f "$CONFIG" ]; then
	echo "${CONFIG} already exists."
else
	echo "Generating ${CONFIG}"
	./marque -c "$CONFIG" -d "/data" -l ":9999" -i createconfig
fi

