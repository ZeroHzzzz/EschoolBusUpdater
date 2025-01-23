# 易校园校车后端

## 项目介绍

一个基于易校园 API 的校车信息获取和缓存服务。

## 系统架构

系统处理两类信息：

-   校车信息：包含线路和时间安排，使用配置的 `uid` 定时获取并存入 Redis
-   用户信息：根据前端在 `Header` 中 `Authorization` 字段传入的 `uid` 实时获取

### 数据流程

1. 校车数据更新：

    ```
    易校园 API -> Updater Service -> Redis 缓存
    ```

2. 用户请求处理：
    ```
    用户请求 -> 中间件验证 -> Controller -> Service -> Redis/远程 API
    ```

## 配置说明

配置文件示例 (`config-example.yaml`):

```yaml
eBus:
    host: https://api.pinbayun.com
    uid: ""

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

## 快速开始

运行服务：

```bash
go run main.go
```

## 展望

后续计划将改项目并入 yxy-go
