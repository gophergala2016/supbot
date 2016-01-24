FROM ubuntu:14.04

RUN apt-get update
RUN apt-get install -y ca-certificates

COPY supbot_linux_amd64 /bin/supbot

RUN mkdir -p /var/supbot
WORKDIR /var/supbot

ENTRYPOINT ["/bin/supbot"]
