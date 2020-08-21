FROM debian:stable

ENV cfg="./marque.json"
RUN echo cfg is ${cfg}

# Needed for downloading titles.
RUN mkdir -p "/etc/ssl/certs"
COPY "/etc/ca-certificates.crt" "/etc/ssl/certs"

WORKDIR "/dist"
RUN mkdir "/log"

COPY "marque" "."
COPY "public" "."
COPY "etc/cfg_docker.json" "${cfg}"

EXPOSE 9999

CMD ["sh", "-c", "/dist/marque -c ${cfg} serve"]

