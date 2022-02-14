package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(Module).Run()
}
