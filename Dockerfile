FROM debian:stable
ARG azure
ENV IS_AZURE=$azure

# Needed for downloading titles.
RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
	apt-get -y install ca-certificates sqlite3

RUN "mkdir" "/data"

WORKDIR "/dist"

COPY "marque" "."
COPY "public" "./public"
COPY "etc/config_docker.json" "./config.json"
COPY "etc/init.sh" "."

EXPOSE 9999

CMD ["bash", "./init.sh"]
