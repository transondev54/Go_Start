# Lesson 1: gRPC & Protocol Buffers

## 📖 Nội dung bài học

1. gRPC là gì & tại sao dùng gRPC?
2. Protocol Buffers (Protobuf) - định nghĩa service
3. Các loại RPC: Unary, Server Streaming, Client Streaming, Bi-directional
4. Interceptors (middleware cho gRPC)
5. Error handling trong gRPC
6. So sánh gRPC vs REST

---

## 1️⃣ gRPC LÀ GÌ?

**gRPC** (Google Remote Procedure Call) là framework RPC hiệu năng cao, dùng HTTP/2 và Protocol Buffers.

### Tại sao dùng gRPC?

| Tiêu chí       | REST/JSON      | gRPC/Protobuf     |
| -------------- | -------------- | ----------------- |
| Protocol       | HTTP/1.1       | HTTP/2            |
| Serialization  | JSON (text)    | Protobuf (binary) |
| Performance    | Trung bình     | Rất cao           |
| Type safety    | Không          | Có (schema)       |
| Streaming      | Không tự nhiên | Native support    |
| Code gen       | Không          | Có (tự động)      |
| Browser hỗ trợ | Rất tốt        | Cần proxy         |

**Khi nào dùng gRPC:**

- Service-to-service communication (microservices)
- Yêu cầu hiệu năng cao & bandwidth nhỏ
- Cần strongly-typed contracts
- Real-time streaming data

---

## 2️⃣ PROTOCOL BUFFERS

### Định nghĩa `.proto` file

```protobuf
// task.proto
syntax = "proto3";

package task;
option go_package = "./proto;task";

// Service definition
service TaskService {
    rpc CreateTask(CreateTaskRequest) returns (Task);
    rpc GetTask(GetTaskRequest) returns (Task);
    rpc ListTasks(ListTasksRequest) returns (stream Task);        // Server streaming
    rpc UpdateTasks(stream UpdateTaskRequest) returns (Summary); // Client streaming
    rpc SyncTasks(stream Task) returns (stream Task);            // Bi-directional
}

// Message definitions
message Task {
    string id         = 1;
    string title      = 2;
    string status     = 3;
    int64  created_at = 4;
}

message CreateTaskRequest {
    string title = 1;
}

message GetTaskRequest {
    string id = 1;
}

message ListTasksRequest {
    string status = 1; // optional filter
}

message UpdateTaskRequest {
    string id     = 1;
    string status = 2;
}

message Summary {
    int32 updated_count = 1;
}
```

### Generate code từ proto

```bash
# Cài đặt tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate
protoc --go_out=. --go-grpc_out=. task.proto
```

---

## 3️⃣ IMPLEMENT gRPC SERVER

### Setup module

```bash
go mod init grpc_demo
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

### Server implementation

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "sync"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    pb "grpc_demo/proto"
)

// taskServer implements TaskServiceServer interface
type taskServer struct {
    pb.UnimplementedTaskServiceServer
    mu    sync.RWMutex
    tasks map[string]*pb.Task
}

func newTaskServer() *taskServer {
    return &taskServer{
        tasks: make(map[string]*pb.Task),
    }
}

// Unary RPC
func (s *taskServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.Task, error) {
    if req.Title == "" {
        return nil, status.Error(codes.InvalidArgument, "title cannot be empty")
    }

    task := &pb.Task{
        Id:        fmt.Sprintf("task-%d", time.Now().UnixNano()),
        Title:     req.Title,
        Status:    "pending",
        CreatedAt: time.Now().Unix(),
    }

    s.mu.Lock()
    s.tasks[task.Id] = task
    s.mu.Unlock()

    return task, nil
}

// Unary RPC
func (s *taskServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
    s.mu.RLock()
    task, ok := s.tasks[req.Id]
    s.mu.RUnlock()

    if !ok {
        return nil, status.Errorf(codes.NotFound, "task %s not found", req.Id)
    }
    return task, nil
}

// Server-streaming RPC
func (s *taskServer) ListTasks(req *pb.ListTasksRequest, stream pb.TaskService_ListTasksServer) error {
    s.mu.RLock()
    defer s.mu.RUnlock()

    for _, task := range s.tasks {
        // Check context cancellation
        if err := stream.Context().Err(); err != nil {
            return status.FromContextError(err).Err()
        }

        // Filter by status if provided
        if req.Status != "" && task.Status != req.Status {
            continue
        }

        if err := stream.Send(task); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterTaskServiceServer(s, newTaskServer())

    log.Println("gRPC server listening on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

---

## 4️⃣ gRPC CLIENT

```go
package main

import (
    "context"
    "io"
    "log"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "grpc_demo/proto"
)

func main() {
    // Connect to server
    conn, err := grpc.Dial(
        "localhost:50051",
        grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewTaskServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Unary call
    task, err := client.CreateTask(ctx, &pb.CreateTaskRequest{Title: "Learn gRPC"})
    if err != nil {
        log.Fatalf("CreateTask: %v", err)
    }
    log.Printf("Created: %s - %s", task.Id, task.Title)

    // Server-streaming call
    stream, err := client.ListTasks(ctx, &pb.ListTasksRequest{})
    if err != nil {
        log.Fatalf("ListTasks: %v", err)
    }
    for {
        t, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("stream.Recv: %v", err)
        }
        log.Printf("Task: %s - %s", t.Id, t.Title)
    }
}
```

---

## 5️⃣ INTERCEPTORS (MIDDLEWARE)

Interceptors là middleware cho gRPC - tương tự HTTP middleware.

```go
// Logging interceptor (unary)
func loggingInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    start := time.Now()

    resp, err := handler(ctx, req)

    log.Printf("method=%s duration=%s err=%v",
        info.FullMethod,
        time.Since(start),
        err,
    )
    return resp, err
}

// Recovery interceptor (panic recovery)
func recoveryInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (resp interface{}, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = status.Errorf(codes.Internal, "panic: %v", r)
        }
    }()
    return handler(ctx, req)
}

// Áp dụng interceptors khi tạo server
s := grpc.NewServer(
    grpc.ChainUnaryInterceptor(
        loggingInterceptor,
        recoveryInterceptor,
    ),
)
```

---

## 6️⃣ ERROR HANDLING

gRPC dùng status codes riêng (không phải HTTP status codes):

```go
import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// Trả về error với code
return nil, status.Error(codes.NotFound, "item not found")
return nil, status.Errorf(codes.InvalidArgument, "invalid id: %s", id)

// Kiểm tra error code ở client
if st, ok := status.FromError(err); ok {
    switch st.Code() {
    case codes.NotFound:
        fmt.Println("not found:", st.Message())
    case codes.InvalidArgument:
        fmt.Println("bad request:", st.Message())
    default:
        fmt.Println("unknown error:", err)
    }
}
```

### Common gRPC status codes

| Code                | Ý nghĩa                 |
| ------------------- | ----------------------- |
| `OK`                | Thành công              |
| `InvalidArgument`   | Input không hợp lệ      |
| `NotFound`          | Không tìm thấy          |
| `AlreadyExists`     | Đã tồn tại              |
| `PermissionDenied`  | Không có quyền          |
| `Unauthenticated`   | Chưa xác thực           |
| `ResourceExhausted` | Rate limit / quota vượt |
| `Internal`          | Lỗi internal server     |
| `Unavailable`       | Service không available |
| `DeadlineExceeded`  | Timeout                 |

---

## 🧠 QUIZ - 5 CÂU HỎI

1. gRPC dùng protocol nào và tại sao nhanh hơn REST?
2. Sự khác nhau giữa server-streaming và bi-directional streaming RPC là gì?
3. Interceptor trong gRPC tương đương với gì trong HTTP server?
4. Khi nào nên dùng `codes.InvalidArgument` vs `codes.Internal`?
5. `UnimplementedTaskServiceServer` embedding có tác dụng gì?

---

## 📌 KEY TAKEAWAYS

- gRPC dùng HTTP/2 + Protobuf → nhanh hơn REST/JSON đáng kể
- Proto file là single source of truth cho API contract
- 4 loại RPC: unary, server-stream, client-stream, bi-directional
- Interceptors = middleware, dùng cho logging/auth/metrics
- Luôn embed `Unimplemented*Server` để forward-compatible
