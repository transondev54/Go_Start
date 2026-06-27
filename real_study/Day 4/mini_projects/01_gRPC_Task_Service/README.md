# Project 1: gRPC Task Service

## 🎯 Mục tiêu

Xây dựng một **gRPC microservice** quản lý tasks với:

- ✅ CRUD operations qua gRPC (Unary RPCs)
- ✅ Server-streaming: stream danh sách tasks theo filter
- ✅ Interceptors: logging, recovery, auth
- ✅ Proper error handling với gRPC status codes
- ✅ gRPC client để test service
- ✅ Unit tests cho server implementation

---

## 📋 Yêu cầu

### Protobuf Service Definition

```protobuf
syntax = "proto3";
package taskpb;

service TaskService {
    rpc CreateTask (CreateTaskRequest) returns (Task);
    rpc GetTask    (GetTaskRequest)    returns (Task);
    rpc UpdateTask (UpdateTaskRequest) returns (Task);
    rpc DeleteTask (DeleteTaskRequest) returns (DeleteTaskResponse);
    rpc ListTasks  (ListTasksRequest)  returns (stream Task);  // server-streaming
    rpc WatchTasks (WatchTasksRequest) returns (stream TaskEvent); // bonus: live updates
}
```

### Data Models

```go
type Task struct {
    ID          string
    Title       string
    Description string
    Status      TaskStatus // PENDING, IN_PROGRESS, DONE, CANCELLED
    Priority    int32      // 1 (low) to 5 (high)
    DueAt       time.Time
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Features cần implement

1. **CRUD RPCs**
   - `CreateTask`: validate input, generate ID, lưu task
   - `GetTask`: tìm theo ID, return `NotFound` nếu không có
   - `UpdateTask`: partial update, validate trạng thái chuyển đổi
   - `DeleteTask`: xoá, return số lượng deleted

2. **Server-streaming**
   - `ListTasks`: stream tasks có thể filter theo status và priority
   - Respect context cancellation trong streaming loop

3. **Interceptors** (áp dụng cho tất cả RPCs)
   - **Logging**: log method, duration, error
   - **Recovery**: catch panic, trả về `Internal` error
   - **Validation**: validate request fields không nil/empty

4. **Error handling**
   - Dùng đúng gRPC status codes
   - Thêm error details (optional)

5. **Client**
   - CLI client gọi đến service
   - Demo tất cả 5 RPC methods

---

## 🏗️ Architecture

```
grpc_task_service/
├── proto/
│   └── task.proto
├── pb/                    # Generated code (protoc output)
│   ├── task.pb.go
│   └── task_grpc.pb.go
├── server/
│   ├── server.go          # TaskServiceServer implementation
│   └── server_test.go
├── interceptors/
│   ├── logging.go
│   ├── recovery.go
│   └── validation.go
├── client/
│   └── main.go            # CLI client
├── go.mod
└── main.go                # Server entry point
```

---

## 🚀 Implementation Steps

### Step 1: Setup module & proto

```bash
mkdir grpc_task_service ; cd grpc_task_service
go mod init grpc_task_service
go get google.golang.org/grpc@latest
go get google.golang.org/protobuf@latest

# Tạo task.proto và generate:
mkdir proto pb
# Viết task.proto...
protoc --go_out=pb --go-grpc_out=pb --proto_path=proto proto/task.proto
```

### Step 2: Implement server

- Tạo `TaskStore` (in-memory: `sync.RWMutex` + `map[string]*Task`)
- Implement tất cả RPCs
- Handle context cancellation trong `ListTasks`

### Step 3: Add interceptors

- Logging interceptor: record start time, log sau khi handler trả về
- Recovery interceptor: defer + recover() → `status.Error(codes.Internal, ...)`
- Chain với `grpc.ChainUnaryInterceptor`

### Step 4: Client

```bash
# Chạy server
go run main.go

# Chạy client (terminal khác)
go run client/main.go
```

### Step 5: Tests

```go
// server/server_test.go
func TestCreateTask(t *testing.T) {
    srv := newTaskServer()

    resp, err := srv.CreateTask(context.Background(), &pb.CreateTaskRequest{
        Title:    "Learn gRPC",
        Priority: 3,
    })

    if err != nil || resp.Id == "" {
        t.Fatalf("CreateTask failed: %v", err)
    }
}
```

---

## 📊 Expected Output

```
$ go run main.go
INFO server listening on :50051

$ go run client/main.go
Created task: {ID: "abc123", Title: "Learn gRPC", Status: PENDING}
Got task: {ID: "abc123", Title: "Learn gRPC", Status: PENDING}
Updated task: {ID: "abc123", Status: IN_PROGRESS}

Streaming all tasks:
  → task-1: Learn gRPC [IN_PROGRESS]
  → task-2: Build service [PENDING]
  → task-3: Write tests [DONE]

Deleted: 1 task(s)
```

---

## 🌟 Bonus Features

- [ ] Bi-directional streaming: `SyncTasks` - client gửi tasks, server xử lý & trả kết quả
- [ ] TLS với self-signed certificate
- [ ] Health check service (grpc health protocol)
- [ ] gRPC reflection (dùng với `grpcurl`)
- [ ] Metadata propagation (request-id)

---

## ✅ Tiêu chí hoàn thành

- [ ] Tất cả 5 RPC methods hoạt động đúng
- [ ] Interceptors được áp dụng và log ra stdout
- [ ] Error codes đúng (`NotFound`, `InvalidArgument`, ...)
- [ ] Server-streaming respect context cancellation
- [ ] `go test ./...` pass
