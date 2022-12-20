FROM ubuntu:22.04

# Needed for downloading titles.
RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
	apt-get -y install ca-certificates sqlite3

RUN "mkdir" "/data"

WORKDIR "/dist"

COPY "marque" "."
COPY "public" "./public"
COPY "etc/config.json" "/data/config.json"

EXPOSE 9999

CMD ["bash", "-c", "/dist/marque -c /data/config.json serve"]
