package main

import (
	"github.com/ShreyashSri/test/internal"
	"github.com/ShreyashSri/test/internal/boot"
)

func main() {
	boot.LoadEnv()
	boot.InitFirebase()
	internal.StartServer()
}
