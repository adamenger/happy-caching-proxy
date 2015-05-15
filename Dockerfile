FROM golang:1.4.2

ADD . /code
WORKDIR /code
RUN go get github.com/elazarl/goproxyy
RUN go build hcp.go

CMD ./hcp
