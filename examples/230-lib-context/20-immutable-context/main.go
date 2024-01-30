package main

import (
	"context"
	"fmt"
)

// It’s  recommended to put the context.Context parameter as the first parameter in a function
func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: someKey's value is %s\n", ctx.Value("someKey"))

	// When using contexts, it’s important to know that the values stored in a specific context.Context are immutable, meaning they can’t be changed.
	// It doesn't modify the context, it wraps parent context
	anotherCtx := context.WithValue(ctx, "someKey", "anotherValue")
	doSomethingElse(anotherCtx)

	fmt.Printf("doSomething: someKey's value is %s\n", ctx.Value("someKey"))

}

func doSomethingElse(ctx context.Context) {
	fmt.Printf("doSomethigElse: someKey's value is %s\n", ctx.Value("someKey"))
}
func main() {
	// ctx := context.TODO() // one of two ways to create an empty (or starting) context; can be used as a placeholder
	ctx := context.Background() // it’s designed to be used where you intend to start a known context; a good option as default

	// Add a new value to a context
	// When using contexts, it’s important to know that the values stored in a specific context.Context are immutable, meaning they can’t be changed.
	// It doesn't modify the context, it wraps parent context
	ctx = context.WithValue(ctx, "someKey", "someValue")
	doSomething(ctx)
}
