# Monitoring

## Perspective

Two dimensions perspective:

1. Service Flow

> Where the request comes and goes?
> Upsteam -> Service -> Downsteam -> Storage

    - QPS(rate)
    - Latency(avg, p99)
    - Spike frequency
    - 

2. Level of Service

> Which layer of service has the issue?
> Business logic -> Framework configure(Gin, KiteX) -> Pod(Host) -> Cluster(Physical Cluster)

## Tools

- Grafana, pure dashboard UI to intergrate different storage(prometheus, OpenTSD, Mysql, ...)
- Metrics, issue analyze tool
- Argos, one stop CD/monitoring
