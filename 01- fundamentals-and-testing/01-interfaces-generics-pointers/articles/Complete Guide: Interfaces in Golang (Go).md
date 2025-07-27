# Complete Guide: Interfaces in Golang (Go)

## Basic Explanation of Interfaces in Go

### What They Are and How They Work

In Go, an interface defines a set of method signatures (i.e., method names, parameters, and return types) without implementing any concrete behavior. Think of interfaces as a behavioral contract: if a type has those methods, it "fulfills the contract" and can be used wherever that interface is expected.

Unlike languages such as Java or C#, Go does not use an `implements` keyword — conformance is implicit. This means that as long as a type has all the required methods, it automatically satisfies the interface. For example:

```go
type MyInterface interface {
    MethodX() string
}
```

Any type that implements `MethodX() string` is considered to implement `MyInterface` without explicit declaration. This is an example of structural typing (often referred to as "duck typing"): if something behaves like a duck, it is a duck.

### Syntax and Simple Example

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

The `Reader` interface (similar to `io.Reader`) specifies that any type with a method `Read([]byte) (int, error)` is a `Reader`. It doesn’t matter what fields or type it is—just the method signature.

Types can implement multiple interfaces by providing the required methods.

### Design Principles Behind Interfaces

Go promotes simplicity and composition. Interfaces are typically small and focused, often having only one or two methods. For example:

* `Write` ➝ `Writer`
* `Notify` ➝ `Notifier`

While naming can be adapted to other languages, Go interfaces are almost always named in English. Go encourages composition over inheritance. Interfaces allow generic code without coupling to specific types.

### Example: `fmt.Stringer`

```go
type Stringer interface {
    String() string
}
```

Any type that implements `String()` can be formatted with `fmt.Println` automatically, enabling decoupled and extensible code.

## Practical Examples

### Example 1: Simple `Shape` Interface

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

type Rectangle struct {
    Width, Height float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func printArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    c := Circle{Radius: 10}
    r := Rectangle{Width: 5, Height: 6}
    printArea(c)
    printArea(r)
}
```

This demonstrates polymorphism through interfaces.

### Example 2: Standard Library Interface (`io.Writer`) and Testing

```go
func SaveData(w io.Writer, data []byte) error {
    _, err := w.Write(data)
    return err
}

func main() {
    data := []byte("Hello, Gophers!\n")

    file, _ := os.Create("out.txt")
    defer file.Close()
    SaveData(file, data)

    var buf bytes.Buffer
    SaveData(&buf, data)
    fmt.Println(buf.String())

    type MockWriter struct {
        Written []byte
    }
    func (m *MockWriter) Write(p []byte) (int, error) {
        m.Written = append(m.Written, p...)
        return len(p), nil
    }

    mw := &MockWriter{}
    SaveData(mw, data)
    fmt.Printf("Mock output: %q\n", string(mw.Written))
}
```

This highlights how interfaces enhance flexibility and testability.

## Comparing Interfaces vs Concrete Types

### Key Differences

* **Abstraction vs Implementation**: Interfaces describe the *what*, structs implement the *how*.
* **Flexibility**: Interfaces allow generic code. Structs are specific.
* **Encapsulation**: Interfaces hide implementation details. Structs expose full internal behavior.
* **Extensibility**: Interfaces must remain stable; adding a method breaks existing types. Structs can evolve without breaking compatibility.
* **Testability**: Interfaces enable mocking. Concrete types may require real resources unless designed to be replaceable.
* **Performance**: Interfaces add a small indirection; structs offer direct access. Prefer clarity over premature optimization.

## Best Practices & Design Patterns

### Guidelines

* **Define interfaces in the consuming package**.
* **Keep interfaces small and focused**.
* **Compose small interfaces to build larger behaviors**:

  ```go
  type ReadWriter interface {
      io.Reader
      io.Writer
  }
  ```
* **Design APIs to accept interfaces and return structs**.
* **Use interfaces in tests, but avoid over-abstraction**.
* **Encapsulate concrete implementations using interfaces if needed**.
* **Know the zero value behavior**: an interface with a nil concrete value is not nil.

### Rob Pike's Advice

> "Don’t design with interfaces initially. Discover them as the design evolves."

---

By following these principles, your Go code becomes cleaner, more extensible, and easier to test and maintain.

---

### References

* [Effective Go](https://go.dev/doc/effective_go)
* [Earthly - Go Interfaces](https://earthly.dev/blog/software-design-go-interfaces/)
* [DEV: Accept Interfaces, Return Structs](https://dev.to/shrsv/designing-go-apis-the-standard-library-way-accept-interfaces-return-structs-410k)
* [Victor Pierre - Five Best Practices](https://victorpierre.dev/blog/five-go-interfaces-best-practices/)
* [Bianca Rosa - Clean Go Code](https://biancarosa.com.br/pt/posts/go_clean_code_3/)
