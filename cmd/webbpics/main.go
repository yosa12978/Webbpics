package main

import (
	"math/rand"
	"time"

	"github.com/joho/godotenv"
	"github.com/yosa12978/webbpics/pkg"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	pkg.Run()
}
