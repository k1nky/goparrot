package main

import (
	"github.com/k1nky/goparrot/internal/config"
)

func main() {
	config.LoadConfig("../internal/config/tests/config01.yaml")
}
