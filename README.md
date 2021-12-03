# Ozon Marketplace Bot

---
## First step

Add your bot token in .env:

```
TOKEN=<here_your_tocken_from_botfather>
```
---

## Build project

### Local

For local assembly you need to perform

```zsh
$ make deps # Installation of dependencies
$ make build # Build project
```
## Running

### For local development

```zsh
$ docker-compose up -d
```

---

## Services

### Grafana:

- http://localhost:3500
- - login `admin`
- - password `MYPASSWORT`

### Metrics:

Metrics of bot

- http://localhost:9105/metrics

### Status:

Service condition and its information

- http://localhost:8500
- - `/live`- Layed whether the server is running
- - `/ready` - Is it ready to accept requests
- - `/version` - Version and assembly information

### Prometheus:

Prometheus is an open-source systems monitoring and alerting toolkit

- http://localhost:9099

### Jaeger UI

Monitor and troubleshoot transactions in complex distributed systems.

- http://localhost:16685

### Graylog

Graylog is a leading centralized log management solution for capturing, storing, and enabling real-time analysis of terabytes of machine data.

- http://localhost:9500
- - login `admin`
- - password `admin`

### Thanks

- [Evald Smalyakov](https://github.com/evald24)
- [Michael Morgoev](https://github.com/zerospiel)
