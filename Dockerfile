FROM golang:1.4.2

ADD . /code
WORKDIR /code
RUN go get github.com/elazarl/goproxy
RUN go build hcp.go

CMD ./hcp
