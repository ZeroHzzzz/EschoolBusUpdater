# 易校园校车后端

# 配置

修改配置文件，如 config-example.yaml 中所示

```yaml
eBus:
    host: https://api.pinbayun.com
    token: "ccdde56f94677246698678908ca17206a3d6c717"

redis:
    host: "127.0.0.1"
    port: 6379
    db: 0
    user: default
    pass: "123456"

timer:
    interval: 1

server:
    port: 8080
```

# 运行

```bash
go run main.go
```
