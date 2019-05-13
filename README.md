# DNS-Proxy

A simple ​DNS to DNS-over-TLS proxy.

Forwards a TCP connection on port 8853 to Cloudflare's DNS-over-TLS server on `1.1.1.1:853`

## Building

To build locally, run
```bash
go build
``` 

Or with Docker, run
```bash
docker build --pull -t dns-proxy .
```

## Running

To run the binary, just run
```bash
./dns-proxy
```

This will run on port `8853`

To run with docker, 
``` bash
docker run -p 8853:8853 dns-proxy
```

You can verify that it works with [kdig](https://www.knot-dns.cz/docs/2.6/html/man_kdig.html)
```bash
$ kdig +tcp @0.0.0.0#8853 cameron.ci
;; ->>HEADER<<- opcode: QUERY; status: NOERROR; id: 15880
;; Flags: qr rd ra; QUERY: 1; ANSWER: 4; AUTHORITY: 0; ADDITIONAL: 0

;; QUESTION SECTION:
;; cameron.ci.         		IN	A

;; ANSWER SECTION:
cameron.ci.         	68	IN	A	185.199.111.153
cameron.ci.         	68	IN	A	185.199.108.153
cameron.ci.         	68	IN	A	185.199.109.153
cameron.ci.         	68	IN	A	185.199.110.153

;; Received 92 B
;; Time 2019-05-14 00:22:10 NZST
;; From 0.0.0.0@8853(TCP) in 210.4 ms
```
