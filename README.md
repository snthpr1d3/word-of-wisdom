# word-of-wisdom
Yet another test task

Приложение включает сервер и клиент, которые решают задачи с использованием Proof of Work. Клиент отправляет результат серверу, а сервер проверяет его и возвращает ответ.

## Запуск

### Docker

**Собрать и запустить сервер:**

```bash
docker build -f Dockerfile.server -t word-of-wisdom-server .
docker run --rm -p 8080:8080 word-of-wisdom-server
```

**Собрать и запустить клиент:**

```bash
docker build -f Dockerfile.client -t word-of-wisdom-cliet .
docker run --rm -p 8080:8080 word-of-wisdom-client
```

### Локально

**Запустить сервер:**

```bash
go run cmd/server/main.go
```
**Запустить клиент:**

```bash
go run cmd/server/main.go
```

## Параметры

### Клиент

```bash
./client -server-address="localhost:8080" -concurrency=4 -conn-timeout=10s -solving-timeout=5s
```

### Сервер

```bash
./server -port="8080" -conn-timeout=10s -pow-difficulty=6 -challenge-length=20 -quotes-file-path="./internal/server/quotes_dump.txt"
```
