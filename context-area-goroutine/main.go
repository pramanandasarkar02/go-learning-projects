package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context){
	for{
		select{
		case <- ctx.Done():
			fmt.Printf("Task cancelled: %v\n", ctx.Err())
			return 
		default:
			fmt.Println("working")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	go task(ctx)

	time.Sleep(3 * time.Second)
}