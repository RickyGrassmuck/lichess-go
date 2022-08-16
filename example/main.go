package main

import (
	"context"
	"fmt"
	lichess "github.com/RickyGrassmuck/lichess-go/lichess"
)

func main() {
	c, err := lichess.NewClient("https://lichess.org")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	c.ApiUser(context.Background(), "joyrida")
}
