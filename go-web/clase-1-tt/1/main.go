package main

import (
	"context"
	"fmt"
)

type RequestType string
const (
	REQUEST RequestType = "request"
)

type Request struct {
	IP		string
	User	string
}

func Handler(ctx context.Context) {
	// logger
	req := ctx.Value(REQUEST).(Request)
	fmt.Printf("ip: %s - user: %s\n", req.IP, req.User)

	// process
	fmt.Println("process finished")
}

func main() {
	ctx := context.Background()

	// request
	ctx = context.WithValue(ctx, REQUEST, Request{IP: "192.168.0.0", User: "Jane"})

	// handler
	Handler(ctx)
}