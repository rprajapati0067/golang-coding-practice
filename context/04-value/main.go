package main

import (
	"context"
	"fmt"
)

type database map[string]bool
type userIDKeyType string

var db database = database{
	"jane": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	processRequest(ctx, "jane")
}
func processRequest(ctx context.Context, userid string) {
	vctx := context.WithValue(ctx, userIDKeyType("userIDKey"), "jane")
	ch := checkMemberShip(vctx)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

func checkMemberShip(ctx context.Context) <-chan bool {
	ch := make(chan bool)

	go func() {
		defer close(ch)
		userid := ctx.Value(userIDKeyType("userIDKey")).(string)
		status := db[userid]
		ch <- status
	}()
	return ch
}
