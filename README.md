# word-of-wisdom
Yet another test task

The application consists of a server and a client that solve tasks using Proof of Work. The client sends the result to the server, which then verifies it and returns a response.

## Running

### Docker

**Build and run the server:**

```bash
docker build -f Dockerfile.server -t word-of-wisdom-server .
docker run --rm -p 8080:8080 word-of-wisdom-server
```

**Build and run the client:**

```bash
docker build -f Dockerfile.client -t word-of-wisdom-cliet .
docker run --rm word-of-wisdom-client -server-address=host.docker.internal:8080
```

### Locally

**Run the server:**

```bash
go run cmd/server/main.go
```
**Run the client:**

```bash
go run cmd/client/main.go
```

## Parameters

### Client

```bash
./client -server-address="localhost:8080" -concurrency=4 -conn-timeout=10s -solving-timeout=5s
```

### Server

```bash
./server -port="8080" -conn-timeout=10s -pow-difficulty=6 -challenge-length=20 -quotes-file-path="./internal/server/quotes_dump.txt"
```
