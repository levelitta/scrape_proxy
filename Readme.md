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
$ apt-get update

$ apt install wget

$ wget https://go.dev/dl/go1.19.5.linux-amd64.tar.gz

$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.5.linux-amd64.tar.gz

$ export PATH=$PATH:/usr/local/go/bin

$ go version
```

#### Export path for golang bin directory

Add to the end of file: ```.bashrc``` this line:
```bash
$ vim ~/.bashrc
```

```
echo 'export PATH=$PATH:/usr/local/go/bin:~/go/bin' >> ~/.bashrc
echo 'export SCRAPE_PROXY_ADDR=:7000' >> ~/.bashrc
echo 'export SCRAPE_PROXY_AUTH_TOKEN={token}' >> ~/.bashrc
```

#### Network config
```bash
$ apt install ufw
$ ufw default deny incoming
$ ufw default allow outgoing
$ ufw allow ssh
$ ufw allow 22
$ ufw allow 7000
$ ufw enable

$ ufw status verbose
```

### Install scrape_proxy
```bash
$ go install github.com/grizmar-realty/scrape_proxy/cmd/scrape_proxy@latest
```

#### Run scrape_proxy
```bash
$ nohup /root/go/bin/scrape_proxy > scrape_proxy.log
```
