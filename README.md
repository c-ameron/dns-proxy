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

# Questions

>Imagine this proxy being deployed in an infrastructure. What would be the security
concerns you would raise?

If the proxy is compromised you control all of the DNS requests, you're also trusting Cloudflare

>How would you integrate that solution in a distributed, microservices-oriented and
containerized architecture?

First, I'd fix up the code a lot more and make it a lot more robust (timeouts, proper error handling, retries etc),
as well as some metrics and logging (which are crucial).

I'd probably run this one per VM, for example as a Daemonset on Kubernetes.
I don't think there's any performance improvement running multiple on the same machine as they're still proxying via upstream.
However running one per service per machine would improve resilience if needed.

>What other improvements do you think would be interesting to add to the project?

There's lots. I hadn't done any Go so used this an opportunity to learn and play around with it.

Essentially its a really basic proxy with no form of robustness or configurability.

* Defensive programming - there needs to be more error handling, timeouts being set, retries etc
* Multiple upstream DNS servers
* Configuration files (and via environment variables), logging and metrics. Be a "12 Factor App"
* UDP Support