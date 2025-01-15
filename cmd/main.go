package main

import (
	"context"
	"sync"

	"github.com/rodrigosscode/easy-user/internal/infrastructure/setup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	setup.
		NewConfig().
		InitLogger().
		WithAppConfig().
		WithDB().
		WithRouter().
		WithWebServer().
		Start(ctx, &wg)

	wg.Wait()
}
