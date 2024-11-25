package main

import (
	"fmt"

	"github.com/aidosgal/image-processing-service/internal/config"
)

func main() {
    cfg := config.MustLoad()

    fmt.Println(cfg)
}
