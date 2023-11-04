# Use an official lightweight Alpine as a parent image
FROM alpine:latest

# Set the version of Traefik you want to install
ARG TRAEFIK_VERSION=latest

# Install Traefik
RUN apk --no-cache add curl \
    && curl -L https://github.com/traefik/traefik/releases/download/v${TRAEFIK_VERSION}/traefik_linux-amd64 -o /usr/local/bin/traefik \
    && chmod +x /usr/local/bin/traefik
    
COPY traefik.yaml /etc/traefik/traefik.yaml
COPY config.yaml /etc/traefik/config.yaml

EXPOSE 80
EXPOSE 8080