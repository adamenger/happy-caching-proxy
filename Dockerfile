FROM golang:1.4.2

RUN apt-get update \
    && apt-get install git -y -q

RUN go get github.com/elazarl/goproxy
RUN git clone https://github.com/adamenger/happy-caching-proxy.git /code

WORKDIR /code
RUN go build hcp.go

CMD ./hcp
