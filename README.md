# Happy Caching Proxy

This is a transparent HTTP proxy which caches ruby gems as they pass through.

Directly inspired by the awesome [Angry Caching Proxy](https://github.com/epeli/angry-caching-proxy) project.

## Usage

```
Usage of ./hcp:
  -dir="cache": directory to cache into
  -listen="8080": Port to listen on
  -verbose=false: enable verbose mode
```

## Running

To get started, clone the repo and build the package.
```
$ git clone https://github.com/adamenger/happy-caching-proxy.git
$ cd happy-caching-proxy
$ go build hcp.go
$ ./hcp.go -dir="gems" -listen=9999 -verbose
2015/05/14 03:16:48 Cache directory does not exist.
2015/05/14 03:16:48 Creating cache directory: gems
2015/05/14 03:16:48 Happy Caching Proxy listening on: localhost:9999
```

## Using

In order to use the proxy you'll have to set the http_proxy variable when running `bundle install`.
```
$ http_proxy=http://localhost:9999 bundle install
```

Now you should see some output generated on HCP:
```
2015/05/14 03:14:28 Cache MISS for gem: cache/aws-sdk-v1-1.60.2.gem
2015/05/14 03:14:28 Downloading: /gems/aws-sdk-v1-1.60.2.gem
2015/05/14 03:14:29 Downloaded: representable-1.6.1.gem (44544 bytes)
2015/05/14 03:14:29 Cache MISS for gem: cache/roar-0.12.0.gem
2015/05/14 03:14:29 Downloading: /gems/roar-0.12.0.gem
2015/05/14 03:14:29 Downloaded: roar-0.12.0.gem (28160 bytes)
2015/05/14 03:14:29 Downloaded: aws-sdk-v1-1.60.2.gem (742912 bytes)
2015/05/14 03:14:30 Cache MISS for gem: cache/aws-sdk-1.60.2.gem
```

## License

Whatever.

That should get you started. Happy caching!
