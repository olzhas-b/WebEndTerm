package main

import (
	"context"
	"runtime"

	"github.com/Torebekov/shop-back/modules/daylight"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ctx := context.Background()
	daylight.Start(ctx)
}
