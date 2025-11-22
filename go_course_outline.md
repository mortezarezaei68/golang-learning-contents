# Go Programming Course Outline

## Introductions
- Introduction to Go
- Setting Up the Go Environment (Go SDK, VS Code, and GOPATH)
- Understanding the Go Workspace and Project Structure
- How Go Code Runs: The Role of `go run`, `go build`, and `go install`
- Your First Go Program
- Go Syntax Overview: Minimalism and Convention
- Operators and Expressions, Conditional Statements

## Data Types and Functions
- Variables and Constants
- Basic Data Types: `int`, `float`, `string`, `bool`, and `rune`
- Working with Strings and Runes in Go
- Declaring and Calling Functions
- Function Parameters and Return Values (Including Multiple Returns)
- Variable Scope: Global vs Local Variables
- Named Return Values and Naked Returns

## Structs and Interface
- Structs: Defining and Using Custom Data Types
- Pointers in Go: Reference and Dereference Explained
- Value vs Reference Semantics in Go
- What are an Interface, an Empty Interface, and a Nil Interface in Go, and How Does Go Decide Which Method to Call?
- Using Interfaces as Function Parameters
- Polymorphism, Composition Over Inheritance
- When Not to Use Interfaces

## Error Handling & Panic Model
- Introduction to Error Handling in Go
- The Error Type Explained and Returning Errors from Functions
- Creating Custom Errors
- What Is a Panic in Go?
- `defer` + `panic` + `recover`: The Safe Trio
- Implementing the `Unwrap()` Method Manually
- Error Chaining and Context Stacking

## Collections
- Introduction to Collections in Go
- Declaring and Initializing Arrays
- What Is a Slice? Understanding the Backing Array
- Creating and Modifying Slices
- Introduction to Maps — Key-Value Data Structures
- Creating Maps with `make()` and Literals
- Performance Considerations: Memory, Copying, and Growth

## File Handling & Encoding
- Reading and Writing Files with the `os` and `io/ioutil` Packages
- Using `bufio` for Buffered Reading and Writing
- Working with File Permissions and Metadata (`os.FileInfo`)
- Handling CSV and JSON Files — Encoding and Decoding Data
- Using the `encoding/json` Package for Struct Serialization
- Error Handling and Resource Management with `defer file.Close()`
- Practical Project: Reading, Transforming, and Writing Data to a File

## OS Fundamentals for Go
- Overview of Go’s Compilation Process
- The Go Toolchain
- Cross Compilation
- Performance and Thread Monitoring
- Goroutines and the Go Scheduler
- The M:N Model
- Concurrency and Parallelism Demo
- Interacting with OS Processes
- Environment & Signals

## Concurrency Basics
- Understanding Concurrency vs Parallelism in Go
- Introducing Goroutines — Lightweight Threads
- Using Channels for Communication and Synchronization
- Buffered vs Unbuffered Channels
- The Select Statement — Coordinating Multiple Channels
- WaitGroups and Synchronization with the `sync` Package
- Common Concurrency Pitfalls: Deadlocks, Race Conditions, and Blocking

## Packages & Tooling
- Understanding Go Modules and the `go.mod` File
- Importing and Organizing Packages
- Creating and Publishing Your Own Packages
- Dependency Management with `go get` and Versioning
- Testing with `go test` and Writing Table-Driven Tests
- Code Formatting, Linting, and Vetting (`fmt`, `go vet`, `golint`)
- Using Go Tooling: `go build`, `go run`, `go doc`, `go install`, `go fmt`

## Context & Concurrency Control
- Introduction to the `context` Package — Managing Lifetime
- Creating and Passing Contexts (`context.Background`, `context.TODO`)
- Timeouts and Deadlines with `context.WithTimeout`
- Cancellation with `context.WithCancel`
- Propagating Context Across Goroutines
- Combining Context with Channels and WaitGroups
- Practical Use Cases — Graceful Shutdowns, HTTP Requests, and Task Management

## Go Runtime and Reflection
- Overview of the Go Runtime System
- Memory Management and Garbage Collection in Go
- Understanding Stack Growth and Escape Analysis
- Using the `runtime` Package for Introspection
- Introduction to Reflection with the `reflect` Package
- Inspecting Struct Fields, Types, and Values Dynamically
- Practical Use Cases — Serialization, Validation, and Generic Utilities

## Design Patterns in Go
- Go’s Approach to Design Patterns — Simplicity over Inheritance
- Factory Pattern — Creating Flexible Constructors
- Chain of Responsibility Pattern — Decoupled Request Handling Pipelines
- Strategy Pattern — Replacing Conditionals with Interfaces
- Observer Pattern — Event-Driven Communication with Channels
- State Pattern — Managing Object Behavior Based on Internal State
- Pipeline Pattern — Composing Concurrent Stages for Data Processing
## Document Links
https://drive.google.com/file/d/1aSpEeLmfsTZc4PA1JT4wRs-wv2Y-_21T/view?usp=sharing

## Practices

### Session 1:

- Is there any destructor in Golang
- How can I override value type in this method

```go
func Test(value int) {
    println("test")
    value = 20
}
```
### Session 2
- why interface is a value type
- customize error handling
- in the code we have practice for panic and you should findout how if panics happened after codes how should be run

```go
defer func() {
	if r := recover(); r != nil {
		println("test")
	}
}()

TestPanic()

println("after panic")

func TestPanic() {
	println("before panic")
	panic("test")
}
```
### Session 3
- why we use wait and close channel in go func?
```go
	go func() {
		wg.Wait()
	}()
```
- what is blocking and deadlock in goroutine and how it would be happened please write it in an example?

- In this method I have implemented fanIn and I want to know why it goes wrong

```go
func fanIn(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil // disable this case
					continue
				}
				out <- v

			case v, ok := <-ch2:
				if !ok {
					ch2 = nil // disable this case
					continue
				}
				out <- v
			}

			// stop when both are nil → dead channel
			if ch1 == nil && ch2 == nil {
				close(out)
				return
			}
		}
	}()

	return out
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// sender 1
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(ch1)
	}()

	// sender 2
	go func() {
		for i := 100; i < 105; i++ {
			ch2 <- i
			time.Sleep(350 * time.Millisecond)
		}
		close(ch2)
	}()

	// merged output
	out := fanIn(ch1, ch2)

	// consumer
	for v := range out {
		fmt.Println("Received:", v)
	}
}
```
