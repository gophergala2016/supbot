FROM golang:1.5.3

RUN apt-get update
RUN apt-get install -y ca-certificates git bzr mercurial make

RUN mkdir -p /src/supbot
RUN git clone https://github.com/gophergala2016/supbot.git /src/supbot
RUN cd /src/supbot && go get -d -v ./... && make build
RUN cp /src/supbot/bin/supbot /bin/supbot

RUN go get -u github.com/pressly/sup/cmd/sup
RUN go build -o /bin/sup github.com/pressly/sup/cmd/sup

VOLUME /root/.ssh

RUN mkdir -p /var/supbot
WORKDIR /var/supbot

ENTRYPOINT ["/bin/supbot"]
