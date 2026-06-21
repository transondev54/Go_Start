# Project 2: CLI Tool with Concurrency

## 🎯 Mục tiêu

Xây dựng một **command-line tool** với:

- ✅ Subcommands support
- ✅ Concurrent file processing
- ✅ Progress tracking
- ✅ Error recovery
- ✅ Performance monitoring
- ✅ Pretty output formatting

---

## 📋 Yêu cầu

### Core Features

1. **Subcommands**
   - `analyze` - Analyze files in directory
   - `convert` - Convert files between formats
   - `batch` - Process multiple files concurrently
   - `report` - Generate analysis report

2. **File Processing**
   - Process multiple files concurrently
   - Progress bar showing completion
   - Error handling per file
   - Graceful degradation

3. **Performance**
   - Worker pool for concurrent processing
   - Configurable number of workers
   - Memory-efficient streaming
   - Performance statistics

4. **Output Formatting**
   - Pretty-printed results
   - Progress indicators
   - Color-coded messages
   - Exportable reports (JSON, CSV)

5. **Error Handling**
   - Individual file error handling
   - Retry logic
   - Skip on error vs fail fast options

---

## 🏗️ Architecture

```
cli/
├── main.go              # Entry point
├── cmd.go               # Command definitions
├── processor.go         # File processing logic
├── worker_pool.go       # Concurrent processing
├── progress.go          # Progress tracking
├── output.go            # Output formatting
└── config.go            # Configuration
```

---

## 📝 Implementation Steps

### Step 1: Define Commands Structure

```go
type Command interface {
    Execute(ctx context.Context, args []string) error
    Usage() string
}

type AnalyzeCmd struct {
    directory string
    workers   int
    format    string
}

type ConvertCmd struct {
    input     string
    output    string
    format    string
}
```

### Step 2: Implement File Processor

```go
type FileProcessor struct {
    input      string
    workerPool *WorkerPool
    progress   *Progress
}

func (fp *FileProcessor) Process(ctx context.Context) error {
    files, err := fp.getFiles()
    if err != nil {
        return err
    }

    results := make(chan ProcessResult, len(files))

    for _, file := range files {
        fp.workerPool.Submit(func() {
            result := fp.processFile(ctx, file)
            results <- result
        })
    }

    // Collect results
    for i := 0; i < len(files); i++ {
        result := <-results
        fp.progress.Update(result)
    }

    return nil
}
```

### Step 3: Create Worker Pool

```go
type WorkerPool struct {
    workers int
    jobs    chan func()
    wg      sync.WaitGroup
}

func (wp *WorkerPool) Submit(job func()) {
    wp.jobs <- job
}

func (wp *WorkerPool) Start(ctx context.Context) {
    for i := 0; i < wp.workers; i++ {
        wp.wg.Add(1)
        go func() {
            defer wp.wg.Done()
            for {
                select {
                case job := <-wp.jobs:
                    if job == nil {
                        return
                    }
                    job()
                case <-ctx.Done():
                    return
                }
            }
        }()
    }
}

func (wp *WorkerPool) Wait() {
    close(wp.jobs)
    wp.wg.Wait()
}
```

### Step 4: Add Progress Tracking

```go
type Progress struct {
    total     int64
    completed int64
    mu        sync.Mutex
}

func (p *Progress) Update(result ProcessResult) {
    p.mu.Lock()
    defer p.mu.Unlock()

    p.completed++
    percentage := (p.completed * 100) / p.total

    fmt.Printf("\r[%-50s] %d%% (%d/%d)",
        strings.Repeat("=", int(percentage/2)),
        percentage,
        p.completed,
        p.total,
    )
}
```

### Step 5: Implement Output Formatting

```go
type OutputFormatter interface {
    Format(results []ProcessResult) string
}

type JSONFormatter struct{}

func (jf *JSONFormatter) Format(results []ProcessResult) string {
    data, _ := json.MarshalIndent(results, "", "  ")
    return string(data)
}

type TableFormatter struct{}

func (tf *TableFormatter) Format(results []ProcessResult) string {
    var buf strings.Builder

    fmt.Fprintf(&buf, "%-30s | %-15s | %-10s\n", "File", "Status", "Duration")
    fmt.Fprintf(&buf, "%s\n", strings.Repeat("-", 60))

    for _, r := range results {
        fmt.Fprintf(&buf, "%-30s | %-15s | %10.2fms\n",
            r.File, r.Status, r.Duration.Milliseconds())
    }

    return buf.String()
}
```

### Step 6: Main Command Handler

```go
func main() {
    if len(os.Args) < 2 {
        printUsage()
        return
    }

    cmd := os.Args[1]
    args := os.Args[2:]

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var command Command

    switch cmd {
    case "analyze":
        command = &AnalyzeCmd{}
    case "convert":
        command = &ConvertCmd{}
    case "batch":
        command = &BatchCmd{}
    default:
        fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
        os.Exit(1)
    }

    if err := command.Execute(ctx, args); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
```

---

## ✅ Test Cases

```bash
# 1. Analyze directory
./cli analyze -dir ./test_files -workers 4 -format json

# 2. Convert files
./cli convert -input input.txt -output output.csv

# 3. Batch processing
./cli batch -dir ./large_dataset -pattern "*.txt" -action process

# 4. Generate report
./cli report -input results.json -output report.html

# 5. Help
./cli -h
./cli analyze -h

# 6. Concurrency test
./cli analyze -dir ./large_dir -workers 8
```

---

## 📊 Performance Targets

- **File processing**: > 100 files/sec per worker
- **Memory efficiency**: < 500MB for large datasets
- **CPU utilization**: < 80%
- **Progress update**: Smooth updates (1-2 times/sec)

---

## 🌟 Bonus Features

- [ ] Configuration file support
- [ ] Resume from checkpoint
- [ ] Parallel directory processing
- [ ] Dry-run mode
- [ ] Verbose/debug logging
- [ ] Parallel upload/download
- [ ] Dependency resolver
- [ ] Watch mode (auto-reprocess on change)
- [ ] Interactive mode (TUI)
- [ ] Plugin system

---

## 📚 Resources

- [Cobra CLI Framework](https://github.com/spf13/cobra)
- [Flags Package](https://pkg.go.dev/flag)
- [Concurrent Patterns](https://go.dev/blog/pipelines)

---

## 💡 Learning Goals

✨ Build production-grade CLI applications
✨ Implement worker pools and concurrency patterns
✨ Handle large-scale data processing
✨ Progress tracking and reporting
✨ Error recovery and resilience
✨ Performance monitoring and optimization
