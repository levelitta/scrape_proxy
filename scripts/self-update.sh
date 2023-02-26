go install github.com/grizmar-realty/scrape_proxy/cmd/scrape_proxy@latest
kill $(pidof scrape_proxy)
nohup /root/go/bin/scrape_proxy > scrape_proxy.log &