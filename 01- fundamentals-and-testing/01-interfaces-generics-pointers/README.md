# Week 01 – Interfaces, Generics and Pointers

## Key Concepts

- Implicit interfaces and polymorphism
- Type parameters and constraints in generics
- Pointer vs value semantics in Go

## Articles & Resources

- [Golang Interfaces Explained – Alex Edwards](https://www.alexedwards.net/blog/interfaces-explained)
- [Go Generics – Official Tutorial](https://go.dev/doc/tutorial/generics)
- [Understanding Go Pointers – Dave Cheney](https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back)

## Code Examples

- `shape_interface.go`: interface `Shape` with `Circle` and `Rectangle`
- `max_generic.go`: generic `Max[T]` function
- `swap_pointers.go`: swap values using pointers

## Practice

- [x] Implemented polymorphic `Shape` example
- [x] Wrote a generic function for max/min
- [x] Created pointer-based `swap` function

## Notes

> Interfaces are satisfied implicitly — no need to declare "implements".
> Generics simplify reusability but may introduce complexity if overused.
> Pointers are passed by value (address), enabling indirect mutation.
