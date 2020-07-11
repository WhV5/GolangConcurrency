/**
* @Author : henry
* @Data: 2020-07-05 11:26
* @Note:
**/

package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("jane", "abc123")
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), "userID", userID)
	ctx = context.WithValue(ctx, "authToken", authToken)
	HandleRequest(ctx)
}

func HandleRequest(ctx context.Context) {
	fmt.Printf(
		"handle response for %v (%v)",
		ctx.Value("userID"),
		ctx.Value("authToken"),
	)
}
