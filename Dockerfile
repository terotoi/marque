FROM debian:stable

# Needed for downloading titles.
RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
	apt-get -y install ca-certificates

WORKDIR "/dist"
COPY "marque" "."
COPY "public" "./public"
COPY "etc/config_docker.json" "./config.json"

EXPOSE 9999

CMD ["sh", "-c", "/dist/marque -c ./config.json serve"]

