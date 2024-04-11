# smartctl_exporter

[![Docker Image Size](https://badgen.net/docker/size/imagelist/smartctl_exporter?icon=docker&label=image%20size)](https://hub.docker.com/r/imagelist/smartctl_exporter)

Prometheus exporter for [smartmontools](https://www.smartmontools.org/) to export the S.M.A.R.T. attributes.

## Deployment

```sh
docker run --detach --privileged -p 9111:9111 imagelist/smartctl_exporter:latest
```

Metrics will be available at http://localhost:9111/metrics

## Options

|option|environment|default|description|
|------|-----------|:-----:|-----------|
|server-addr|SERVER_ADDR||Listen address. If not set listen `0.0.0.0`|
|server-port|SERVER_PORT|9111|Listen port|
|server-user|SERVER_USER|admin|Username|
|server-pass|SERVER_PASS|admin|Password|
|disable-auth|DISABLE_AUTH|false|Disable basic authentication|
|server-tls|SERVER_TLS|true|Enable TLS|
|version|VERSION|false|Show version and exit|

## Grafana Dashboard

![](https://user-images.githubusercontent.com/8357481/164513823-175e1d32-2ba1-41a8-a7f4-76b1b5f07f09.png)
