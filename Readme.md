# SCRAPE_PROXY
Request proxy service with custom parsers

## Install scrape_proxy on remote virtual host 


### Copy ssh key to remote server

exec next command on local machine
```bash
$ cat ~/.ssh/id_rsa.pub
```

copy result into the file on the remote server
```bash
$ mkdir .ssh

$ vim ./.ssh/authorized_keys
```

### Install golang
```bash
$ wget https://go.dev/dl/go1.19.5.linux-amd64.tar.gz

$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.5.linux-amd64.tar.gz

$ export PATH=$PATH:/usr/local/go/bin

$ go version
```

#### Export path for golang bin directory

Add to the end of file: ```.bashrc``` this line:
```
export PATH=$PATH:/usr/local/go/bin:~/go/bin
export SCRAPE_PROXY_ADDR=:7000
```

### Install scrape_proxy
```bash
$ go install github.com/grizmar-realty/scrape_proxy/cmd/scrape_proxy@latest
```

#### Run scrape_proxy
```bash
$ export SCRAPE_PROXY_ADDR=:7000
$ nohup /root/go/bin/scrape_proxy > scrape_proxy.log
```