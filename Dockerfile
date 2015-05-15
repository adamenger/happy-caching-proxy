FROM golang:1.4.2

ADD . /code
WORKDIR /code
RUN go build hcp.go

CMD ./hcp
