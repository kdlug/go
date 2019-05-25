# Maps

Map is a reference type. Meaning it holds a reference to the underlying data structure, much like a pointer.   

A gotcha that reference types introduce is that maps can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic. You can read more about maps here.  

Therefore, you should never initialize an empty map variable:  

```go
var m map[string]string
```

Instead, you can initialize an empty map like we were doing above, or use the make keyword to create a map for you:

```go
m = map[string]string{}

// OR

m = make(map[string]string)
```

An interesting property of referenve type is that you can modify them without passing them as a pointer. 

```go
type MyMap map[string]string

func (m MyMap) Add(key, value string) {
    m[key] = value
}
```
