package main

import (
	"flag"
	"fmt"
	"github.com/elazarl/goproxy"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	listen          = flag.String("listen", "8080", "Port to listen on")
	cache_directory = flag.String("dir", "cache", "directory to cache into")
	verbose         = flag.Bool("verbose", false, "enable verbose mode")
)

// Here's where we do the heavy lifting of the file downloading
func fileDownloader(file string, url string) {

	filename := get_filename(file)

	out, err := os.Create(fmt.Sprintf("%s/%s", *cache_directory, filename))
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Downloaded: %s (%d bytes)", filename, n)
	}
}

// Check to see if cache directory exists, if not, create it
func initialize_cache(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Printf("Cache directory does not exist.\n")
		log.Printf("Creating cache directory: %s \n", directory)
		os.Mkdir(directory, 0700)
		return
	} else {
		gem_count, _ := ioutil.ReadDir(*cache_directory)
		log.Printf("Cache directory initialized: %s \n", directory)
		log.Printf("%d packages in the cache \n", len(gem_count))
	}
}

// return just the filename of the gem
func get_filename(path string) string {
	sl := strings.Split(path, "/")
	filename := sl[len(sl)-1]
	return filename
}

// TODO there has to be a better way to do this
//resorting to this because the stdlib insists on returning ports on SSL
func generate_url(scheme string, host string, path string) string {
	if scheme == "https" {
		host, _, err := net.SplitHostPort(host)
		if err != nil {
			log.Fatal("Fatal error generating package url!")
		}
		return fmt.Sprintf("%s://%s%s", scheme, host, path)
	} else {
		return fmt.Sprintf("%s://%s%s", scheme, host, path)
	}
}

// check our cache, if not, call the downloader
// returns file object which the proxy will serve back
func check_or_cache(gem string, url string) *os.File {
	gem_file := get_filename(gem)
	full_gem_path := fmt.Sprintf("%s/%s", *cache_directory, gem_file)

	if _, err := os.Stat(full_gem_path); os.IsNotExist(err) {
		log.Printf("Cache MISS for package: %s\n", full_gem_path)
		log.Printf("Downloading: %s\n", gem)
		fileDownloader(gem, url)

		gem_file, _ := os.Open(full_gem_path)
		return gem_file
	} else {
		log.Printf("Cache HIT for package: %s\n", full_gem_path)
		gem_file, _ := os.Open(full_gem_path)
		return gem_file
	}
}

// This is a wrapper around the std http.Response, here we pump the os.File Writer into the body of the response
func fileResponse(r *http.Request, contentType string, status int, body *os.File) *http.Response {
	resp := &http.Response{}
	resp.Request = r
	resp.TransferEncoding = r.TransferEncoding
	resp.Header = make(http.Header)
	resp.Header.Add("Content-Type", contentType)
	resp.StatusCode = status
	resp.Body = ioutil.NopCloser(body)
	return resp
}

// where the magic happens
func main() {

	// Parse all cli flags
	flag.Parse()

	// initialize the cache
	initialize_cache(*cache_directory)

	// set up the proxy
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.Verbose = *verbose

	// Currently only match gem files
	r, err := regexp.Compile("^.*\\.(gem|deb)$")
	if err != nil {
		log.Fatal("There is a problem with your regex.\n")
		return
	}

	// add a conditional to only serve requests that match *.gem
	proxy.OnRequest(goproxy.UrlMatches(r)).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			gem_url := generate_url(r.URL.Scheme, r.URL.Host, r.URL.Path)
			gem_file := check_or_cache(r.URL.Path, gem_url)
			return r, fileResponse(r, "application/octet", 200, gem_file)
		})

	// To protect and serve!
	log.Printf("Happy Caching Proxy listening on: localhost:%s \n", *listen)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *listen), proxy))
}
