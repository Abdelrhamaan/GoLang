package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// func main() {
// 	// 1️⃣ Create a base context
// 	// context is a type from context package
// 	// context includes a deadline, a cancelation signal, and values scoped to a request
// 	// context is used to pass values between API boundaries
// 	// context todo is a placeholder context or we plane to use it later
// 	// context todo shouldn't carry any values, cancelation signal or deadline
// 	// if you see the implementation of context.TODO() and context.background() they nearly do the same thing
// 	ctx := context.TODO()

// 	// 2️⃣ Derive a child context that carries a value
// 	ctxWithValue := context.WithValue(ctx, "name", "Mohammed")

// 	// ✅ Use fmt.Printf / fmt.Println with proper verbs
// 	fmt.Println("Parent Context is:", ctx) // still prints a pointer (that's expected)

// 	// If you really want to see the concrete type:
// 	fmt.Printf("Parent Context type: %T\n", ctx)

// 	// The derived context – also a pointer, but we can show its type:
// 	fmt.Printf("Derived Context type: %T\n", ctxWithValue)

// 	// 👉 Retrieve the stored value and assert its type
// 	if name, ok := ctxWithValue.Value("name").(string); ok {
// 		fmt.Println("Ctx Value (string):", name) // → Mohammed
// 	} else {
// 		fmt.Println("Ctx Value: <unexpected type>")
// 	}
// }

// func checkOddEvent(ctx context.Context, num int) string {
// 	select {
// 	// context.Done() returns empy channel which mean cancellation signal
// 	// we recieve a non nil value here
// 	case <-ctx.Done():
// 		return "Oeration Canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("Num %d is even", num)
// 		}
// 		return fmt.Sprintf("Num %d is odd", num)
// 	}
// }

// func main() {
// 	ctx := context.TODO()

// 	result := checkOddEvent(ctx, 5)
// 	fmt.Println("Result num is: ", result)

// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
// 	// cancel context before main is finished
// 	// it will send a cancelletion signal
// 	defer cancel()
// 	// this will implement cancel mechanism due to time out and any line after that will not executed
// 	// and defer cancel() will be called
// 	time.Sleep(3 * time.Second)
// 	result = checkOddEvent(ctx, 10)
// 	fmt.Println("Result num is: ", result)

// }

func doWOrk(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Working Cancelled")
			// return is important here to go out from inifinte loop
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func logWithContext(ctx context.Context, msg string) {
	requestValue := ctx.Value("requestId")
	log.Printf("Request log %v - %v", requestValue, msg)
}

func main() {
	// context background is not something that run in background
	// its abase step which we can drive another contexts from it
	ctxRoot := context.Background()
	ctx, cancel := context.WithTimeout(ctxRoot, 2*time.Second)

	defer cancel()
	// =========================
	// we can also cancel it manually without timeout by using withCancel
	// ctx, cancel := context.WithCancel(ctxRoot)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	cancel()
	// }()
	// ===========================
	// you can add any number of key value to context
	ctx = context.WithValue(ctx, "requestId", "jkdjfkdjfkdjfkjdf")
	ctx = context.WithValue(ctx, "ip", "124556535fd")
	go doWOrk(ctx)

	time.Sleep(3 * time.Second)

	if requestIdValue := ctx.Value("requestId"); requestIdValue != nil {
		fmt.Println("Request Id Value is: ", requestIdValue)
	} else {
		fmt.Println("No Request Id value found...")
	}
	logWithContext(ctx, "test logs with request")

}
