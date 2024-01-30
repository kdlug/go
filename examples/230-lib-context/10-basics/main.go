package main

// https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
import (
	"context"
	"fmt"
)

// It’s  recommended to put the context.Context parameter as the first parameter in a function
func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: someKey's value is %s\n", ctx.Value("someKey"))
}

func main() {
	// ctx := context.TODO() // one of two ways to create an empty (or starting) context; can be used as a placeholder
	ctx := context.Background() // it’s designed to be used where you intend to start a known context; a good option as default

	// Add a new value to a context
	ctx = context.WithValue(ctx, "someKey", "someValue")
	doSomething(ctx)
}
