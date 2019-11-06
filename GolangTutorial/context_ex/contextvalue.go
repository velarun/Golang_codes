package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func newFunc(ctx context.Context, key string) {
	fmt.Println(ctx.Value(key))
}

func cancelFunc(ctx context.Context) error {
	return errors.New("Failed")
}

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "color", "black")

	newFunc(ctx, "color")

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		err := cancelFunc(ctx)

		if err != nil {
			cancel()
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Halted")
	case <-time.After(time.Second * 1):
		fmt.Println("Succeed")
	}

	ctxx := context.Background()
	ctxx, _ = context.WithTimeout(ctxx, 1*time.Second)

	req, _ := http.NewRequest(http.MethodGet, "http://www.google.com", nil)
	client := &http.Client{}

	req = req.WithContext(ctxx)
	time, ok := ctxx.Deadline()

	fmt.Println(time, ok)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed: ", err)
		return
	}
	fmt.Println("Response code:", resp.StatusCode)

}
