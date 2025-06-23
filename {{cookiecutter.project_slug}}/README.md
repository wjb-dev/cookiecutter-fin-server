# {{ cookiecutter.project_name }}

{{ cookiecutter.description }}

## gRPC Server Validation Commands

Use `grpcurl` to validate the running gRPC server.

---

### 🔎 List Available Services

```bash
    grpcurl -plaintext localhost:50051 list
```

---

### ❤️ Health Checks

#### Check All Services

```bash
    grpcurl -plaintext -d '{}' \
    localhost:50051 grpc.health.v1.Health/Check
```

#### Check EchoService Health

```bash
    grpcurl -plaintext -d '{"service":"v1.EchoService"}' \
    localhost:50051 grpc.health.v1.Health/Check
```

#### Check PingService Health

```bash
    grpcurl -plaintext -d '{"service":"v1.PingService"}' \
    localhost:50051 grpc.health.v1.Health/Check
```

---

### 📘 Describe Services

#### EchoService

```bash
    grpcurl -plaintext localhost:50051 describe v1.EchoService
```

#### PingService

```bash
    grpcurl -plaintext localhost:50051 describe v1.PingService
```

---

### 🧪 Test RPCs

#### Test EchoService

```bash
    grpcurl -plaintext -d '{"message":"Hello Echo"}' \
    localhost:50051 v1.EchoService/Echo
```

#### Test PingService

```bash
    grpcurl -plaintext -d '{}' \
    localhost:50051 v1.PingService/Ping
```

