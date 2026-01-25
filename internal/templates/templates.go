package templates

import (
	"embed"
	"fmt"
)

//go:embed embedded/*
var embeddedFS embed.FS

// Template represents a project template
type Template struct {
	Name        string
	Description string
	Directories []string
	Files       []FileTemplate
}

// FileTemplate represents a file to be generated
type FileTemplate struct {
	Path     string
	Template string
	Content  string
}

var builtInTemplates = map[string]Template{
	"go-api": {
		Name:        "go-api",
		Description: "Go REST API with clean architecture",
		Directories: []string{
			"cmd/api",
			"internal/handler",
			"internal/service",
			"internal/repository",
			"internal/model",
			"pkg/config",
		},
		Files: []FileTemplate{
			{Path: "cmd/api/main.go", Content: goAPIMainTmpl},
			{Path: "internal/handler/handler.go", Content: goHandlerTmpl},
			{Path: "internal/service/service.go", Content: goServiceTmpl},
			{Path: "internal/repository/repository.go", Content: goRepositoryTmpl},
			{Path: "internal/model/model.go", Content: goModelTmpl},
			{Path: "pkg/config/config.go", Content: goConfigTmpl},
			{Path: "README.md", Content: readmeTmpl},
			{Path: ".gitignore", Content: gitignoreGoTmpl},
		},
	},
	"go-cli": {
		Name:        "go-cli",
		Description: "Go CLI application with Cobra",
		Directories: []string{
			"cmd",
			"internal",
			"pkg",
		},
		Files: []FileTemplate{
			{Path: "main.go", Content: goCLIMainTmpl},
			{Path: "cmd/root.go", Content: goCLIRootTmpl},
			{Path: "README.md", Content: readmeTmpl},
			{Path: ".gitignore", Content: gitignoreGoTmpl},
		},
	},
	"go-lib": {
		Name:        "go-lib",
		Description: "Go library/package",
		Directories: []string{
			"internal",
			"examples",
		},
		Files: []FileTemplate{
			{Path: "{{.ProjectName}}.go", Content: goLibMainTmpl},
			{Path: "{{.ProjectName}}_test.go", Content: goLibTestTmpl},
			{Path: "examples/main.go", Content: goLibExampleTmpl},
			{Path: "README.md", Content: readmeTmpl},
			{Path: ".gitignore", Content: gitignoreGoTmpl},
		},
	},
	// Learning templates
	"learn-concurrency": {
		Name:        "learn-concurrency",
		Description: "Learn Go concurrency patterns",
		Directories: []string{
			"01-goroutines",
			"02-channels",
			"03-select",
			"04-sync",
			"05-patterns",
		},
		Files: []FileTemplate{
			{Path: "README.md", Content: learnConcurrencyReadmeTmpl},
			{Path: "01-goroutines/main.go", Content: learnGoroutinesTmpl},
			{Path: "02-channels/main.go", Content: learnChannelsTmpl},
			{Path: "03-select/main.go", Content: learnSelectTmpl},
			{Path: "04-sync/main.go", Content: learnSyncTmpl},
			{Path: "05-patterns/main.go", Content: learnPatternsTmpl},
			{Path: ".gitignore", Content: gitignoreGoTmpl},
		},
	},
	"learn-testing": {
		Name:        "learn-testing",
		Description: "Learn Go testing techniques",
		Directories: []string{
			"unit",
			"table",
			"mock",
			"benchmark",
		},
		Files: []FileTemplate{
			{Path: "README.md", Content: learnTestingReadmeTmpl},
			{Path: "unit/calculator.go", Content: learnUnitCalcTmpl},
			{Path: "unit/calculator_test.go", Content: learnUnitTestTmpl},
			{Path: "table/validator.go", Content: learnTableValidatorTmpl},
			{Path: "table/validator_test.go", Content: learnTableTestTmpl},
			{Path: "benchmark/string_builder.go", Content: learnBenchFuncTmpl},
			{Path: "benchmark/string_builder_test.go", Content: learnBenchTestTmpl},
			{Path: ".gitignore", Content: gitignoreGoTmpl},
		},
	},
	"learn-dsa": {
		Name:        "learn-dsa",
		Description: "Practice Data Structures & Algorithms",
		Directories: []string{
			"datastructures/stack",
			"datastructures/queue",
			"datastructures/linkedlist",
			"datastructures/tree",
			"algorithms/sorting",
			"algorithms/searching",
			"algorithms/recursion",
		},
		Files: []FileTemplate{
			{Path: "README.md", Content: learnDSAReadmeTmpl},
			// Data Structures
			{Path: "datastructures/stack/stack.go", Content: dsaStackTmpl},
			{Path: "datastructures/stack/stack_test.go", Content: dsaStackTestTmpl},
			{Path: "datastructures/queue/queue.go", Content: dsaQueueTmpl},
			{Path: "datastructures/queue/queue_test.go", Content: dsaQueueTestTmpl},
			{Path: "datastructures/linkedlist/linkedlist.go", Content: dsaLinkedListTmpl},
			{Path: "datastructures/linkedlist/linkedlist_test.go", Content: dsaLinkedListTestTmpl},
			// Algorithms
			{Path: "algorithms/sorting/sorting.go", Content: dsaSortingTmpl},
			{Path: "algorithms/sorting/sorting_test.go", Content: dsaSortingTestTmpl},
			{Path: "algorithms/searching/searching.go", Content: dsaSearchingTmpl},
			{Path: "algorithms/searching/searching_test.go", Content: dsaSearchingTestTmpl},
			{Path: "algorithms/recursion/recursion.go", Content: dsaRecursionTmpl},
			{Path: "algorithms/recursion/recursion_test.go", Content: dsaRecursionTestTmpl},
			{Path: ".gitignore", Content: gitignoreGoTmpl},
		},
	},
	"go-grpc": {
		Name:        "go-grpc",
		Description: "Go gRPC service template",
		Directories: []string{
			"cmd/server",
			"cmd/client",
			"internal/service",
			"proto",
		},
		Files: []FileTemplate{
			{Path: "cmd/server/main.go", Content: goGRPCServerTmpl},
			{Path: "cmd/client/main.go", Content: goGRPCClientTmpl},
			{Path: "proto/service.proto", Content: goGRPCProtoTmpl},
			{Path: "Makefile", Content: goGRPCMakefileTmpl},
			{Path: "README.md", Content: readmeTmpl},
			{Path: ".gitignore", Content: gitignoreGoTmpl},
		},
	},
	"go-worker": {
		Name:        "go-worker",
		Description: "Background worker with job queue",
		Directories: []string{
			"cmd/worker",
			"internal/job",
			"internal/queue",
		},
		Files: []FileTemplate{
			{Path: "cmd/worker/main.go", Content: goWorkerMainTmpl},
			{Path: "internal/job/job.go", Content: goWorkerJobTmpl},
			{Path: "internal/queue/queue.go", Content: goWorkerQueueTmpl},
			{Path: "README.md", Content: readmeTmpl},
			{Path: ".gitignore", Content: gitignoreGoTmpl},
		},
	},
	"go-tui": {
		Name:        "go-tui",
		Description: "Terminal UI app with Bubbletea",
		Directories: []string{
			"internal/ui",
		},
		Files: []FileTemplate{
			{Path: "main.go", Content: goTUIMainTmpl},
			{Path: "internal/ui/model.go", Content: goTUIModelTmpl},
			{Path: "README.md", Content: readmeTmpl},
			{Path: ".gitignore", Content: gitignoreGoTmpl},
		},
	},
	"fullstack": {
		Name:        "fullstack",
		Description: "Go backend + React/Vite/Bun/Tailwind frontend",
		Directories: []string{
			"backend/cmd/api",
			"backend/internal/handler",
			"backend/internal/middleware",
			"frontend/src/components",
			"frontend/src/hooks",
			"frontend/src/lib",
		},
		Files: []FileTemplate{
			// Backend files
			{Path: "backend/cmd/api/main.go", Content: fullstackBackendMainTmpl},
			{Path: "backend/internal/handler/handler.go", Content: fullstackHandlerTmpl},
			{Path: "backend/internal/middleware/cors.go", Content: fullstackCorsTmpl},
			{Path: "backend/go.mod", Content: fullstackBackendGoModTmpl},
			// Frontend files
			{Path: "frontend/package.json", Content: fullstackPackageJsonTmpl},
			{Path: "frontend/tsconfig.json", Content: fullstackTsconfigTmpl},
			{Path: "frontend/tsconfig.node.json", Content: fullstackTsconfigNodeTmpl},
			{Path: "frontend/vite.config.ts", Content: fullstackViteConfigTmpl},
			{Path: "frontend/tailwind.config.js", Content: fullstackTailwindConfigTmpl},
			{Path: "frontend/postcss.config.js", Content: fullstackPostcssConfigTmpl},
			{Path: "frontend/index.html", Content: fullstackIndexHtmlTmpl},
			{Path: "frontend/src/main.tsx", Content: fullstackMainTsxTmpl},
			{Path: "frontend/src/App.tsx", Content: fullstackAppTsxTmpl},
			{Path: "frontend/src/index.css", Content: fullstackIndexCssTmpl},
			{Path: "frontend/src/lib/api.ts", Content: fullstackApiTsTmpl},
			{Path: "frontend/src/components/Header.tsx", Content: fullstackHeaderTmpl},
			// Root files
			{Path: "README.md", Content: fullstackReadmeTmpl},
			{Path: "Makefile", Content: fullstackMakefileTmpl},
			{Path: ".gitignore", Content: fullstackGitignoreTmpl},
		},
	},
}

// GetTemplate returns a template by name
func GetTemplate(name string) (Template, error) {
	t, ok := builtInTemplates[name]
	if !ok {
		return Template{}, fmt.Errorf("template not found: %s", name)
	}
	return t, nil
}

// GetAllTemplates returns all available templates
func GetAllTemplates() []Template {
	result := make([]Template, 0, len(builtInTemplates))
	for _, t := range builtInTemplates {
		result = append(result, t)
	}
	return result
}

// Template content definitions

var goAPIMainTmpl = `package main

import (
	"fmt"
	"log"
	"net/http"

	"{{.ModuleName}}/internal/handler"
	"{{.ModuleName}}/pkg/config"
)

func main() {
	cfg := config.Load()

	h := handler.New()

	http.HandleFunc("/", h.Health)
	http.HandleFunc("/api/v1/", h.HandleAPI)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
`

var goHandlerTmpl = `package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) HandleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello from {{.ProjectName}}"})
}
`

var goServiceTmpl = `package service

// Service handles business logic
type Service struct{}

// New creates a new Service instance
func New() *Service {
	return &Service{}
}
`

var goRepositoryTmpl = `package repository

// Repository handles data persistence
type Repository struct{}

// New creates a new Repository instance
func New() *Repository {
	return &Repository{}
}
`

var goModelTmpl = `package model

// Define your data models here
`

var goConfigTmpl = `package config

import "os"

type Config struct {
	Port string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{Port: port}
}
`

var goCLIMainTmpl = `package main

import "{{.ModuleName}}/cmd"

func main() {
	cmd.Execute()
}
`

var goCLIRootTmpl = `package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "{{.ProjectName}}",
	Short: "{{.ProjectName}} - A brief description",
	Long:  "{{.ProjectName}} - A longer description of your CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to {{.ProjectName}}!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`

var goLibMainTmpl = `// Package {{.ProjectName}} provides ...
package {{.PackageName}}

// Hello returns a greeting message
func Hello() string {
	return "Hello from {{.ProjectName}}!"
}
`

var goLibTestTmpl = `package {{.PackageName}}

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello from {{.ProjectName}}!"
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
`

var goLibExampleTmpl = `package main

import (
	"fmt"
	"{{.ModuleName}}"
)

func main() {
	fmt.Println({{.PackageName}}.Hello())
}
`

var readmeTmpl = `# {{.ProjectName}}

{{.Description}}

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

` + "```bash" + `
go mod download
` + "```" + `

### Running

` + "```bash" + `
go run ./cmd/...
` + "```" + `

## License

{{.License}}
`

var gitignoreGoTmpl = `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output
/bin/
/dist/

# Dependency directories
/vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db

# Environment
.env
.env.local
`

// ============== Learning Templates ==============

var learnConcurrencyReadmeTmpl = `# Learn Go Concurrency

This project contains examples to learn Go concurrency patterns.

## Sections

1. **01-goroutines** - Basic goroutine usage
2. **02-channels** - Channel communication
3. **03-select** - Select statement for multiplexing
4. **04-sync** - Synchronization primitives (Mutex, WaitGroup)
5. **05-patterns** - Common concurrency patterns

## How to Run

` + "```bash" + `
cd 01-goroutines && go run main.go
` + "```" + `
`

var learnGoroutinesTmpl = `package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic goroutine
	go sayHello("World")

	// Anonymous goroutine
	go func() {
		fmt.Println("Hello from anonymous goroutine!")
	}()

	// Multiple goroutines
	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Printf("Goroutine %d\n", n)
		}(i)
	}

	// Wait for goroutines (simple approach - use sync.WaitGroup in production)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done!")
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
`

var learnChannelsTmpl = `package main

import "fmt"

func main() {
	// Unbuffered channel
	ch := make(chan string)
	go func() {
		ch <- "Hello from channel!"
	}()
	msg := <-ch
	fmt.Println(msg)

	// Buffered channel
	buffered := make(chan int, 3)
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println(<-buffered, <-buffered, <-buffered)

	// Channel with range
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	for n := range numbers {
		fmt.Printf("Received: %d\n", n)
	}

	// Channel direction (send-only, receive-only)
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	pingPong(ping, pong)
}

func pingPong(ping chan<- string, pong <-chan string) {
	ping <- "ping"
	// This demonstrates send-only and receive-only channels
}
`

var learnSelectTmpl = `package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	// Select receives from whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received", msg2)
		}
	}

	// Select with timeout
	timeout := make(chan string)
	select {
	case msg := <-timeout:
		fmt.Println(msg)
	case <-time.After(10 * time.Millisecond):
		fmt.Println("Timeout!")
	}

	// Non-blocking select with default
	nonBlocking := make(chan string)
	select {
	case msg := <-nonBlocking:
		fmt.Println(msg)
	default:
		fmt.Println("No message available")
	}
}
`

var learnSyncTmpl = `package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// WaitGroup - wait for goroutines to finish
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Worker %d done\n", n)
		}(i)
	}
	wg.Wait()
	fmt.Println("All workers done!")

	// Mutex - protect shared data
	var mu sync.Mutex
	counter := 0
	var wg2 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Printf("Counter: %d\n", counter)

	// Atomic operations
	var atomicCounter int64
	var wg3 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}
	wg3.Wait()
	fmt.Printf("Atomic counter: %d\n", atomicCounter)

	// Once - do something only once
	var once sync.Once
	for i := 0; i < 3; i++ {
		once.Do(func() {
			fmt.Println("This prints only once!")
		})
	}
}
`

var learnPatternsTmpl = `package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Worker Pool Pattern
	fmt.Println("=== Worker Pool ===")
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		fmt.Printf("Result: %d\n", <-results)
	}

	// Fan-out, Fan-in Pattern
	fmt.Println("\n=== Fan-out, Fan-in ===")
	input := generate(1, 2, 3, 4, 5)
	c1 := square(input)
	c2 := square(input)
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(10 * time.Millisecond)
		results <- j * 2
	}
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
`

// ============== Learn Testing Templates ==============

var learnTestingReadmeTmpl = `# Learn Go Testing

This project contains examples for Go testing techniques.

## Sections

1. **unit/** - Basic unit tests
2. **table/** - Table-driven tests
3. **benchmark/** - Benchmark tests

## How to Run

` + "```bash" + `
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run benchmarks
go test -bench=. ./benchmark/
` + "```" + `
`

var learnUnitCalcTmpl = `package unit

import "fmt"

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

var ErrDivideByZero = fmt.Errorf("cannot divide by zero")
`

var learnUnitTestTmpl = `package unit

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(5, 3)
	want := 2
	if got != want {
		t.Errorf("Subtract(5, 3) = %d; want %d", got, want)
	}
}

func TestDivide(t *testing.T) {
	t.Run("valid division", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 5 {
			t.Errorf("Divide(10, 2) = %d; want 5", got)
		}
	})

	t.Run("divide by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Error("expected error for divide by zero")
		}
	})
}
`

var learnTableValidatorTmpl = `package table

import "regexp"

// IsValidEmail checks if the email format is valid
func IsValidEmail(email string) bool {
	pattern := ` + "`" + `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$` + "`" + `
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// IsValidAge checks if age is within valid range
func IsValidAge(age int) bool {
	return age >= 0 && age <= 150
}
`

var learnTableTestTmpl = `package table

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"valid email", "user@example.com", true},
		{"valid with subdomain", "user@mail.example.com", true},
		{"missing @", "userexample.com", false},
		{"missing domain", "user@", false},
		{"empty string", "", false},
		{"with plus", "user+tag@example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidEmail(tt.email)
			if got != tt.want {
				t.Errorf("IsValidEmail(%q) = %v; want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestIsValidAge(t *testing.T) {
	tests := []struct {
		age  int
		want bool
	}{
		{0, true},
		{25, true},
		{150, true},
		{-1, false},
		{151, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := IsValidAge(tt.age)
			if got != tt.want {
				t.Errorf("IsValidAge(%d) = %v; want %v", tt.age, got, tt.want)
			}
		})
	}
}
`

var learnBenchFuncTmpl = `package benchmark

import "strings"

// ConcatPlus concatenates strings using +
func ConcatPlus(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}

// ConcatBuilder concatenates strings using strings.Builder
func ConcatBuilder(strs []string) string {
	var builder strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}
`

var learnBenchTestTmpl = `package benchmark

import "testing"

var testStrings = []string{"hello", "world", "this", "is", "a", "test", "of", "string", "concatenation"}

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatPlus(testStrings)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatBuilder(testStrings)
	}
}
`

// ============== gRPC Templates ==============

var goGRPCServerTmpl = `package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// TODO: Register your gRPC service here
	// pb.RegisterYourServiceServer(s, &server{})

	log.Printf("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
`

var goGRPCClientTmpl = `package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// TODO: Create client and make RPC calls
	// c := pb.NewYourServiceClient(conn)
	// resp, err := c.YourMethod(context.Background(), &pb.Request{})

	log.Println("Connected to gRPC server")
}
`

var goGRPCProtoTmpl = `syntax = "proto3";

package {{.PackageName}};

option go_package = "{{.ModuleName}}/proto";

service {{.ProjectName}}Service {
  rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}
`

var goGRPCMakefileTmpl = `.PHONY: proto build run-server run-client

proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

build:
	go build -o bin/server ./cmd/server
	go build -o bin/client ./cmd/client

run-server:
	go run ./cmd/server

run-client:
	go run ./cmd/client
`

// ============== Worker Templates ==============

var goWorkerMainTmpl = `package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"{{.ModuleName}}/internal/job"
	"{{.ModuleName}}/internal/queue"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown gracefully
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Shutting down...")
		cancel()
	}()

	// Create job queue
	q := queue.New(10)

	// Start workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i, q)
	}

	// Add some sample jobs
	for i := 1; i <= 10; i++ {
		q.Enqueue(job.Job{ID: i, Payload: fmt.Sprintf("Task %d", i)})
	}

	// Wait for context cancellation
	<-ctx.Done()
	log.Println("Worker stopped")
}

func worker(ctx context.Context, id int, q *queue.Queue) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			j, ok := q.Dequeue()
			if !ok {
				continue
			}
			log.Printf("Worker %d processing: %s", id, j.Payload)
			j.Process()
		}
	}
}
`

var goWorkerJobTmpl = `package job

import (
	"log"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

func (j *Job) Process() error {
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	log.Printf("Job %d completed: %s", j.ID, j.Payload)
	return nil
}
`

var goWorkerQueueTmpl = `package queue

import (
	"{{.ModuleName}}/internal/job"
)

type Queue struct {
	jobs chan job.Job
}

func New(size int) *Queue {
	return &Queue{
		jobs: make(chan job.Job, size),
	}
}

func (q *Queue) Enqueue(j job.Job) {
	q.jobs <- j
}

func (q *Queue) Dequeue() (job.Job, bool) {
	select {
	case j := <-q.jobs:
		return j, true
	default:
		return job.Job{}, false
	}
}

func (q *Queue) Close() {
	close(q.jobs)
}
`

// ============== TUI Templates ==============

var goTUIMainTmpl = `package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"{{.ModuleName}}/internal/ui"
)

func main() {
	p := tea.NewProgram(ui.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
`

var goTUIModelTmpl = `package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)
)

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func NewModel() Model {
	return Model{
		choices:  []string{"Option 1", "Option 2", "Option 3", "Option 4"},
		selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := titleStyle.Render("{{.ProjectName}}") + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		style := itemStyle
		if m.cursor == i {
			cursor = ">"
			style = selectedStyle
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += style.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, choice)) + "\n"
	}

	s += "\n(q to quit)\n"
	return s
}
`

// ============== Fullstack Templates ==============

var fullstackBackendMainTmpl = `package main

import (
	"log"
	"net/http"

	"{{.ProjectName}}-backend/internal/handler"
	"{{.ProjectName}}-backend/internal/middleware"
)

func main() {
	h := handler.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", h.Health)
	mux.HandleFunc("/api/message", h.GetMessage)

	// Wrap with CORS middleware
	corsHandler := middleware.CORS(mux)

	log.Println("🚀 Backend server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
`

var fullstackHandlerTmpl = `package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Hello from Go backend! 🚀",
		"project": "{{.ProjectName}}",
	})
}
`

var fullstackCorsTmpl = `package middleware

import "net/http"

// CORS middleware to allow frontend to communicate with backend
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
`

var fullstackBackendGoModTmpl = `module {{.ProjectName}}-backend

go 1.21
`

var fullstackPackageJsonTmpl = `{
  "name": "{{.ProjectName}}-frontend",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "tsc && vite build",
    "preview": "vite preview"
  },
  "dependencies": {
    "react": "^18.2.0",
    "react-dom": "^18.2.0"
  },
  "devDependencies": {
    "@types/react": "^18.2.0",
    "@types/react-dom": "^18.2.0",
    "@vitejs/plugin-react": "^4.2.0",
    "autoprefixer": "^10.4.16",
    "postcss": "^8.4.32",
    "tailwindcss": "^3.4.0",
    "typescript": "^5.3.0",
    "vite": "^5.0.0"
  }
}
`

var fullstackTsconfigTmpl = `{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "skipLibCheck": true,
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "react-jsx",
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true
  },
  "include": ["src"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
`

var fullstackTsconfigNodeTmpl = `{
  "compilerOptions": {
    "composite": true,
    "skipLibCheck": true,
    "module": "ESNext",
    "moduleResolution": "bundler",
    "allowSyntheticDefaultImports": true,
    "strict": true
  },
  "include": ["vite.config.ts"]
}
`

var fullstackViteConfigTmpl = `import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
`

var fullstackTailwindConfigTmpl = `/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
`

var fullstackPostcssConfigTmpl = `export default {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
}
`

var fullstackIndexHtmlTmpl = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.ProjectName}}</title>
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="/src/main.tsx"></script>
  </body>
</html>
`

var fullstackMainTsxTmpl = `import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
`

var fullstackAppTsxTmpl = `import { useState, useEffect } from 'react'
import { api } from './lib/api'
import { Header } from './components/Header'

function App() {
  const [message, setMessage] = useState<string>('')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchMessage = async () => {
      try {
        const data = await api.getMessage()
        setMessage(data.message)
      } catch (error) {
        setMessage('Failed to connect to backend')
      } finally {
        setLoading(false)
      }
    }
    fetchMessage()
  }, [])

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
      <Header />
      <main className="container mx-auto px-4 py-16">
        <div className="text-center">
          <h1 className="text-5xl font-bold text-white mb-8">
            Welcome to <span className="text-purple-400">{{.ProjectName}}</span>
          </h1>
          <div className="bg-white/10 backdrop-blur-lg rounded-xl p-8 max-w-md mx-auto">
            <p className="text-gray-300 mb-4">Message from Go Backend:</p>
            {loading ? (
              <div className="animate-pulse bg-purple-500/20 h-8 rounded"></div>
            ) : (
              <p className="text-2xl font-semibold text-purple-400">{message}</p>
            )}
          </div>
          <div className="mt-12 flex justify-center gap-4">
            <a
              href="http://localhost:8080/api/health"
              target="_blank"
              className="px-6 py-3 bg-purple-600 hover:bg-purple-700 text-white rounded-lg transition-colors"
            >
              Check API Health
            </a>
          </div>
        </div>
      </main>
    </div>
  )
}

export default App
`

var fullstackIndexCssTmpl = `@tailwind base;
@tailwind components;
@tailwind utilities;

body {
  font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;
}
`

var fullstackApiTsTmpl = `const API_BASE = '/api'

export const api = {
  async getMessage(): Promise<{ message: string; project: string }> {
    const res = await fetch(` + "`${API_BASE}/message`" + `)
    if (!res.ok) throw new Error('Failed to fetch')
    return res.json()
  },

  async checkHealth(): Promise<{ status: string }> {
    const res = await fetch(` + "`${API_BASE}/health`" + `)
    if (!res.ok) throw new Error('Failed to fetch')
    return res.json()
  },
}
`

var fullstackHeaderTmpl = `export function Header() {
  return (
    <header className="border-b border-white/10">
      <nav className="container mx-auto px-4 py-4 flex justify-between items-center">
        <div className="text-xl font-bold text-white">{{.ProjectName}}</div>
        <div className="flex gap-6 text-gray-300">
          <a href="#" className="hover:text-purple-400 transition-colors">Home</a>
          <a href="#" className="hover:text-purple-400 transition-colors">About</a>
          <a href="#" className="hover:text-purple-400 transition-colors">Contact</a>
        </div>
      </nav>
    </header>
  )
}
`

var fullstackReadmeTmpl = `# {{.ProjectName}}

Fullstack application with Go backend and React frontend.

## Tech Stack

**Backend:**
- Go 1.21+
- Net/HTTP (standard library)

**Frontend:**
- React 18 + TypeScript
- Vite (dev server & build)
- Tailwind CSS
- Bun (package manager)

## Getting Started

### Prerequisites
- Go 1.21+
- Bun (https://bun.sh)

### Installation

` + "```bash" + `
# Install frontend dependencies
cd frontend
bun install
` + "```" + `

### Running

` + "```bash" + `
# Terminal 1: Start backend
make backend

# Terminal 2: Start frontend
make frontend
` + "```" + `

Or use the Makefile:
` + "```bash" + `
make dev  # Runs both in parallel
` + "```" + `

### URLs
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
`

var fullstackMakefileTmpl = `.PHONY: dev backend frontend install

# Run both servers
dev:
	@echo "Starting backend and frontend..."
	@make -j2 backend frontend

# Run Go backend
backend:
	cd backend && go run ./cmd/api

# Run React frontend
frontend:
	cd frontend && bun run dev

# Install frontend dependencies
install:
	cd frontend && bun install

# Build for production
build:
	cd backend && go build -o ../dist/api ./cmd/api
	cd frontend && bun run build
`

var fullstackGitignoreTmpl = `# Dependencies
node_modules/
vendor/

# Build
dist/
build/

# Environment
.env
.env.local

# IDE
.idea/
.vscode/
*.swp

# OS
.DS_Store
Thumbs.db

# Go
*.exe
*.exe~
*.dll
*.so
*.dylib

# Logs
*.log
`

// ============== Learn DSA Templates ==============

var learnDSAReadmeTmpl = `# Data Structures & Algorithms Practice

Practice implementing common data structures and algorithms in Go.

## Structure

` + "```" + `
datastructures/
├── stack/       # Stack implementation
├── queue/       # Queue implementation
├── linkedlist/  # Linked List implementation
└── tree/        # Binary Tree (coming soon)

algorithms/
├── sorting/     # Sorting algorithms
├── searching/   # Searching algorithms
└── recursion/   # Recursive problems
` + "```" + `

## How to Practice

1. Open any ` + "`*_test.go`" + ` file to see the expected behavior
2. Implement the functions in the corresponding ` + "`.go`" + ` file
3. Run tests: ` + "`go test ./...`" + `

## Quick Commands

` + "```bash" + `
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test -v ./datastructures/stack/
go test -v ./algorithms/sorting/
` + "```" + `
`

var dsaStackTmpl = `package stack

// Stack represents a LIFO data structure
type Stack struct {
	items []int
}

// New creates a new empty stack
func New() *Stack {
	return &Stack{items: []int{}}
}

// Push adds an element to the top of the stack
// TODO: Implement this
func (s *Stack) Push(item int) {
	// Your code here
}

// Pop removes and returns the top element
// Returns -1 if stack is empty
// TODO: Implement this
func (s *Stack) Pop() int {
	// Your code here
	return -1
}

// Peek returns the top element without removing it
// Returns -1 if stack is empty
// TODO: Implement this
func (s *Stack) Peek() int {
	// Your code here
	return -1
}

// IsEmpty returns true if the stack has no elements
// TODO: Implement this
func (s *Stack) IsEmpty() bool {
	// Your code here
	return true
}

// Size returns the number of elements in the stack
// TODO: Implement this
func (s *Stack) Size() int {
	// Your code here
	return 0
}
`

var dsaStackTestTmpl = `package stack

import "testing"

func TestPush(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Size() = %d; want 3", s.Size())
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if got := s.Pop(); got != 3 {
		t.Errorf("Pop() = %d; want 3", got)
	}
	if got := s.Pop(); got != 2 {
		t.Errorf("Pop() = %d; want 2", got)
	}
	if s.Size() != 1 {
		t.Errorf("Size() = %d; want 1", s.Size())
	}
}

func TestPopEmpty(t *testing.T) {
	s := New()
	if got := s.Pop(); got != -1 {
		t.Errorf("Pop() on empty stack = %d; want -1", got)
	}
}

func TestPeek(t *testing.T) {
	s := New()
	s.Push(42)
	
	if got := s.Peek(); got != 42 {
		t.Errorf("Peek() = %d; want 42", got)
	}
	// Peek shouldn't remove the element
	if s.Size() != 1 {
		t.Errorf("Size() after Peek() = %d; want 1", s.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Error("IsEmpty() = false; want true for new stack")
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Error("IsEmpty() = true; want false after Push")
	}
}
`

var dsaQueueTmpl = `package queue

// Queue represents a FIFO data structure
type Queue struct {
	items []int
}

// New creates a new empty queue
func New() *Queue {
	return &Queue{items: []int{}}
}

// Enqueue adds an element to the back of the queue
// TODO: Implement this
func (q *Queue) Enqueue(item int) {
	// Your code here
}

// Dequeue removes and returns the front element
// Returns -1 if queue is empty
// TODO: Implement this
func (q *Queue) Dequeue() int {
	// Your code here
	return -1
}

// Front returns the front element without removing it
// Returns -1 if queue is empty
// TODO: Implement this
func (q *Queue) Front() int {
	// Your code here
	return -1
}

// IsEmpty returns true if the queue has no elements
// TODO: Implement this
func (q *Queue) IsEmpty() bool {
	// Your code here
	return true
}

// Size returns the number of elements in the queue
// TODO: Implement this
func (q *Queue) Size() int {
	// Your code here
	return 0
}
`

var dsaQueueTestTmpl = `package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Size() = %d; want 3", q.Size())
	}
}

func TestDequeue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// FIFO: first in, first out
	if got := q.Dequeue(); got != 1 {
		t.Errorf("Dequeue() = %d; want 1", got)
	}
	if got := q.Dequeue(); got != 2 {
		t.Errorf("Dequeue() = %d; want 2", got)
	}
}

func TestFront(t *testing.T) {
	q := New()
	q.Enqueue(42)
	q.Enqueue(100)

	if got := q.Front(); got != 42 {
		t.Errorf("Front() = %d; want 42", got)
	}
	// Front shouldn't remove the element
	if q.Size() != 2 {
		t.Errorf("Size() after Front() = %d; want 2", q.Size())
	}
}
`

var dsaLinkedListTmpl = `package linkedlist

// Node represents a node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	Head *Node
	size int
}

// New creates a new empty linked list
func New() *LinkedList {
	return &LinkedList{}
}

// Append adds a new node at the end
// TODO: Implement this
func (l *LinkedList) Append(value int) {
	// Your code here
}

// Prepend adds a new node at the beginning
// TODO: Implement this
func (l *LinkedList) Prepend(value int) {
	// Your code here
}

// Delete removes the first node with the given value
// Returns true if deleted, false if not found
// TODO: Implement this
func (l *LinkedList) Delete(value int) bool {
	// Your code here
	return false
}

// Find returns true if value exists in the list
// TODO: Implement this
func (l *LinkedList) Find(value int) bool {
	// Your code here
	return false
}

// Size returns the number of nodes
// TODO: Implement this
func (l *LinkedList) Size() int {
	// Your code here
	return 0
}

// ToSlice converts the linked list to a slice (for testing)
func (l *LinkedList) ToSlice() []int {
	var result []int
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}
`

var dsaLinkedListTestTmpl = `package linkedlist

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	got := l.ToSlice()
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v; want %v", got, want)
	}
}

func TestPrepend(t *testing.T) {
	l := New()
	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	got := l.ToSlice()
	want := []int{3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v; want %v", got, want)
	}
}

func TestDelete(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if !l.Delete(2) {
		t.Error("Delete(2) = false; want true")
	}
	got := l.ToSlice()
	want := []int{1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() after delete = %v; want %v", got, want)
	}
}

func TestFind(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if !l.Find(2) {
		t.Error("Find(2) = false; want true")
	}
	if l.Find(99) {
		t.Error("Find(99) = true; want false")
	}
}
`

var dsaSortingTmpl = `package sorting

// BubbleSort sorts a slice using bubble sort algorithm
// Time: O(n²), Space: O(1)
// TODO: Implement this
func BubbleSort(arr []int) []int {
	// Your code here
	return arr
}

// InsertionSort sorts a slice using insertion sort algorithm
// Time: O(n²), Space: O(1)
// TODO: Implement this
func InsertionSort(arr []int) []int {
	// Your code here
	return arr
}

// MergeSort sorts a slice using merge sort algorithm
// Time: O(n log n), Space: O(n)
// TODO: Implement this
func MergeSort(arr []int) []int {
	// Your code here
	return arr
}

// QuickSort sorts a slice using quick sort algorithm
// Time: O(n log n) average, Space: O(log n)
// TODO: Implement this
func QuickSort(arr []int) []int {
	// Your code here
	return arr
}
`

var dsaSortingTestTmpl = `package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{5, 2, 8, 1, 9}, []int{1, 2, 5, 8, 9}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		got := BubbleSort(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BubbleSort(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	got := InsertionSort([]int{5, 2, 8, 1, 9})
	want := []int{1, 2, 5, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertionSort() = %v; want %v", got, want)
	}
}

func TestMergeSort(t *testing.T) {
	got := MergeSort([]int{5, 2, 8, 1, 9})
	want := []int{1, 2, 5, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeSort() = %v; want %v", got, want)
	}
}

func TestQuickSort(t *testing.T) {
	got := QuickSort([]int{5, 2, 8, 1, 9})
	want := []int{1, 2, 5, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("QuickSort() = %v; want %v", got, want)
	}
}
`

var dsaSearchingTmpl = `package searching

// LinearSearch finds the index of target in arr
// Returns -1 if not found
// Time: O(n)
// TODO: Implement this
func LinearSearch(arr []int, target int) int {
	// Your code here
	return -1
}

// BinarySearch finds the index of target in a SORTED arr
// Returns -1 if not found
// Time: O(log n)
// TODO: Implement this
func BinarySearch(arr []int, target int) int {
	// Your code here
	return -1
}

// TwoSum finds two numbers that add up to target
// Returns their indices, or [-1, -1] if not found
// TODO: Implement this
func TwoSum(arr []int, target int) [2]int {
	// Your code here
	return [2]int{-1, -1}
}
`

var dsaSearchingTestTmpl = `package searching

import "testing"

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 5, 3, 9, 2}
	
	if got := LinearSearch(arr, 9); got != 3 {
		t.Errorf("LinearSearch(arr, 9) = %d; want 3", got)
	}
	if got := LinearSearch(arr, 99); got != -1 {
		t.Errorf("LinearSearch(arr, 99) = %d; want -1", got)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 5, 9} // sorted
	
	if got := BinarySearch(arr, 5); got != 3 {
		t.Errorf("BinarySearch(arr, 5) = %d; want 3", got)
	}
	if got := BinarySearch(arr, 1); got != 0 {
		t.Errorf("BinarySearch(arr, 1) = %d; want 0", got)
	}
	if got := BinarySearch(arr, 99); got != -1 {
		t.Errorf("BinarySearch(arr, 99) = %d; want -1", got)
	}
}

func TestTwoSum(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	
	got := TwoSum(arr, 9)
	// 2 + 7 = 9, indices 0 and 1
	if got != [2]int{0, 1} && got != [2]int{1, 0} {
		t.Errorf("TwoSum(arr, 9) = %v; want [0, 1]", got)
	}
}
`

var dsaRecursionTmpl = `package recursion

// Factorial calculates n!
// TODO: Implement recursively
func Factorial(n int) int {
	// Your code here
	return 0
}

// Fibonacci returns the nth Fibonacci number
// TODO: Implement recursively
func Fibonacci(n int) int {
	// Your code here
	return 0
}

// ReverseString reverses a string recursively
// TODO: Implement recursively
func ReverseString(s string) string {
	// Your code here
	return ""
}

// SumArray calculates sum of array recursively
// TODO: Implement recursively
func SumArray(arr []int) int {
	// Your code here
	return 0
}
`

var dsaRecursionTestTmpl = `package recursion

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{10, 3628800},
	}

	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{10, 55},
	}

	for _, tt := range tests {
		if got := Fibonacci(tt.n); got != tt.want {
			t.Errorf("Fibonacci(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	if got := ReverseString("hello"); got != "olleh" {
		t.Errorf("ReverseString(\"hello\") = %q; want \"olleh\"", got)
	}
}

func TestSumArray(t *testing.T) {
	if got := SumArray([]int{1, 2, 3, 4, 5}); got != 15 {
		t.Errorf("SumArray([1,2,3,4,5]) = %d; want 15", got)
	}
}
`
