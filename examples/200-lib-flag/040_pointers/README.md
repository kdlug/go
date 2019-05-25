# Pointers

When should we use pointers?
- When we want to improve performance. If we have a large chunk of data and we don't want to pass that around through a program we could just pass the address when that data is stored  
- When we need to change something at a certain location.

Everything in GO is passed by value (not by copy, not by value). 

Creating pointer
- Pointer types use an asterisk(*) prefix to type pointed to, f.ex
```go
var a *int // a pointer to an integer
```
- Use the addres of operator (&) to get address of variable

```go
a := "foo"
&a // gets an address of a
```

Dereferencing pointers
- By using an asterisk(*) prefix before a variable name
```go
var a *int
b := "foo"
a := &b

fmt.Println(a) // address
fmt.Println(*a) // foo
```
- Complex types (eg. pointers to structs) are automatically dereferenced


Types with internal pointers
- slices
- maps

When we copy slices or maps, we actually copy only address which points to underlying data.