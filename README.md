# Happy Caching Proxy

This is a transparent http proxy which caches ruby gems as they pass through. Directly inspired by (Angry Caching Proxy)[https://github.com/epeli/angry-caching-proxy]

## Usage
```
Usage of ./hcp:
  -dir="cache": directory to cache into
  -listen="8080": Port to listen on
  -verbose=false: enable verbose mode
```

## Running
```
$ git clone https://github.com/adamenger/happy-caching-proxy.git
$ cd happy-caching-proxy
$ go build hcp.go
$ ./hcp.go -dir="gems" -listen=9999 -verbose
```

