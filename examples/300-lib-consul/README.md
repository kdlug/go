# Consul library

Consul is a tool for service discovery and key value store for configurations.

## Run consul service

```bash
docker-compose up -d
```

Open UI in your browser

```browser
http://localhost:8500
```

## Get consul go library

```go
go get github.com/hashicorp/consul/api
```

Import consul library:

```go
import (
    consulapi "github.com/hashicorp/consul/api"
)
```

## Simple usage via Consul KV CLI

### Go to consul container

```bash
sudo docker exec -it consul ash
```

### Add config

```consul
consul kv put foo bar
```

### Get config value

```consul
consul kv get foo
// bar
```

### Update value

```consul
consul kv put foo baz
```

### Delete key

```consul
consul kv delete foo
```

### Add key with path

```consul
consul kv put service/config/timeout 100
consul kv put service/config/max_connections 1000
consul kv put service/config/username "john"
```

### Get all configs

```consul
consul kv get -recurse
```

### Delete recursive

```consul
consul kv delete -recurse service
```

### Export (json)

```consul
consul kv export
```

### Import (json)

```consul
consul kv import
```

## Resources

- [Consul API in GO](http://techblog.zeomega.com/devops/golang/2015/06/09/consul-kv-api-in-golang.html)
- [Consul Quick GUide](https://www.tutorialspoint.com/consul/consul_quick_guide.htm)
