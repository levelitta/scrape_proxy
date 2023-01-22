# SCRAPE_PROXY
Request proxy service with custom parsers

## Install scrape_proxy on remote virtual host 

### Install golang
```bash
$ wget https://go.dev/dl/go1.19.5.linux-amd64.tar.gz

$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.5.linux-amd64.tar.gz

$ go version
```

#### Export path for golang bin directory

Add to the end of file: ```.bashrc``` this line
```
export PATH=$PATH:~/go/bin
```

### Install scrape_proxy
```bash
$ go install github.com/grizmar-realty/scrape_proxy/cmd/scrape_proxy@latest
```

#### Run scrape_proxy
```bash
$ SCRAPE_PROXY_ADDR=:7000 /root/go/bin/scrape_proxy
```