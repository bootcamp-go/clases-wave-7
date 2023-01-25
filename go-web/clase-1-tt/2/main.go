package main

import (
	"context"
	"fmt"
	"time"
)

func service(url string, ch chan string) {
	fmt.Println("request url:", url)
	time.Sleep(time.Second * 5)
	ch <- "data"
}

func fetchData(ctx context.Context, url string) (string, error) {
	
	ch := make(chan string)
	go service(url, ch)

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case data := <- ch:
			return data, nil
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	// not needed but in case its needed to be cancalled manually
	go func ()  {
		time.Sleep(time.Second)
		cancel()
	}()

	data, err := fetchData(ctx, "www.google.com")
	if err != nil {
		fmt.Println("request error:", err)
		return
	}

	fmt.Println("request ok:", data)
}